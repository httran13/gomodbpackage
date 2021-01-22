package gomodbpackage

import "time"

func SayHi() string{
	return "Hello from gomodbpackage current time: "+time.Now().String()
}
