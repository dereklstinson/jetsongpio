package intern

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

//var debuggpio = true

const gpiodirectory = "/sys/class/gpio"

//func init() {
//	err := syscall.Setfsgid(999)
//	if err != nil {
//		panic(err)
//	}

//}

//GPIO is a gpio pin
type GPIO struct {
	fd        int  //file descriptor
	p         uint //linux pin number
	directory string
	enabled   bool
}

//CreateGPIO creates a gpio usually the directory is in /sys/class/gpio.  If directory is passed as "" then methods will use
// /sys/class/gpio.
func CreateGPIO(linuxpinnum uint, directory string) GPIO {
	if directory == "" {
		directory = gpiodirectory
	}
	return GPIO{
		p:         linuxpinnum,
		directory: directory,
	}
}

func (g GPIO) String() string {
	return fmt.Sprintf("%d", g.p)
}

var exportsleeptime = time.Duration(20) * time.Millisecond

//SetExportSleepTimeMS when doing export it takes time for the os to export the file information.  So, there needs to be a little bit
//of sleep time so the os can catch up.  The default is set to 20 ms but this function allows you to change the time if need be.
func SetExportSleepTimeMS(nmilliseconds uint64) {
	exportsleeptime = time.Duration(nmilliseconds) * time.Millisecond
}

//Export exports pin
func (g GPIO) Export() error {

	file, err := syscall.Open(g.directory+"/export", os.O_WRONLY, 777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(g.directory+"/export", os.O_WRONLY, os.ModePerm)
	//	defer file.Close()
	if err != nil {
		return err
	}
	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", g.p)))

	//_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	//syscall.Getuid()
	time.Sleep(exportsleeptime)
	return nil
}

//Unexport removes the pin
func (g GPIO) Unexport() error {

	file, err := syscall.Open(g.directory+"/unexport", os.O_WRONLY, 777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(g.directory+"/unexport", os.O_WRONLY, os.ModePerm)
	//defer file.Close()
	if err != nil {
		return err
	}

	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", g.p)))
	//	_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	return nil
}

//SetDirection sets if the pin is an input or output
func (g GPIO) SetDirection(outflag bool) error {

	//file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/direction", g.directory, gpio), os.O_WRONLY, os.ModePerm)
	//defer file.Close()

	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/direction", g.directory, g.p), os.O_WRONLY, 0777)
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

//SetValue sets the value for gpio true sets it heigh, false sets it low
func (g GPIO) SetValue(high bool) error {
	if !g.enabled {
		return fmt.Errorf("pin %v needs to be enabled", g)
	}
	//	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", g.directory, gpio), os.O_WRONLY, os.ModePerm)
	//	defer file.Close()
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/value", g.directory, g.p), os.O_WRONLY, 0777)
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

//Getvalue gets the value for gpuio if high is true then it is high. if false then low.
func (g GPIO) GetValue() (high bool, err error) {
	if !g.enabled {
		return false, fmt.Errorf("pin %v needs to be enabled", g)
	}
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/value", g.directory, g.p), os.O_RDONLY, 0777)
	defer syscall.Close(file)
	//	file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/value", g.directory, gpio), os.O_RDONLY, os.ModePerm)
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
func (g GPIO) SetEdge(edge Edge) error {
	file, err := syscall.Open(fmt.Sprintf("%s/gpio%d/edge", g.directory, g.p), os.O_WRONLY, 0777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(fmt.Sprintf("%s/gpio%d/edge", g.directory, gpio), os.O_WRONLY, os.ModePerm)
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

//Enable enables the pin so that programs can read the input or set the output.
func (g *GPIO) Enable() (err error) {
	g.fd, err = syscall.Open(fmt.Sprintf("%s/gpio%d/edge", g.directory, g.p), syscall.O_RDONLY|syscall.O_NONBLOCK, 0777)
	if err == nil {
		g.enabled = true
	}
	return err
}

//Disable disables the pin from being read or maybe outputed.
func (g *GPIO) Disable() error {
	if g.enabled != true {
		return fmt.Errorf("pin %v not enabled", g)
	}
	g.enabled = false
	return syscall.Close(g.fd)
}
