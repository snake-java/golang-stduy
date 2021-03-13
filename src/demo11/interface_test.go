package demo11

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHellworld()
}

type GOProgrammer struct {
}

func (gop *GOProgrammer) WriteHellworld() {
	fmt.Println(1232313)
}

func Test_Goprogrammer(t *testing.T) {
	goprogrammer := new(GOProgrammer)
	goprogrammer.WriteHellworld()
}
