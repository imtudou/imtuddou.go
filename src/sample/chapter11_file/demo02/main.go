package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "../../../../doc/222.txt"
	os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Println(filePath)
}
