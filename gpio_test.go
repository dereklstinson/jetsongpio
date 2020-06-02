package jetsongpio

import (
	"fmt"
	"runtime"
	"testing"

	"./intern"
)

func Test_export(t *testing.T) {
	runtime.LockOSThread()
	const GPIO = uint32(79)
	var err error
	err = intern.Export(GPIO)
	defer intern.Unexport(GPIO)
	if err != nil {
		t.Error(err)
	}

	//err = intern.Setdirection(GPIO, false)
	//if err != nil {
	//	t.Error(err)
	//}
	ishigh, err := intern.Getvalue(GPIO)
	if err != nil {
		t.Error(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}
	err = intern.Setvalue(GPIO, true)
	if err != nil {
		t.Error(err)
	}
	ishigh, err = intern.Getvalue(GPIO)
	if err != nil {
		t.Error(err)
	}
	if ishigh {
		fmt.Printf("GPIO %d is High\n", GPIO)
	} else {
		fmt.Printf("GPIO %d is LOW\n", GPIO)
	}

	/*	type args struct {
			gpio uint32
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := export(tt.args.gpio); (err != nil) != tt.wantErr {
					t.Errorf("export() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}*/
}
