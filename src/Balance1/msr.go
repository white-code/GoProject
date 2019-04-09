package Balance1

import "fmt"

type BlanceMsr struct {
	Allblance map[string]Blance
}

func(p *BlanceMsr)registeBalancer(name string,b Blance){
	p.Allblance[name] = b
}
var msr  = &BlanceMsr{
	Allblance:make(map[string]Blance),
}
func Registerbalancer(name string,b Blance){
	msr.registeBalancer(name,b)
}

func Doblance(name string,insts []*Instance)(inst *Instance,err error){
	balance,ok := msr.Allblance[name]
	if !ok{
		err = fmt.Errorf("Not find %s balancer",name)
		return
	}
	fmt.Printf("user %s balancer\n",name)
	inst,err = balance.Doblance(insts)
	return
}
func (p *BlanceMsr) printMethod(){
	for k,_ := range p.Allblance{
		fmt.Printf("负载均衡算法：%s\n",k)
	}
}
func PrintMethod(){
	msr.printMethod()
}