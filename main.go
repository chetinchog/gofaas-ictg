package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func main() {
	greetDude := Hello("")
	fmt.Println(aurora.Yellow(greetDude))
	greetChe := Hello("Che")
	fmt.Println(aurora.Yellow(greetChe))
}
