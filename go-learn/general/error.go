package main

import (
	// "fmt" // error here, because fmt is not used
	// if still want to use fmt, use _ to ignore it
	"fmt"
	"reflect"
	"time"
)

// cannot declare a short variable declaration outside a block
// myVar := 1
var myVar = 1

func Main() {
	// { // if put this block here, it will not compile
	// var a int // error here, because a is not used

	x := 1
	_ = x
	// fmt.Println(x) // print 1
	// {
	// 	// 被称为变量幽灵
	// 	fmt.Println(x) // print 1
	// 	x := 2
	// 	fmt.Println(x) // print 2
	// }
	// fmt.Println(x) // print 1

	// var y = nil // error here, because nil is not a type
	var y interface{} = nil // ok
	_ = y

	// init nil is not suitable for slice and map type

	// when map is allocated by make, it can target with capacity
	// but when map is allocated by map, it can not target with capacity
	var m1 = make(map[string]int, 10)
	var m2 = map[string]int{"a": 1, "b": 2}
	_ = m1
	_ = m2

	// string is not a pointer type, so it can not be nil
	// var s1 string = nil // error here, because string is not a pointer type

	// nil is only available with pointer, channel, func, interface, map, or slice type

	// copyFunction()

	// range 返回的是键值，而不是只有值
	// rangeFunction()

	// multiArr := [][]int{{1, 3}, {1, 2}, {2, 3}}
	// fmt.Println(multiArr)

	// fmt.Println(changeStringText("something When Here", 0, 'S'))

	go func() {
		// 这儿并发代码
	}()
	// switchCaseFunction("a")
	// useChannelFunction()
	// tryDefer()
	reflectInGo()
}

func rangeFunction() {
	arrX := []string{"a", "b", "c"}
	for v := range arrX {
		fmt.Println(v) // print 0 1 2
	}

	for _, v := range arrX {
		fmt.Println(v) // print a b c
	}
}

func copyFunction() {
	x := [3]int{1, 2, 3}

	// 数组在函数中传参数是值复制
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // print [7 2 3]
	}(x)
	fmt.Println(x) // print [1 2 3]

	// 使用指针(pointer)来传参
	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(*arr) // print [7 2 3]
	}(&x)
	fmt.Println(x) // print [7 2 3]
}

func changeStringText(t string, num int, t2 byte) string {
	// string cannot be change, will be error
	// t[num] = t2 // error here

	// string convert to byte is copy (might be some memory lack), can instead of map[string] []byte, or using range to avoid memory allocation for increase performance.
	tBytes := []byte(t)
	tBytes[num] = t2

	return string(tBytes)
}

func switchCaseFunction(t string) {
	// fallthrough in golang mean continue search the case
	// unlike other language, switch case is stop or break by "break" keyword
	switch t {
	case "a":
		fmt.Println("yes a")
		// fallthrough
	case "b":
		fmt.Println("yes b")
		// fallthrough
	default:
		fmt.Println("not found")
	}
}

func useChannelFunction() {
	// will cause to deadlock
	var ch chan int
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}
	// get first result
	fmt.Println("result: ", <-ch)
	// do other work
	time.Sleep(2 * time.Second)
}

func tryDefer() {
	a := 1
	defer fmt.Println("after a ====>", a)
	fmt.Println("before a =====>", a)
	a++
	fmt.Println("a =====>", a)
}

func reflectInGo() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{"Bob", 20}
	// reflect.TypeOf() return the type of the variable
	fmt.Println(reflect.TypeOf(p)) // print main.Person
	// reflect.ValueOf() return the value of the variable
	fmt.Println(reflect.ValueOf(p)) // print {Bob 20}
	// reflect.ValueOf().NumField() return the number of the field of the variable
	fmt.Println(reflect.ValueOf(p).NumField()) // print 2
	// reflect.ValueOf().NumMethod() return the number of the method of the variable
	fmt.Println(reflect.ValueOf(p).NumMethod()) // print 0
	// reflect.ValueOf().IsValid() return the variable is valid or not
	fmt.Println(reflect.ValueOf(p).IsValid()) // print true
	// reflect.ValueOf().Field() return the field of the variable
	fmt.Println(reflect.ValueOf(p).Field(0)) // print Bob
	// reflect.ValueOf().FieldByName() return the field of the variable by name
	fmt.Println(reflect.ValueOf(p).FieldByName("Name")) // print Bob
	// reflect.TypeOf().Kind() return the kind of the variable
	fmt.Println(reflect.TypeOf(p).Kind()) // print struct
}
