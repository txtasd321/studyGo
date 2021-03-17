# 接口(interface)

接口是一种类型，是一种特殊的类型，他规定了变量有哪些方法，
在编程中会遇到一下场景：
我不关心一个变量是什么类型，我只关心等调用它的什么方法

## 接口的定义

```go
type 接口名 interface {
    方法名1(参数1,参数2...)(返回值1, 返回值2...)
    方法名2(参数1,参数2...)(返回值1, 返回值2...
                     ...
}
```

用来给变量\参数\返回值等设置类型

## 接口的实现

一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以成为这个接口类型的变量。

![](http://qn.kuinit.cn/20210312171049.png)



```go
package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫走")
}

// func (c cat) eat() {
// 	fmt.Println("吃鱼")
// }

// !有没有传参很重要，方法必须与接口相同，上例则不行
func (c cat) eat(str string) {
	fmt.Printf("猫吃%s...\n", str)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("jidong")
}

func (c chicken) eat(str string) {
	fmt.Println("吃饲料！")
}

func main() {
	var a1 animal

	bc := cat{
		name: "淘气",
		feet: 4,
	}
	kfc := chicken{
		feet: 2,
	}

	a1 = bc
	a1.eat("鲈鱼")
	fmt.Println(a1)
	fmt.Printf("%T", a1)
	a1 = kfc
	fmt.Println(a1)
}

```

### 使用值接收者实现接口与使用指针接收者实现接口的区别？

使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存

使用指针接收者实现接口，只能存结构体指针

## 接口和类型的关系

多个类型可以实现同 一个接口

一个类型可以实现多个接口

## 空接口

空接口没有必要起名字，通常定义成下面的格式

```go
type xx interface {} // 空接口没有必要起名字
```

当不知道传递的是什么时，使用空接口

所有的类型都实现了空接口，也就是任意类型都能保存到空接口中。

只有当两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口，不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗

### 类型断言

空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？

![](http://qn.kuinit.cn/20210313164743.png)