package constants

const (
	// MySQLDataPath is the path of MySQL data dir.
	MySQLDataPath = "/var/lib/mysql"

	// MySQLConfPath is the path of MySQL conf dir.
	MySQLConfPath = "/etc/mysql"

	// MySQLInitConfPath is the path of MySQL conf dir generated by moco-init.
	MySQLInitConfPath = "/etc/mysql-conf.d"

	// MySQLConfName is the filename of my.cnf.
	MySQLConfName = "my.cnf"

	// RunPath is the path of variable files which concerns MySQLd.
	RunPath = "/run"

	// LogDirPath is the path of /var/log/mysql.
	LogDirPath = "/var/log/mysql"

	// MySQLSlowLogName is the filename of slow query log for MySQL.
	MySQLSlowLogName = "mysql.slow"

	// TmpPath is the path for /tmp.
	TmpPath = "/tmp"

	// MyCnfSecretPath is the path for my.cnf formated credentials for CLI
	MyCnfSecretPath = "/mysql-credentials"
)

const (
	// FluentBitConfigName is a filename for fluent-bit conf.
	FluentBitConfigName = "fluent-bit.conf"

	// FluentBitConfigPath is the path for fluent-bit config directory.
	FluentBitConfigPath = "/fluent-bit/etc"
)

// Environment variables
const (
	PodNameEnvKey      = "POD_NAME"
	PodNamespaceEnvKey = "POD_NAMESPACE"
	ClusterNameEnvKey  = "CLUSTER_NAME"
)

// Secret keys to clone data from an external mysqld
const (
	CloneSourceHostKey         = "HOST"
	CloneSourcePortKey         = "PORT"
	CloneSourceUserKey         = "USER"
	CloneSourcePasswordKey     = "PASSWORD"
	CloneSourceInitUserKey     = "INIT_USER"
	CloneSourceInitPasswordKey = "INIT_PASSWORD"
)
