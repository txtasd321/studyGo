package main

import (
	"fmt"
	"os"
)

// TODO写文件
func main() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err %v", err)
		return
	}
	// write
	fileObj.Write([]byte("zhoulin mengbi\n"))
	// writeString
	fileObj.WriteString("周林解释不了！")
	fileObj.Close()

}
