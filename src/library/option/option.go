package option

type Option struct {
	ServerPort string
	AdminPort  string
}

var o *Option

func init() {
	o = &Option{}
	o.ServerPort = ":8087"
	o.AdminPort = ":8086"
}

func Get() *Option {
	return o
}
