package main

import (
	"exam/11_interrface/service"
	"fmt"
)

func main() {
	num := 1
	var ptr *int = &num
	fmt.Println(ptr)
	user0 := service.User{}
	user1 := service.User{1, "gosha", 15}
	admin := service.admin{}
}
