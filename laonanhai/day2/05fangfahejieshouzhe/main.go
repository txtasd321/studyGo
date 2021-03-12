package main

import "fmt"

// TODO:方法和接受者
// dog这是一个狗的结构体

// ?方法中使用值传递不能给改变原有的值，传递的是拷贝的值
// ?方法中使用指针传递能够在方法中改变传递者的内容

// !给自定义类型加方法
// !不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法
type dog struct {
	name string
}

func newGog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// * 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:wang", d.name)
}

func main() {
	d1 := newGog("zhou")
	d1.wang()
}
