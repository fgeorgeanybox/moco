package constants

// MySQL user names for MOCO
const (
	AdminUser       = "moco-admin"
	AgentUser       = "moco-agent"
	ReplicationUser = "moco-repl"
	CloneDonorUser  = "moco-clone-donor"
	ExporterUser    = "moco-exporter"
	BackupUser      = "moco-backup"
	ReadOnlyUser    = "moco-readonly"
	WritableUser    = "moco-writable"
)

// my.cnf filenames for different kind of users.
const (
	AdminMyCnf    = AdminUser + "-my.cnf"
	ExporterMyCnf = ExporterUser + "-my.cnf"
	BackupMyCnf   = BackupUser + "-my.cnf"
	ReadOnlyMyCnf = ReadOnlyUser + "-my.cnf"
	WritableMyCnf = WritableUser + "-my.cnf"
)
