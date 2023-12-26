package main

import "fmt"

type AInterface interface {
	Test01()
	Test02()
}

type BInterface interface {
	Test01()
	Test02()
}

type CInterface interface {
	AInterface
	BInterface
}

func main() {
	fmt.Println("111")
}
