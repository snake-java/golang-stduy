package demo12

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("......")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	SpeakTo(host)
// }

type ProgrammerInterface interface {
	sayhello()
}

type GoProgrammerInterface struct {
}

func (t *GoProgrammerInterface) sayhello() {
	fmt.Println("Go==========ProgrammerInterface")
}

type JavaProgrammerInterface struct {
}

func (t *JavaProgrammerInterface) sayhello() {
	fmt.Println("JAVA=======ProgrammerInterface")
}

func sayhellbyto(gop ProgrammerInterface) {
	gop.sayhello()
}

func Test_sayHelloByTO(t *testing.T) {
	g := new(GoProgrammerInterface)
	j := new(JavaProgrammerInterface)
	sayhellbyto(g)
	sayhellbyto(j)
}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("CHAO")
}
