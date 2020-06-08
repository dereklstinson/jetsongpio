package jetsongpio

import (
	"fmt"
	"testing"
)

func TestGetPresetDefs(t *testing.T) {
	JetsonModel, err := GetPresetDefs()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(JetsonModel.Model())

	for i := 1; i <= 40; i++ {
		bpif, err := JetsonModel.GetBoardPins(i)
		if err != nil {
			t.Error(err)
		}
		fmt.Println("BP: ", i)
		fmt.Println(bpif)
	}
}
