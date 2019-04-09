package Balance1

import (
	"errors"
)
func init(){
	msr.registeBalancer("round",&Randonmethod{})
}
type Roundmethod struct {
	curIndex int
}

func (p *Roundmethod) Doblance(insts []*Instance) (inst *Instance,err error){
	if len(insts) == 0{
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1)%lens
	return
}