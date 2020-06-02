package intern

import (
	"fmt"
	"os"
	"syscall"
)

const gpiodir = "/sys/class/gpio"

//Export exports pin
func Export(gpio uint32) error {
	file, err := os.OpenFile(gpiodir+"/export", os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	return nil
}

//Unexport removes the pin
func Unexport(gpio uint32) error {
	file, err := os.OpenFile(gpiodir+"/unexport", os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	return nil
}

//Setdirection sets if the pin is an input or output
func Setdirection(gpio uint32, outflag bool) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/direction", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}
	if outflag {
		file.Write([]byte("out"))
	} else {
		file.Write([]byte("in"))
	}
	return nil
}

//Setvalue sets the value for gpio
func Setvalue(gpio uint32, high bool) error {
	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}
	if high {
		file.Write([]byte("1"))
	} else {
		file.Write([]byte("0"))
	}
	return nil
}

//Getvalue gets the value for gpuio
func Getvalue(gpio uint32) (high bool, err error) {
	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return false, err
	}
	val := make([]byte, 1)

	_, err = file.Read(val)
	if err != nil {
		return false, err
	}
	if string(val) != "0" {
		high = true

	} else {
		high = false
	}
	return high, nil
}

//Setedge sets the edge of the pin
func Setedge(gpio uint32, edge Edge) error {

	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/edge", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(edge.val))

	return err
}

//Edge is used for interupt detection if the pin allows it
type Edge struct {
	val string
}

//Rising sets e to Rising and returns an Edge set of that value
func (e *Edge) Rising() Edge {
	e.val = "rising"
	return *e
}

//Falling sets e to Falling and returns an Edge set of that value
func (e *Edge) Falling() Edge {
	e.val = "falling"
	return *e
}

//None sets e to None and returns an Edge set of that value
func (e *Edge) None() Edge {
	e.val = "none"
	return *e
}

//Both sets e to Both and returns an Edge set of that value
func (e *Edge) Both() Edge {
	e.val = "both"
	return *e
}

//Fdopen opens gpio
func Fdopen(gpio uint32) (fd int, err error) {
	return syscall.Open(fmt.Sprintf("%s/gpio%d/edge", gpiodir, gpio), syscall.O_RDONLY|syscall.O_NONBLOCK, 0777)

}

//Fdclose closes gpio
func Fdclose(fd int) error {
	return syscall.Close(fd)
}
