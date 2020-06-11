package intern

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

//const p.pwmdirectory = "sys/devices/7000a000.pwm/pwm/pwmchip0"

//PWM is the pwm pin
type PWM struct {
	p            uint
	pwmdirectory string
}

//CreatePWM creates a pwm using the pwm pin value (not the gpio value) and the pwmdirectory
//Since at current time this can only be tested with the jetson nano the full path needs to be placed.
//pindefs given don't have the full path (at least for the nano so it gives /sys/devices/7000a000.pwm",
//but full path is "/sys/devices/7000a000.pwm/pwm/pwmchip0").
//
// pwm directory can also be found at /sys/class/pwm/pwmchipn where pwmchip(n) can be 0 or something else.
//
func CreatePWM(p uint, pwmdirectory string) PWM {

	return PWM{
		p:            p,
		pwmdirectory: pwmdirectory,
	}
}

//Export exports pin
func (p PWM) Export() error {

	file, err := syscall.Open(p.pwmdirectory+"/export", os.O_WRONLY, 777)
	defer syscall.Close(file)
	if err != nil {
		return err
	}
	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", p.p)))

	//_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	//syscall.Getuid()
	time.Sleep(exportsleeptime)
	return nil
}

//Unexport removes the pin
func (p PWM) Unexport() error {

	file, err := syscall.Open(p.pwmdirectory+"/unexport", os.O_WRONLY, 777)
	defer syscall.Close(file)
	//file, err := os.OpenFile(gpiodir+"/unexport", os.O_WRONLY, os.ModePerm)
	//defer file.Close()
	if err != nil {
		return err
	}

	_, err = syscall.Write(file, []byte(fmt.Sprintf("%d", p.p)))
	//	_, err = fmt.Fprintf(file, "%d", gpio)
	if err != nil {
		return err
	}
	return nil
}

//SetPeriod sets the period in ns (nano seconds)
func (p PWM) SetPeriod(ns uint) error {
	file, err := syscall.Open(fmt.Sprintf("%s/pwm%d/period", p.pwmdirectory, p.p), os.O_WRONLY, 0777)
	defer syscall.Close(file)

	_, err = syscall.Write(file, []byte(fmt.Sprint(ns)))
	return err
}

//GetPeriod returns the period.
//Sometimes the period is already set when it is exported.  You might want to get that
//info before doing anything.
func (p PWM) GetPeriod() (ns uint, err error) {
	file, err := syscall.Open(fmt.Sprintf("%s/pwm%d/period", p.pwmdirectory, p.p), os.O_RDONLY, 0777)
	defer syscall.Close(file)
	buffer := make([]byte, 128)

	n, err := syscall.Read(file, buffer)
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(string(bytes.TrimSpace(buffer[:n])))
	if err != nil {
		return 0, err
	}
	return uint(val), err

}

//SetDutyCycle sets the duty cycle in ns (nano seconds)
func (p PWM) SetDutyCycle(ns uint) error {
	file, err := syscall.Open(fmt.Sprintf("%s/pwm%d/duty_cycle", p.pwmdirectory, p.p), os.O_WRONLY, 0777)
	defer syscall.Close(file)

	_, err = syscall.Write(file, []byte(fmt.Sprint(ns)))
	return err
}

//Enable enables the PWM
func (p PWM) Enable() error {
	file, err := syscall.Open(fmt.Sprintf("%s/pwm%d/enable", p.pwmdirectory, p.p), os.O_WRONLY, 0777)
	defer syscall.Close(file)

	_, err = syscall.Write(file, []byte("1"))
	return err
}

//Disable disables the PWM
func (p PWM) Disable() error {
	file, err := syscall.Open(fmt.Sprintf("%s/pwm%d/enable", p.pwmdirectory, p.p), os.O_WRONLY, 0777)
	defer syscall.Close(file)

	_, err = syscall.Write(file, []byte("0"))
	return err
}
