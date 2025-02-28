package submod

import "log"

func SayHello() {
	panic(nil)
	//log.Println("Hello!")
}

func SayHelloV1() {
	log.Println("SayHelloV1!")
}
