package node

func New() *Node {
	n := &Node{}
	n.Port = "9"
	return n
}

type Node struct {
	NodeId      string `名称`
	NodeName    string
	ServerIp    string
	Port        string
	PublicCert  string
	PrivateCert string
	Connection  Connection
}

// 注册发现
func (n *Node) Register() {

}

// 注册发现
func (n *Node) getServerIp() string {
	return "127.0.0.1"
}
