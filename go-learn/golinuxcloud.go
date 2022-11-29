package main

import (
	"fmt"
	"log"

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
	// fmt.Println(factorialFunction(4))
	// fmt.Println(exponentialFunction(5, 3))
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

func factorialFunction(num int) (returnNum int) {
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
	returnNum = num * factorialFunction(num-1)
	return returnNum
}

func exponentialFunction(num int, topNum int) (returnNum int) {
	for i := topNum; i > 0; i-- {
		if i == topNum {
			returnNum = num
		} else {
			returnNum *= num
		}
	}

	return returnNum
}
