package main

// import "fmt"
// import "unsafe"

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
	// "unsafe"
)

// global var
var (
	globalVar  = "globalVar"
	globalVar2 = "globalVar"
)

func init() {
	// 全局变量定义 》 init函数 》 main函数
	// golang will program first after go to main function
}

func HelloMain() {
	// chap3()
	// pointer()
	// pointerExercise()
	// interviewQuestion()
	// chap4()
	// chap5()
	// chap6()
	// exeChap6()
	// chap7()
	// exeChap7()
	// chap8()
	// chap9()
	// chap10()
	// chap11()
	// exeChap11()
	// chap12()
	// chap13()
	chap14()
}

// chapter03/demo1
func chap3() {
	// 定义变量/声明变量
	var i int
	// 给i赋值
	i = 10
	// 使用变量
	fmt.Println("i=", i)

	// variable method 1
	// 指定变量类型， 声明后诺不赋值， 使用默认值
	// int 默认值为0
	var p int
	fmt.Println("p=", p)

	// variable method 2
	// 根据值自行判断变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	// variable method 3
	// 省略var， 注意 ：= 左侧的变量不应该是已经声明过的， 否则会导致编译错误
	name := "tom"
	fmt.Println("name=", name)

	// multiple variable
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// multiple variable with different type
	var n4, name2, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name2=", name2, "n5=", n5)

	// test for int8 range
	// var n6 int8 = 128 wrong
	// formula for int8, int32, int64
	// 8-1 (-0 move to -128)
	// -128 ~ 127
	// -2 ^ 7 ~ 2 ^ 7 - 1

	// test for uint8 range
	// only absolute value
	// 0 ~ 255
	// 0 ~ 2 ^ 8 - 1

	// fmt.Printf for format print
	// 查看变量类型， 可以使用%T
	// fmt.Printf("type of n4 %t\n", n4)

	// 怎么查看一个变量的字节大小
	// var n6 int64 = 10
	// fmt.Printf("type of n6 %t  size of n6 %d", n6, unsafe.Sizeof(n6))

	// 可以用小就用小， 不用大 （保小不保大）
	// 数据类型越小， 使用的空间越少， 增加效率
	// var age byte = 89

}

// pointer
func pointer() {
	// basic data type
	var i int = 10
	// i 的地址
	fmt.Println("i的地址=", &i)

	// 下面的var ptr *int = &i
	// 1. ptr 是一个指针变量
	// 2. ptr 的类型 *int
	// 3. ptr 本身的值&i
	var ptr *int = &i
	// ptr 的地址
	fmt.Println("ptr=", ptr)
	fmt.Println("*ptr指向的值=", *ptr)
}

func pointerExercise() {
	var num int = 9
	fmt.Println("num address", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Println("num=", num)
}

func interviewQuestion() {
	a := 10
	b := 20

	// question: a exchange with b and without temp num
	a = a + b // a = 10 + 20
	b = a - b // b = 30 - 20 = 10(o:a)
	a = a - b // a = 30(a) - 10(b) = 20(o:b)
	// a, b = b, a
	fmt.Println("a=", a, "b=", b)
}

func chap4() {
	var name string
	var age int
	var salary float32
	var isPass bool
	// method 1
	// fmt.Println("Please enter your name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Please enter your age: ")
	// fmt.Scanln(&age)
	// fmt.Println("Please enter your salary: ")
	// fmt.Scanln(&salary)
	// fmt.Println(name, " ", age, " ", salary)
	// fmt.Printf("name type %T age type %T salary type %T\n", name, age, salary)
	// fmt.Printf("name is %v age is %v sal is %v isPass is %v", name, age, salary, isPass)

	// method 2
	fmt.Println("Please enter your name, age, salary, and use ' ' to separate: ")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	fmt.Printf("name is %v age is %v sal is %v isPass is %v", name, age, salary, isPass)
}

func chap5() {
	// fmt.Println("Enter the pyramid height:")
	var height int = 10
	// fmt.Scanf("%d", &height)

	for i := 0; i < height; i++ {
		for k := 0; k < height-i; k++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2-1; j++ {
			if j == 0 || j == i*2-2 {
				fmt.Print("*")
			} else if i == height-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// 9 x 9 multiple
	for i := 0; i <= 9; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(i, " * ", j, " = ", i*j, " ")
		}
		fmt.Println()
	}

	// 153 = 1 * 1 * 1 + 3 * 3 * 3 + 5 * 5 * 5
	var num int = 153
	var first int = num / 100
	var second int = num % 10
	var last int = num / 10 % 10
	if num == int(math.Pow(float64(first), 3)+math.Pow(float64(second), 3)+math.Pow(float64(last), 3)) {
		fmt.Println("num is a pyramidal number")
	} else {
		fmt.Println("num is not a pyramidal number")
	}

	// monthly days
	for i := 1; i <= 12; i++ {
		switch i {
		case 1, 3, 5, 7, 8, 10, 12:
			fmt.Println(i, " has 31 days")
		case 2:
			fmt.Println(i, " has 28 days")
		default:
			fmt.Println(i, " has 30 days")
		}
	}

	// for loop until number == 99
	var number int
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		if n == 99 {
			break
		} else {
			number++
		}
		fmt.Println(n)
	}

	// goto
	fmt.Println("ok1")
	// goto label1
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	// label1:
	fmt.Println("ok6")
	fmt.Println("ok7")
	fmt.Println("ok8")
}

