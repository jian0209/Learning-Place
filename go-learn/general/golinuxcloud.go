package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// func GoLinuxCloudMain() {
func main() {
	// mapTutorial()
	// EnvConfigTutorial()
	// fmt.Println(FactorialFunction(4))
	// fmt.Println(ExponentialFunction(5, 3))
	// TypeAssertionFunction()
	// CastingFunction()
	// StructFunction()
	// LearnFunctionFunction()
	// FlowControlFunction()
	FileControlFunction()
}

func MapTutorial() {
	// adding value to map
	employees := make(map[string]int)
	employees["Bob"] = 1
	employees["John"] = 2
	fmt.Printf("%v\n", employees)

	// accessing value to map
	firstEmployee := employees["Bob"]
	newEmployee := employees["Mike"]
	fmt.Println(firstEmployee)
	fmt.Println(newEmployee)

	// changing value in map
	employees["Bob"] = 123
	fmt.Println(employees)

	// deleting value in map
	delete(employees, "Bob")
	fmt.Println(employees)

	// create a new map
	users := map[string]int{
		"Bob":  12,
		"Jack": 24,
	}
	// check the key is inside the map
	value, isExist := users["Jack"]
	if isExist {
		fmt.Printf("Yes the key is exist: %d\n", value)
	} else {
		fmt.Printf("Noooo the key does not exist, %v\n", isExist)
	}

	// for loop to loop in map
	for k, v := range users {
		fmt.Printf("%s is %d years old \n", k, v)
	}
}

func EnvConfigTutorial() {
	conf, err := loadConfig("./")
	if err != nil {
		log.Fatalf("can't load environment app.env: %v\n", err)
		return
	}

	fmt.Printf("-----%s------\n", "Reading Environment variables Using Viper package")
	fmt.Printf("%s = %v \n", "Application_Environment", conf.AppEnv)
	fmt.Printf(" %s = %s \n", "DB_DRIVE", conf.DBDriver)
	fmt.Printf(" %s = %s \n", "Server_Listening_Address", conf.ServerAddress)
	fmt.Printf(" %s = %s \n", "Database_User_Name", conf.DBUser)
	fmt.Printf(" %s = %s \n", "Database_User_Password", conf.DBPass)
}

func loadConfig(path string) (config Config, err error) {

	// read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func FactorialFunction(num int) (returnNum int) {
	// first method
	// for i := num; i > 0; i-- {
	// 	if i == num {
	// 		returnNum = num
	// 	} else {
	// 		returnNum *= i
	// 	}
	// }

	// second method
	if num == 0 {
		return 1
	}
	returnNum = num * FactorialFunction(num-1)
	return returnNum
}

func ExponentialFunction(num int, topNum int) (returnNum int) {
	for i := topNum; i > 0; i-- {
		if i == topNum {
			returnNum = num
		} else {
			returnNum *= num
		}
	}

	return returnNum
}

func TypeAssertionFunction() {
	// an interface which has string
	var checkInterface interface{} = "Something here"

	// assigning value of interface type of checkType variable
	checkType := checkInterface.(string)
	// printing the concrete value
	fmt.Println(checkType)

	// checkTypeInt := checkInterface.(int)
	// // fatal/panic here, because checkInterface don't have int type of value
	// fmt.Println(checkTypeInt)

	// check correctly with exist type
	checkTypeInt, ok := checkInterface.(int)
	if ok {
		fmt.Println(checkTypeInt)
	} else {
		fmt.Printf("Don't have this type of assertion, got type %T \n", checkInterface)
	}
}

func CastingFunction() {
	// Converting a type into another type is an operation called casting
	type N [2]int
	var _ = N{1, 2}

	var a = 3.14       // this is float64
	var _ int = int(a) // numerical types can be casted, in this case will be rounded to 3

	// var i interface{} = "hello" // a new empty interface that contains a string
	// _, _ = i.(int)              // ok will be false (the second empty)
	// var _ = i.(int)             // this will be panic
	// _, _ = i.(string)           // ok will be true (the second empty)

	// convert string to int using strconv.atoi
	// convert string from os.Arg[] to int
	dataInput := os.Args[0]
	value, err := strconv.Atoi(dataInput)
	if err != nil {
		fmt.Printf("This %v was encountered converting string to integer\n", value)
	}
	fmt.Printf("Value entered is %s of type %T to be converted to data type of %T the value is %02d.\n", dataInput, dataInput, value, value)

	// convert int to string using strconv.itoa
	fmt.Println("Enter any number")
	var inputValue int
	_, err = fmt.Scanf("%d", &inputValue)
	if err != nil {
		fmt.Println(err)
	}
	value2 := strconv.Itoa(inputValue)
	fmt.Printf("Value entered is %02d of %T data type to be converted to %T data type the value is %s\n", inputValue, inputValue, value2, value2)
}

