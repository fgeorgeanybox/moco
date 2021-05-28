package e2e

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	mocov1beta1 "github.com/cybozu-go/moco/api/v1beta1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

//go:embed testdata/backup.yaml
var backupYAML string

//go:embed testdata/restore.yaml
var restoreYAML string

var _ = Context("backup", func() {
	if doUpgrade {
		return
	}

	var restorePoint time.Time

	It("should create a bucket", func() {
		kubectlSafe(nil, "run", "--command", "make-bucket", "--image=moco-backup:dev", "--",
			"s3cmd", "--host=minio.default.svc:9000", "--host-bucket=minio.default.svc:9000", "--no-ssl",
			"--access_key=minioadmin", "--secret_key=minioadmin", "mb", "s3://moco")
	})

	It("should construct a source cluster", func() {
		kubectlSafe(fillTemplate(backupYAML), "apply", "-f", "-")
		Eventually(func() error {
			cluster, err := getCluster("backup", "source")
			if err != nil {
				return err
			}
			for _, cond := range cluster.Status.Conditions {
				if cond.Type != mocov1beta1.ConditionHealthy {
					continue
				}
				if cond.Status == corev1.ConditionTrue {
					return nil
				}
				return fmt.Errorf("cluster is not healthy: %s", cond.Status)
			}
			return errors.New("no health condition")
		}).Should(Succeed())

		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-writable", "source", "--",
			"-e", "CREATE DATABASE test")
		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-writable", "source", "--",
			"-D", "test", "-e", "CREATE TABLE t (i INT PRIMARY KEY AUTO_INCREMENT, data TEXT NOT NULL) ENGINE=InnoDB")
		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-writable", "source", "--",
			"-D", "test", "--init_command=SET autocommit=1", "-e", "INSERT INTO t (data) VALUES ('aaa')")
	})

	It("should take a full dump", func() {
		kubectlSafe(nil, "-n", "backup", "create", "job", "--from=cronjob/moco-backup-source", "backup-1")
		Eventually(func() error {
			out, err := kubectl(nil, "-n", "backup", "get", "jobs", "backup-1", "-o", "json")
			if err != nil {
				return err
			}
			job := &batchv1.Job{}
			if err := json.Unmarshal(out, job); err != nil {
				return err
			}
			for _, cond := range job.Status.Conditions {
				if cond.Type != batchv1.JobComplete {
					continue
				}
				if cond.Status == corev1.ConditionTrue {
					return nil
				}
			}
			return errors.New("backup-1 has not been finished")
		}).Should(Succeed())
	})

	It("should take an incremental backup", func() {
		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-writable", "source", "--",
			"-D", "test", "--init_command=SET autocommit=1", "-e", "INSERT INTO t (data) VALUES ('bbb')")
		time.Sleep(1100 * time.Millisecond)
		restorePoint = time.Now().UTC()
		time.Sleep(1100 * time.Millisecond)
		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-admin", "source", "--",
			"-D", "test", "--init_command=SET autocommit=1", "-e", "FLUSH LOCAL BINARY LOGS")
		kubectlSafe(nil, "moco", "-n", "backup", "mysql", "-u", "moco-writable", "source", "--",
			"-D", "test", "--init_command=SET autocommit=1", "-e", "INSERT INTO t (data) VALUES ('ccc')")
		time.Sleep(100 * time.Millisecond)

		kubectlSafe(nil, "-n", "backup", "create", "job", "--from=cronjob/moco-backup-source", "backup-2")
		Eventually(func() error {
			out, err := kubectl(nil, "-n", "backup", "get", "jobs", "backup-2", "-o", "json")
			if err != nil {
				return err
			}
			job := &batchv1.Job{}
			if err := json.Unmarshal(out, job); err != nil {
				return err
			}
			for _, cond := range job.Status.Conditions {
				if cond.Type != batchv1.JobComplete {
					continue
				}
				if cond.Status == corev1.ConditionTrue {
					return nil
				}
			}
			return errors.New("backup-2 has not been finished")
		}).Should(Succeed())

		cluster, err := getCluster("backup", "source")
		Expect(err).NotTo(HaveOccurred())
		Expect(cluster.Status.Backup.BinlogSize).NotTo(Equal(int64(0)))
	})

	It("should destroy the source then restore the backup data", func() {
		kubectlSafe(nil, "-n", "backup", "delete", "mysqlclusters", "source")

		tmpl, err := template.New("").Parse(restoreYAML)
		Expect(err).NotTo(HaveOccurred())
		buf := new(bytes.Buffer)
		err = tmpl.Execute(buf, struct {
			MySQLVersion string
			RestorePoint string
		}{
			mysqlVersion,
			restorePoint.Format(time.RFC3339),
		})
		Expect(err).NotTo(HaveOccurred())

		kubectlSafe(buf.Bytes(), "apply", "-f", "-")
		Eventually(func() error {
			cluster, err := getCluster("backup", "target")
			if err != nil {
				return err
			}
			for _, cond := range cluster.Status.Conditions {
				if cond.Type != mocov1beta1.ConditionHealthy {
					continue
				}
				if cond.Status == corev1.ConditionTrue {
					return nil
				}
			}
			return errors.New("target is not healthy")
		}).Should(Succeed())

		out := kubectlSafe(nil, "moco", "-n", "backup", "mysql", "target", "--",
			"-N", "-D", "test", "-e", "SELECT COUNT(*) FROM t")
		count, err := strconv.Atoi(strings.TrimSpace(string(out)))
		Expect(err).NotTo(HaveOccurred())
		Expect(count).To(Equal(2))
	})
})