func calculator(n1 int, n2 int, operator byte) int {
	var res int
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("operator error")
	}
	return res
}

func bonaNum(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return bonaNum(n-1) + bonaNum(n-2)
	}
}

func question2(n int) int {
	// f(1) = 3; f(n) = 2 * f(n-1) + 1
	if n == 1 {
		return 3
	} else {
		return 2*question2(n-1) + 1
	}
}

func question3(n int) int {
	// monkey eat half of total plus one peach one day, when days 10, peach left 1, how many peaches on day 1 to 10
	if n == 10 {
		return 1
	} else {
		return (question3(n+1) + 1) * 2
	}
}

func changeOutSideValue(n *int) {
	*n += 10
	fmt.Println("inside n=", *n)
}

func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func chap6() {
	fmt.Println(calculator(10, 20, '+'))
	fmt.Println(bonaNum(3))
	fmt.Println(question2(1))

	var num1 int = 10
	changeOutSideValue(&num1)
	fmt.Println("outside n=", num1)
	fmt.Println("sum", sum(1, 2, 3, 4))
}

// method 1
func numExchange(n1 int, n2 int) (int, int) {
	return n2, n1
}

// method 2
// func numExchange(n1 int, n2 int) (int, int) {
// 	t := n1
// 	n1 = n2
// 	n2 = t
// 	return n1, n2
// }

// method 3
// func numExchange(n1 int, n2 int) (int, int) {
// 	n1 = n1 + n2
// 	n2 = n1 - n2
// 	n1 = n1 - n2
// 	return n1, n2
// }

func exeChap6() {
	fmt.Println(numExchange(10, 20))

	// unknown name function
	// res1 := func(n1 int, n2 int) int {
	// 	return n1 + n2
	// }(10, 20)
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}

	res2 := res1(10, 20)

	fmt.Println(res2)
}

// (累加器)
func AddUpper() func(int) int {
	// 返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
	// 闭包是类， 函数是操作， n是字段
	// 当反复调用函数， 因为n是初始化一次，因此没调用一次就进行累计

	// var n int = 10 只进行一次， 后续的调用只跑return
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n += x
		str += "a"
		fmt.Println(str)
		return n
	}
}

func defSum(n1 int, n2 int) int {
	// 当执行defer， 会将defer后面的语句压入到独立的栈中（defer）（暂时不执行）
	// 当函数执行完毕， 再从defer栈， 按照先入后出的方式出栈， 执行defer语句
	defer fmt.Println("ok1 n1=", n1) // n1 = 10
	defer fmt.Println("ok2 n2=", n2) // n2 = 20

	n1++ // n1 = 11
	n2++ // n2 = 21

	res := n1 + n2
	fmt.Println("ok3 res=", res)

	// defer example
	// 预防database开了忘了关
	// connect := openDatabase()
	// defer connect.close()
	return res
}

func chap7() {
	// 使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) // 11
	// fmt.Println(f(2)) // 13
	fmt.Println(f(3))

	// defer(在函数中，程序员经常需要创建资源（数据库连接，文件句柄，锁等），为了在函数执行完毕后，及时的释放资源，Go的设计者提供defer（延时机制）)
	defSum(10, 20)

	// 值传递是value， 引用传递其中是地址拷贝 & * 指标
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func exeChap7() {
	f := makeSuffix(".jpg")
	fmt.Println(f("test"))
}

func chap8() {
	// golang using utf-8
	// slice mean array?
	str := "hello北"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	// string convert int
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("failed to convert", err)
	} else {
		fmt.Println("success to convert", n)
	}

	// int convert string
	str3 := strconv.Itoa(123)
	fmt.Printf("str=%v, str=%T\n", str3, str3)

	// string convert byte
	bytes := []byte(str)
	fmt.Printf("bytes=%v\n", bytes)

	// byte convert string
	str4 := string([]byte{97, 98, 99})
	fmt.Printf("str4=%v\n", str4)

	// decimal convert to binary
	str5 := strconv.FormatInt(123, 2)
	fmt.Printf("str5=%v\n", str5)
}