func StructFunction() {
	type contact struct {
		phone  int
		email  string
		street string
	}
	type user struct {
		firstName      string
		lastName       string
		age            int
		isAwesome      bool
		height         float64
		contactDetails contact
	}

	user1 := user{
		firstName: "Jack",
		lastName:  "Boh",
		age:       10,
		isAwesome: false,
		height:    12.3,
		contactDetails: contact{
			phone:  12345,
			email:  "world@gmail.cmo",
			street: "123 street",
		}}
	fmt.Println("First Name: ", user1.firstName)
	fmt.Println("Last Name: ", user1.lastName)
	fmt.Println("Age: ", user1.age)
	fmt.Println("Is Awesome: ", user1.isAwesome)
	fmt.Println("Height: ", user1.height)
	fmt.Println("phone: ", user1.contactDetails.phone)
	fmt.Println("email: ", user1.contactDetails.email)

	// anonymous struct
	person := struct {
		name, nationality string
		age               int
		score             float64
	}{
		name:        "Job",
		nationality: "SmM",
		age:         34,
		score:       20.34,
	}
	fmt.Println(person)
}

func LearnFunctionFunction() {
	res := make(chan string)
	go sendMessage(res)
	// check if the after execution value of the channel
	fmt.Printf("Response of channel: %v\n", <-res)
	printTechStack("Go", "Gin", "gRPC", "Docker", "Kubernetes")

	// sort function
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Amit", "Kumar", 42},
		{"Ram", "Singh", 18},
		{"Tulip", "Sharma", 51},
	}
	fmt.Println(people)
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
}

func sendMessage(msg chan string) {
	msg <- "Hello From GoLinux Academy"
}

func printTechStack(name ...string) {
	for _, techName := range name {
		fmt.Printf("===> %s\n", techName)
	}
}

func FlowControlFunction() {
	product := "Hello WOrld"
	for index, character := range product {
		switch character {
		case 'o', 'O':
			if character == 'o' {
				fmt.Println("Lowercase o at position ", index)
				break
			}
			fmt.Println("Uppercase O at position ", index)
		case 'd':
			fmt.Println("d at position ", index)
		}
	}

	// regexp
	asString := ""
	// check if string contains negative number
	var negative = regexp.MustCompile(`-`)
	// check if string contains floating number
	var floatingPoint = regexp.MustCompile(`\d?\.\d`)
	// check if string contains and email address
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative Number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating Point")
	case email.MatchString(asString):
		fmt.Println("It is an email")
	default:
		fmt.Println("Something else!")
	}

	words := []string{
		"hi",
		"hello",
		"friend",
	}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 3:
			fmt.Println(word, " is a short word")
		case wordLen > 5:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, " is exactly the right length")
		}
	}
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Phone     int
	}
	u := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@gmail.com",
		Phone:     1234567,
	}
	values := reflect.ValueOf(u)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Printf("Field: %s\t Value: %v\n", typesOf.Field(i).Name, values.Field(i).Interface())
	}

	// go loop in map
	user := map[string]interface{}{
		"FirstName": "John",
		"LastName":  "Doe",
		"Country":   "Go Land",
		"Email":     "johndoe@gmail.com",
	}
	for k, v := range user {
		fmt.Printf("Key: %s === Value: %s\n", k, v)
	}

	// go loop in interface
	items := []interface{}{
		1,
		"Something here",
		3.14,
	}
	for _, item := range items {
		switch item.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case float64:
			fmt.Println("float64")
		default:
			fmt.Println("unknown")
		}
	}

	// go loop in chan
	c := make(chan string, 10)
	go testChannel(c)
	for i := range c {
		println(i)
	}
}

func testChannel(c chan string) {
	users := []string{"user1", "user2", "user3", "user4", "user5"}
	for _, user := range users {
		c <- user
	}
	close(c)
}

func FileControlFunction() {
	// create file
	// file, err := os.Create("./heihei.txt")
	// if err != nil {
	// 	fmt.Println("create file error: ", err)
	// 	return
	// }
	// defer file.Close()

	// open file
	// O_RDONLY : Open a file for read only operations
	// O_WRONLY : Open a file for write only operations
	// O_RDWR : Open a file for read-write
	// O_APPEND :It appends data to the file when writing
	// O_CREATE: It creates a file if none exists.
	openFile, err := os.OpenFile("heihei.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	defer openFile.Close()

	// write file
	// b := []byte("Hello World!")
	// _, err = openFile.Write(b)
	// if err != nil {
	// 	fmt.Println("Write file error: ", err)
	// 	return
	// }
	// b2 := []byte("Log file write!!")
	// _, err = openFile.Write(b2)
	// if err != nil {
	// 	fmt.Println("Write file error: ", err)
	// }
	b := []byte("Hello World!!!\n")
	// err = ioutil.WriteFile("heihei.txt", b, 0644)
	openFile.Write(b)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// read file
	readFile, err := ioutil.ReadFile("heihei.txt")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(string(readFile))
}
