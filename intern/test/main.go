package main

import (
	"fmt"

	intern ".."
)

func main() {

	const GPIO = uint32(79)
	var err error
	err = intern.Export(GPIO)
	if err != nil {
		fmt.Println(err)
	}
	defer intern.Unexport(GPIO)
	err = intern.Setdirection(GPIO, true)
	if err != nil {
		fmt.Println(err)
	}

	fd, err := intern.Fdopen(GPIO)
	if err != nil {
		fmt.Println(err)
	}
	defer intern.Fdclose(fd)
	ishigh, err := intern.Getvalue(GPIO)
	if err != nil {
		fmt.Println(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}
	err = intern.Setvalue(GPIO, true)
	if err != nil {
		fmt.Println(err)
	}
	ishigh, err = intern.Getvalue(GPIO)
	if err != nil {
		fmt.Println(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}
}
