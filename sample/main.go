package main

import (
	"fmt"

	gb "github.com/takoyaki-3/go-binaryfile"
)

func main(){
	var buf []byte
	if err:=gb.LoadFromPath("./main.go",&buf);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(buf))

	gb.DumpToFile(&buf,"main2.go")
}
