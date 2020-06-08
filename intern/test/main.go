package main

import (
	"fmt"
	"syscall"

	intern ".."
)

/*
func main() {
	const GPIO = uint32(79)
	var err error
	err = intern.Export(GPIO)
	if err != nil {
		fmt.Println(err)
	}
}
*/
func main() {

	const GPIO = uint32(79)
	var err error

	groups, err := syscall.Getgroups()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GroupID", groups)
	pin79 := intern.CreateGPIO(79, "")
	err = pin79.Export()

	if err != nil {
		fmt.Println(err)
	} else {
		defer pin79.Unexport()
	}

	err = pin79.SetDirection(true)

	if err != nil {
		fmt.Println(err)
	}
	err = pin79.Enable()

	if err != nil {
		fmt.Println(err)
	}
	defer pin79.Disable()

	ishigh, err := pin79.GetValue()

	if err != nil {
		fmt.Println(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}
	err = pin79.SetValue(true)
	if err != nil {
		fmt.Println(err)
	}
	ishigh, err = pin79.GetValue()
	if err != nil {
		fmt.Println(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}
}
