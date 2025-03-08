package main

import (
	"fmt"
	teller "my-study-go/teller"
	"os"
	"strings"
)

const master string = "Hugo"

var version int = 1

/*
	함수 밖에서는 := 사용 불가
	Go 언어에서 이렇게 규칙을 정한 이유는 패키지 레벨 코드의 명확성과 일관성을 유지하기 위해서이다.
	:= 연산자는 함수 내에서 코드를 간결하게 작성할 수 있도록 도와주지만
	패키지 레벨에서는 더 명시적인 var 선언을 요구한다
*/

func typeTest() {
	name := "Tester"
	age := 20

	fmt.Println("name:", name)
	fmt.Println("age:", age)
}

func multiply(a, b int) int {
	return a * b
}

func getNumberAndStringByName(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 반환변수 이름을 지정할 수 있다
func getNumberAndStringByName2(name string) (length int, upperCasedName string) {
	length = len(name)
	upperCasedName = strings.ToUpper(name)
	// 반환변수 이름을 지정하면 함수 내에서 변수를 생략할 수 있다. 이것을 naked return 이라고 한다.
	return
}

func printNames(names ...string) {
	// defer 는 함수가 종료되기 전에 실행된다.
	defer fmt.Println("printNames 함수 종료")
	fmt.Println(names)
}

func mergeNames(names ...string) string {
	mergedName := ""

	// 루프 사용법 1
	for index, name := range names {
		fmt.Println("index:", index, "name:", name)
		mergedName += name + " "
	}

	// 루프 사용법 2
	// for index :=0; index < len(names); index++{
	// 	mergedName += names[index] + " "
	// }

	return mergedName
}

func checkAdult(age int) bool {
	// 조건문에서 변수 선언 가능 이 경우 조건문 내에서만 사용 가능
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func checkMaster(name string) bool {
	switch name {
	case master:
		fmt.Println("Hello Master")
		return true
	case "Tester":
		fmt.Println("Hi Tester")
		return false
	default:
		fmt.Println("Who are you? ", name)
		return false
	}

	// 표현식이 없는 switch 사용할 때 조건문을 case 문에 넣을 수 있다.
	// switch {
	// case name == master:
	// 	fmt.Println("Hello Master")
	// 	return true
	// case name == "Tester":
	// 	fmt.Println("Hi Tester")
	// 	return false
	// default:
	// 	fmt.Println("Who are you? ", name)
	// 	return false
	// }
}

func main() {
	teller.SayHello()
	teller.Say(master)
	teller.SayGoodbye()
	fmt.Println("version:", version)
	typeTest()

	fmt.Println("multiply 2 * 3 : ", multiply(2, 3))

	number, string := getNumberAndStringByName(master)

	fmt.Println("getNumberAndStringByName")
	fmt.Println("number:", number)
	fmt.Println("string:", string)

	number2, string2 := getNumberAndStringByName2(master)
	fmt.Println("getNumberAndStringByName2")
	fmt.Println("number2:", number2)
	fmt.Println("string2:", string2)

	printNames("Hugo", "Tester", "Master")

	mergedName := mergeNames("Hugo", "Tester", "Master")
	fmt.Println("mergeNames")
	fmt.Println("mergedName:", mergedName)

	adult := checkAdult(20)
	fmt.Println("checkAdult, 20살은 성인인가?", adult)

	adult = checkAdult(15)
	fmt.Println("checkAdult, 15살은 성인인가?", adult)

	master := checkMaster(master)
	fmt.Println("checkMaster, ", master)

	master = checkMaster("Tester")
	fmt.Println("checkMaster, ", master)

	master = checkMaster("John")
	fmt.Println("checkMaster, ", master)

	os.Exit(0)
}
