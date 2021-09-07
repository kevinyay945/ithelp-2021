package main

import (
	"first-golang-project/util"
	"fmt"
)

func main() {
	util.PrintHelloWorld()
	fmt.Println(util.ExportedVar)
	fmt.Println(util.Var1.ExportedProperty)
	// fmt.Println(util.Var1.unexportedProperty) // util.Var1.unexportedProperty undefined (cannot refer to unexported field or method unexportedProperty)
	// fmt.Println(util.unexportedVar) // cannot refer to unexported name util.unexportedVar
}
