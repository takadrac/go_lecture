// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

//import "fmt"
/*
const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}
*/
/*
func getOsName() string {
	return "mac"
}
*/

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func thirdPartyConnectDB() {
	/* panic & recover */

	panic("Unable to connect database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	/*
		fmt.Println(Pi, Username, Password)
		t, f := true, false
		fmt.Printf("%T %v %t\n", t, 1, t)
		fmt.Printf("%T %v %t\n", f, 1, f)

		var s string = "14"
		i, _ := strconv.Atoi(s)
		fmt.Printf("%T %v %d\n", i, 1, i)
	*/

	/* if statement */
	num := 5
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	/* for statement */
	for i := 0; i < 10; i++ {

		if i == 3 {
			fmt.Println("continue")
			continue //return to for statement
		}

		if i > 5 {
			fmt.Println("break")
			break //finish for statement
		}
		fmt.Println(i)
	}
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
	/*
		for {
			fmt.Println("hello")
		}
	*/

	/* range statement */
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	/* swich statement */
	//os := getOsName() //"mac"
	/*
		switch os := getOsName(); os {
		case "mac":
			fmt.Println("Mac!!")
		case "windows":
			fmt.Println("Windows!!")
		default:
			fmt.Println("Default!!")
		}
	*/
	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	}

	/*defer statement*/
	//defer fmt.Println("world") //execute at the end of function

	//fmt.Println("hello")
	/*
		file, _ := os.Open("./lesson_definef.go")
		defer file.Close()
		data := make([]byte, 100)
		file.Read(data)
		fmt.Println(string(data))
	*/

	/* log statement */
	// go logging github
	log.Println("logging!")
	//log.Printf("%T %v", "test", "test")

	//log.Fatalln("error!!") //break program
	LoggingSettings("test.log")

	/* error handling */
	file, err := os.Open("./lesson_define.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error!")
	}
	fmt.Println(count, string(data))

	/* panic & recover */
	save()
	fmt.Println("OK?")
}
