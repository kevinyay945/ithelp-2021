package util

var ExportedVar string = "GLOBAL VAR"
var unexportedVar string = "unexported var"

type ExportedStruct struct {
	ExportedProperty   string
	unexportedProperty string
}

var Var1 ExportedStruct = ExportedStruct{
	ExportedProperty:   "exported",
	unexportedProperty: "unexported",
}
