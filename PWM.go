package jetsongpio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	intern "./intern"
)

//PWMdefer contains pwm pin information.  If the pin doesn't have pwm support then
//PWMsysDirectory() returns "" and PWMid() returns -1 then that pin doesn't have PWM capabilities.
//
type PWMdefer interface {
	PWMsysDirectory() string
	PWMid() int
}

//PWM is a pin that outputs PWM signals
type PWM struct {
	p      intern.PWM
	period uint
}

//IsAlreadyExported checks to see if the pin is already exported.
func IsAlreadyExported(err error) bool {
	if err.Error() == "Warning: Pin Already Exported" {
		return true
	}
	return false
}

//CreatePWMPin will create and export the pin.  If already exported pin
//pin will still be availble.  It will return an error "Warning: Pin Already Exported".
//It is up to the developer what to do with that.
func CreatePWMPin(p PWMdefer) (pwm *PWM, err error) {
	if p.PWMid() < 0 {
		return nil, errors.New("INVALID PIN")
	}
	pwm = new(PWM)
	path, err := getpwmdirs(p)
	if err != nil {
		return nil, err
	}
	pwm.p = intern.CreatePWM(uint(p.PWMid()), path)
	err = pwm.p.Export()
	if err != nil {
		if err.Error() == "device or resource busy" {
			err = fmt.Errorf("Warning: Pin Already Exported")
			return
		} else {
			return nil, err
		}
	}

	return pwm, err
}

//SetFrequency sets the period of the pwm
//
//For the nano the frequency is between 25Hz up to 160,000Hz
//or a period of 40,000,000 ns down to 6,250ns
//
//SetFrequency auto sets duty cycle to 0.  So, the duty cycle needs to be set again.
func (p *PWM) SetFrequency(Hz uint) error {
	err := p.SetDutyTime(0)
	if err != nil {
		return err
	}
	p.period = 1e9 / Hz
	return p.p.SetPeriod(p.period)

}

//SetPeriod sets the period it is the inverse of set freqency
func (p *PWM) SetPeriod(ns uint) error {
	p.period = ns
	return p.p.SetPeriod(p.period)
}

//GetPeriod gets the period fo the frequency.
func (p *PWM) GetPeriod() (ns uint) {
	var err error
	if p.period == 0 {
		p.period, err = p.p.GetPeriod()
		if err != nil {
			fmt.Println(err)
		}
	}
	return p.period
}

//SetDutyTime sets the duty cycle to a certain amount of time.  Must be less
//than the period of the signal.
func (p *PWM) SetDutyTime(ns uint) error {
	period := p.GetPeriod()
	if ns > period {
		return fmt.Errorf("DutyTimeLength (%d) larger than period (%d)", ns, period)
	}

	return p.p.SetDutyCycle(ns)
}

//SetDutyCycleRatio will set the duty cycle a ratio of on/period.
//Ratio can't be larger than 1 or negative.
func (p *PWM) SetDutyCycleRatio(ratio float64) error {
	if ratio <= 1 && ratio >= 0 {

		return p.p.SetDutyCycle((uint)((float64)(p.period) * ratio))
	}
	return errors.New("Unsupported Ratio.  Ratio Needs to be <=1 and >=0")
}

//Enable enables the pwm.
func (p *PWM) Enable() error {
	return p.p.Enable()
}

//Disable disables the pwm.
func (p *PWM) Disable() error {

	return p.p.Disable()
}
func (p *PWM) Close() error {
	p.p.SetDutyCycle(0)
	//p.p.SetPeriod(0)
	p.p.Disable()
	return p.p.Unexport()
}
func getpwmdirs(p PWMdefer) (path string, err error) {
	path = p.PWMsysDirectory()
	if strings.Contains(path, "/pwm/pwmchip") {
		return path, nil
	}
	path = path + "/pwm"
	dircomponents, err := ioutil.ReadDir(path)
	for _, subdir := range dircomponents {

		if strings.Contains(subdir.Name(), "pwmchip") {
			path = path + "/" + subdir.Name()
			return path, nil
		}
	}
	return "", errors.New("INVALID PATH")
}
