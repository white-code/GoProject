package Balance1

import (
	"errors"
	"math/rand"
)
func init(){
	msr.registeBalancer("random",&Randonmethod{})
}

type Randonmethod struct {

}

func (p *Randonmethod) Doblance(insts []*Instance) (inst *Instance,err error){
	if len(insts) == 0{
		err = errors.New("No instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}