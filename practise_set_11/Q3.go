package main

import "fmt"

type telephone interface {
	lift()
	disconnet()
}

type SmartPhone struct {
	phoneType string 
	lift      func()
	disconnect func()
}

func lift(s SmartPhone,pt string) {
	s.phoneType = pt;
	fmt.Printf("Lift The %v!!!",s.phoneType)
}
func disconnect(s SmartPhone,pt string) {
	s.phoneType = pt;
	fmt.Printf("Disconnect The %v!!!",s.phoneType)
}

func main() {
	smart := SmartPhone{}
	lift(smart,"SamrtPhone");
	disconnect(smart,"SamrtPhone")
}