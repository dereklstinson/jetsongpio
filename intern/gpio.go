package intern

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

//var debuggpio = true

const gpiodir = "/sys/class/gpio"

//func init() {
//	err := syscall.Setfsgid(999)
//	if err != nil {
//		panic(err)
//	}

//}

var exportsleeptime = time.Duration(20) * time.Millisecond

//SetExportSleepTimeMS when doing export it takes time for the os to export the file information.  So, there needs to be a little bit
//of sleep time so the os can catch up.  The default is set to 20 ms but this function allows you to change the time if need be.
func SetExportSleepTimeMS(nmilliseconds uint64) {
	exportsleeptime = time.Duration(nmilliseconds) * time.Millisecond
}

//Export exports pin
func Export(gpio uint32) error {

	file, err := syscall.Open(gpiodir+"/export", os.O_WRONLY, 777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(gpiodir+"/export", os.O_WRONLY, os.ModePerm)
	//	defer file.Close()
	if err != nil {
		return err
	}
	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", gpio)))

	//_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	//syscall.Getuid()
	time.Sleep(exportsleeptime)
	return nil
}

//Unexport removes the pin
func Unexport(gpio uint32) error {

	file, err := syscall.Open(gpiodir+"/unexport", os.O_WRONLY, 777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(gpiodir+"/unexport", os.O_WRONLY, os.ModePerm)
	//defer file.Close()
	if err != nil {
		return err
	}

	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", gpio)))
	//	_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	return nil
}

//Setdirection sets if the pin is an input or output
func Setdirection(gpio uint32, outflag bool) error {

	//file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/direction", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	//defer file.Close()

	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/direction", gpiodir, gpio), os.O_WRONLY, 0777)
	defer syscall.Close(file)
	if err != nil {
		fmt.Println("Error In opening file direction")
		return err
	}
	if outflag {
		_, err = syscall.Write(file, []byte("out"))
		//_, err = file.Write([]byte("out"))

	} else {
		_, err = syscall.Write(file, []byte("in"))

		//_, err = file.Write([]byte("in"))

	}
	return err
}

//Setvalue sets the value for gpio
func Setvalue(gpio uint32, high bool) error {
	//	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	//	defer file.Close()
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_WRONLY, 0777)
	defer syscall.Close(file)
	if err != nil {
		fmt.Println("Set Value: Error in opening file ")
		return err
	}
	if high {
		_, err = syscall.Write(file, []byte("1"))
		//	_, err = file.Write([]byte("1"))
		if err != nil {
			return err
		}
	} else {
		_, err = syscall.Write(file, []byte("0"))

		//	_, err = file.Write([]byte("0"))
		if err != nil {
			return err
		}
	}
	return nil
}

//Getvalue gets the value for gpuio
func Getvalue(gpio uint32) (high bool, err error) {
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_RDONLY, 0777)
	defer syscall.Close(file)
	//	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", gpiodir, gpio), os.O_RDONLY, os.ModePerm)
	//	defer file.Close()
	if err != nil {
		return false, err
	}
	val := make([]byte, 1)
	_, err = syscall.Read(file, val)
	//_, err = file.Read(val)
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
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/edge", gpiodir, gpio), os.O_WRONLY, 0777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/edge", gpiodir, gpio), os.O_WRONLY, os.ModePerm)
	//defer file.Close()
	if err != nil {
		return err
	}
	_, err = syscall.Write(file, []byte(edge.val))
	//_, err = file.Write([]byte(edge.val))

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
