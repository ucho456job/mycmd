package util

import "fmt"

func BlueFont(str string) string {
	return "\033[34m" + str + "\033[0m"
}

func GreenFont(str string) string {
	return "\033[32m" + str + "\033[0m"
}

func YellowFont(str string) string {
	return "\033[33m" + str + "\033[0m"
}

func RedFont(str string) string {
	return "\033[31m" + str + "\033[0m"
}

func PrintWarnMsg(msg string) {
	fmt.Println(YellowFont(msg))
}

func PrintErrMsg(msg string, err error) {
	fmt.Printf(RedFont("error: "+msg+": %v\n"), err)
}
