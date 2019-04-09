package Balance1

import "strconv"

type Instance struct {
	host string
	port int
}


func NewInstance(host string,port int) *Instance{
	return &Instance{
		host:host,
		port:port,
	}
}


func (p *Instance)GetHost() string{
	return p.host
}

func (p *Instance)GetPort() int{
	return p.port
}
//所有结构体只要是自己定义string方法，输出则会自己进行格式转换h j h j
func (p *Instance) String() string{
	return p.host + ":" + strconv.Itoa(p.port)
}