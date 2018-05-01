package main

import (
	"fmt"
	"goinaction/code/chapter5/listing74/entities"
)

func main() {
	// 由于内部类型 user 是未公开的,这段代码无法直接通过结构字面量的方式初 始化该内部类型。
	// 不过,即便内部类型是未公开的,内部类型里声明的字段依旧是公开的。
	// 既然 内部类型的标识符提升到了外部类型,这些公开的字段也可以通过外部类型的字段的值来访问
	a := entities.Admin{
		Rights: 10,
	}

	a.Name = "Bill"
	a.Email = "Bill@email.com"

	fmt.Printf("User : %v\n", a)
}
