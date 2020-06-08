package jetsongpio

import (
	"errors"
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
	p intern.PWM
}

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
		return nil, err
	}
	pwm.Settings(0, 0)
	return pwm, nil
}
func (p *PWM) Settings(periodns, dutycyclens uint) error {
	err := p.p.SetPeriod(periodns)
	if err != nil {
		return err
	}
	return p.p.SetDutyCycle(dutycyclens)
}
func (p *PWM) Enable() error {
	return p.p.Enable()
}
func (p *PWM) Disable() error {
	return p.p.Disable()
}
func (p *PWM) Close() error {
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
