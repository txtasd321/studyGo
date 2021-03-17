package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// TODO利用Open打开文件使用Read方法读
func readFormFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128)	// 指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
		}
		if err != nil {
			fmt.Println("read from file failed, err:%v", err)
			return
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// 利用Bufio读文件
func readFromFileBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	// 关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
			return
		}
		fmt.Print(line)
	}

}

func readFromByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFileBufio()
	// readFormFile()
	readFromByIoutil()
}
