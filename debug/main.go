package main

import (
	"fmt"

	"github.com/ArakiTakaki/Context"
)

func main() {

	fmt.Println("--------------------------------------")
	context.Put("key", "value")
	context.Put("This", "Value")
	tes, err := context.Get("key")
	if err != nil {
		panic(err)
	}

	var swt bool
	var key string

	tes, err = context.Get("This")
	fmt.Println("This : " + tes.(string))

	key = "This"
	swt = context.Remove(key)
	debug(swt, key)

	key = "is"
	swt = context.Remove(key)
	debug(swt, key)

	fmt.Println(tes)
	fmt.Println(context.Len())
	fmt.Println("--------------------------------------")
}

// debug test
func debug(swt bool, key string) {
	if swt {
		fmt.Println(key + "を削除しました")
	} else {
		fmt.Println(key + "がみつかりませんでした")
	}
}
