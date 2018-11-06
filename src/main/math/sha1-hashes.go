package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	//md5 类似
	h := sha1.New()

	//写入处理内容
	h.Write([]byte(s))

	//额外追加字节
	bs := h.Sum(nil)

	fmt.Println(s)

	fmt.Printf("%x\n", bs)
}
