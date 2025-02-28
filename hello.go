package submod

import "log"

func SayHello() {
	panic(nil)
	//log.Println("Hello!")
}

func SayHelloV2() {
	log.Println("SayHelloV2!")
}

func SayHelloV1() {
	log.Println("SayHelloV1!")
}
