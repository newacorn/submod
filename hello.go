package submod

import "log"
import "github.com/newacorn/anothermod"

func SayHello() {
	panic(nil)
	//log.Println("Hello!")
}

func SayHelloV2() {
	log.Println("SayHelloV2!")
	anothermod.HelloWorld()
}

func SayHelloV1() {
	log.Println("SayHelloV1!")
}