func test() {
	defer func() {
		err := recover() // recover()可以捕获异常
		if err != nil {  // 说明捕获到异常
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

func makeError(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("file not found")
	}
}

func chap9() {
	// new
	num := new(int) // return pointer *int
	fmt.Printf("type=%T, value=%v, address=%v\n", num, num, &num)

	// make 用来分配内存， 主要分配引用类型， channel，map，slice

	// golang 错误处理(panic)(defer, panic, recover)
	test()
	fmt.Println("ok")
	err := makeError("config.ini")
	if err != nil {
		// if failed, output err and close program
		panic(err)
	}
	fmt.Println("ok2")
}

func chap10() {
	// array
	chickenWeight := [...]float64{3.0, 5.0, 1.0, 3.4, 2.0, 50.0}
	for i, v := range chickenWeight {
		fmt.Printf("index=%v, value=%v\n", i, v)
	}
	fmt.Println(chickenWeight)

	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100)
	}

	fmt.Println(intArr)

	var intArr2 [5]int

	for i := 0; i < len(intArr); i++ {
		intArr2[i] = intArr[len(intArr)-(i+1)]
	}

	fmt.Println(intArr2)

	// slice
	// slice 可以看做是array的一个view， 只是指向array的一部分， 可以改变slice的长度， 但是不能改变array的长度
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice element=", slice)
	fmt.Println("slice length=", len(slice))
	fmt.Println("slice capacity=", cap(slice))
	// type slice struct {
	// 	ptr *[5]int
	// 	len int
	// 	cap int
	// }

	// slice []int = make(type, len, cap)
	var slice2 []int = make([]int, 2, 4)
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2 length=", len(slice2))
	fmt.Println("slice2 capacity=", cap(slice2))

	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println("slice3=", slice3)
	fmt.Println("slice3 length=", len(slice3))
	fmt.Println("slice3 capacity=", cap(slice3))

	slice3 = append(slice3, "jerry")
	slice3 = append(slice3, slice3...)
}

func chap11() {
	str := "hello world"
	slice := str[5:]
	fmt.Println(slice)
}

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func exeChap11() {
	fmt.Println(fbn(10))

	// sort, arrange
	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		rand.Seed(time.Now().UnixNano())
		intArr[i] = rand.Intn(100)
	}
	fmt.Println(intArr)
	// bubble sort
	for i := 0; i < len(intArr); i++ {
		for j := 0; j < len(intArr)-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				intArr[j] += intArr[j+1]
				intArr[j+1] = intArr[j] - intArr[j+1]
				intArr[j] -= intArr[j+1]
			}
		}
	}
	fmt.Println(intArr)
}

type stu struct {
	name    string
	age     int
	address string
}

func chap12() {
	// map
	// map[key]value
	// map的声明是不会分配内存的，初始化需要make， 分配内存后才能赋值和使用
	// make的作用是给map分配数据空间
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "asd"
	a["no2"] = "asd123"
	fmt.Println(a)

	heroes := make(map[string]string)
	heroes["heroesNo1"] = "asd"
	heroes["heroesNo2"] = "asd123"
	fmt.Println(heroes)

	student := make(map[string]map[string]string)
	student["sn1"] = make(map[string]string, 3)
	student["sn1"]["name"] = "tom"
	student["sn1"]["sex"] = "male"
	student["sn1"]["address"] = "jln utara"

	student["sn2"] = make(map[string]string, 3)
	student["sn2"]["name"] = "jenny"
	student["sn2"]["sex"] = "female"
	student["sn2"]["address"] = "mukim bakri"
	delete(student["sn2"], "address")
	fmt.Println(student)

	// searching in map
	val, found := student["sn1"]
	if found {
		fmt.Println(val)
	} else {
		fmt.Println("Record not found")
	}

	// looping in map
	for k, v := range student {
		fmt.Printf("key=%v, value=%v\n", k, v)
		for k2, v2 := range v {
			fmt.Printf("key2=%v, value2=%v\n", k2, v2)
		}
	}

	// sort map key
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	var keys []int
	// k, v
	for k := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, ": ", map1[k])
	}

	// remove all map key
	student = make(map[string]map[string]string)

	students := make(map[string]stu, 10)
	students["sn1"] = stu{"tom", 20, "jln utara"}
	students["sn2"] = stu{"jenny", 21, "mukim bakri"}
	students["sn3"] = stu{"jerry", 22, "jln utara"}
	fmt.Println(students)

}

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func (c Cat) Say() {
	fmt.Println("My name is", c.Name)
}

func chap13() {
	// create one cat
	cat1 := Cat{"tom", 10, "black", "eat"}
	fmt.Println(cat1)
	cat1.Say()

	cat2 := Cat{
		Name:  "jenny",
		Age:   11,
		Color: "white",
		Hobby: "play",
	}

	fmt.Println(cat2)

	// Say() wrong
}

type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone start")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera start")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func chap14() {
	// 让phone实现usb接口的方法
	// 先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(phone)
	computer.Working(camera)
}
