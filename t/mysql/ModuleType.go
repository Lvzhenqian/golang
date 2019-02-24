package mysql

import "time"

type Kubernetes struct {
	ID              int32     `xorm:"autoincr unique"`
	FromAddress     string    `xorm:"pk unique notnull"`
	CreateTime      time.Time `xorm:"created"`
	UpDateTime      time.Time `xorm:"updated"`
	Apiserver       string    `xorm:"notnull"`
	ApiserverUrl    string
	IPIPMOD         string
	KubeBin         string
	KubeSSL         string
	ServiceCIDR     string
	ClusterCIDR     string
	SecurePort      int
	InsecurePort    int
	ClusterDNSSvcIP string
	NodePortRange   string
	ClusterK8sSvcIP string
	DockerData      string
	KubeletData     string
	KubeProxy       string
	BackUpDir       string
	ProxyMode       string
	MasterID        int32
	NodeID          int32
	ShrinkNode      int32
	ExtendNode      int32
}

type Master struct {
	ID     int32 `xorm:"pk autoincr unique"`
	HostIP string
}

type Node struct {
	ID         int32 `xorm:"pk autoincr unique"`
	HostIP     string
	InternetIP string
}

type ExtendNode struct {
	ID         int32 `xorm:"pk autoincr unique"`
	HostIP     string
	InternetIP string
}

type ShrinkNode struct {
	ID     int32 `xorm:"pk autoincr unique"`
	HostIP string
}

type SSH struct {
	ID          int `xorm:"pk autoincr unique"`
	SshPort     int
	SshUser     string
	SshPassword string
	SshKeyFile  string
	SshKeyPass  string
}

type BaseService struct {
	ID                int `xorm:"pk autoincr unique"`
	Nfs               string
	ReplicasNginx     int
	Dashboard         string
	Grafana           string
	Wss               string
	CaKey             string
	CaCrt             string
	NginxAuthUser     string
	NginxAuthPassword string
	MysqlInK8s        bool
	MysqlAddress      string
	MysqlUser         string
	MysqlPassword     string
}

type MatchvsEngine struct {
	ID                int `xorm:"pk autoincr unique"`
	Gateway           int
	Mvs               int
	Hotel             int
	Live              int
	Dispatcher        int
	Directory         int
	EnableCCagent     bool
	EnableGsRegit     bool
	NodeID            int32
	ZkAddress         string
	GsLogDownUrl      string
	License           string
	LicenseServer     string
	VsbrainAddress    string
	EnableInstruments bool
	HarborEnable      bool
	HarborUser        string
	HarborPassword    string
	HarborEmail       string
}
