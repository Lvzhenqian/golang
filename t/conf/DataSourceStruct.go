package conf

type PythonConf struct {
	BootstrapToken string            `json:"BOOTSTRAP_TOKEN"`
	NODES          Nodes             `json:"NODES"`
	EXTENDNODE     map[string]string `json:"EXTENDNODE"`
	SHRINKNODE     Shrinknode        `json:"SHRINKNODE"`
	SSH_SETTING    SSH_SETTING       `json:"SSH_SETTING"`
	ETCD           Etcd              `json:"ETCD"`
	NFS            NFS               `json:"NFS"`
	NGINX          Nginx             `json:"NGINX"`
	ENGINE         Engine            `json:"ENGINE"`
	MYSQL          Mysql             `json:"MYSQL"`
}

type Engine struct {
	ReplicasGateway    int64  `json:"REPLICAS_GATEWAY"`
	ReplicasMVS        int64  `json:"REPLICAS_MVS"`
	ReplicasDirectory  int64  `json:"REPLICAS_DIRECTORY"`
	ReplicasDispatcher int64  `json:"REPLICAS_DISPATCHER"`
	ReplicasHotel      int64  `json:"REPLICAS_HOTEL"`
	ReplicasLive       int64  `json:"REPLICAS_LIVE"`
	EnableCcAgent      int64  `json:"ENABLE_CC_AGENT"`
	EnableGsRegit      int64  `json:"ENABLE_GS_REGIT"`
	AgentNodeID        int64  `json:"AGENT_NODE_ID"`
	ZookeeperCcAddress string `json:"ZOOKEEPER_CC_ADDRESS"`
	LogDownloadURL     string `json:"LOG_DOWNLOAD_URL"`
	BeatsLimitid       string `json:"BEATS_LIMITID"`
	BeatsVsbrainAddr   string `json:"BEATS_VSBRAIN_ADDR"`
	Licenes            string `json:"LICENES"`
	EnableInstruments  int64  `json:"ENABLE_INSTRUMENTS"`
	HarborEnable       int64  `json:"HarborEnable"`
	HarborUser         string `json:"HarborUser"`
	HarborPasswd       string `json:"HarborPasswd"`
	HarborEmail        string `json:"HarborEmail"`
}

type Etcd struct {
	IPS []string `json:"IPS"`
}

type Mysql struct {
	UseMysqlInK8S int64  `json:"USE_MYSQL_IN_K8S"`
	MysqlAddress  string `json:"MYSQL_ADDRESS"`
	MysqlUser     string `json:"MYSQL_USER"`
	MysqlPassword string `json:"MYSQL_PASSWORD"`
}

type NFS struct {
	IPS string `json:"IPS"`
}

type Nginx struct {
	ReplicasNginx int64  `json:"REPLICAS_NGINX"`
	Dashboard     string `json:"DASHBOARD"`
	Grafana       string `json:"GRAFANA"`
	Wss           string `json:"WSS"`
	CAKey         string `json:"CA_KEY"`
	CACRT         string `json:"CA_CRT"`
	AuthUser      string `json:"AUTH_USER"`
	AuthPassword  string `json:"AUTH_PASSWORD"`
}

type Nodes struct {
	APISERVER     string            `json:"APISERVER"`
	APISERVER_URL string            `json:"APISERVER_URL"`
	IPS           map[string]string `json:"IPS"`
}

type SSH_SETTING struct {
	SSH_PORT     string `json:"SSH_PORT"`
	SSH_USER     string `json:"SSH_USER"`
	SSH_PASSWORD string `json:"SSH_PASSWORD"`
}

type Shrinknode struct {
	NODES []string `json:"NODES"`
}
