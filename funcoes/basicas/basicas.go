package main

import "fmt"

func main() {
	f1()
	f2("p1", "p2")
	f3()
	resultf3, resultf4 := f3(), f4("adsas1", "asdasd2")
	fmt.Println(resultf3)
	fmt.Println(resultf4)
	f4("p1", "p2")
	f5()
}

func f1() {
	fmt.Println("f1")
}

func f2(p1 string, p2 string) {
	fmt.Println(p1, p2)
}

func f3() string {
	return "f3"
}

// parametros do mesmo tipo pode ser declarados assim
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "return1", "return2"
}
