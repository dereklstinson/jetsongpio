package jetsongpio

import (
	"errors"
	"fmt"

	intern "./intern"
)

//GPIOGlobalDef returns the global GPIO pin number
type GPIOGlobalDef interface {
	GPIOGlobal() (int, error)
}

//GPIOin is a gpio that reads the pin
type GPIOin struct {
	g intern.GPIO
}

//CreateGPIOin creates a gpiopin if there is an error the pin will not be usable.
func CreateGPIOin(pin GPIOGlobalDef) (g *GPIOin, err error) {
	gpin, err := pin.GPIOGlobal()
	if err != nil {
		return
	}
	g = new(GPIOin)
	g.g = intern.CreateGPIO(uint(gpin), "")
	err = g.g.Export()
	if err != nil {
		if err != nil {
			err = fmt.Errorf("Warning: Pin Already Exported")
			return
		}
	}
	err = g.g.SetDirection(false)
	if err != nil {
		err2 := g.g.Unexport()
		if err2 != nil {
			fmt.Println(err2)
		}

		return
	}
	err = g.g.Enable()
	if err != nil {
		err2 := g.g.Unexport()
		if err2 != nil {
			fmt.Println(err2)
		}
		return
	}
	return
}

//Close Disables the pin and unexports it.
func (g GPIOin) Close() error {
	var err error
	err = g.g.Disable()
	if err != nil {
		fmt.Println(err)
	}
	err = g.g.Unexport()
	if err != nil {
		fmt.Println(err)

	}
	return err
}

//GetSignal gets the signal from the GPIO pin
func (g GPIOin) GetSignal() (s Signal) {
	val, err := g.g.GetValue()
	if err != nil {
		fmt.Println("(g GPIOin)GetSignal() error: ", err)
		return 0
	}
	if val {
		return s.HIGH()
	}
	return s.LOW()

}

//GPIOout is a gpio that writes to the pin
type GPIOout struct {
	g intern.GPIO
}

//CreateGPIOout creates a GPIOout if there is an error the pin will not be usable.
func CreateGPIOout(pin GPIOGlobalDef, initial Signal) (g *GPIOout, err error) {
	flg := initial
	if initial != flg.HIGH() && initial != flg.LOW() {
		return g, errors.New("Initial Signal needs to be High or Low")
	}
	gpin, err := pin.GPIOGlobal()
	if err != nil {
		return
	}
	g = new(GPIOout)
	g.g = intern.CreateGPIO(uint(gpin), "")
	err = g.g.Export()
	if err != nil {
		err = fmt.Errorf("Warning: Pin Already Exported")
		return
	}
	err = g.g.SetDirection(true)
	if err != nil {
		err2 := g.g.Unexport()
		fmt.Println(err2)
		return
	}
	err = g.g.Enable()
	if err != nil {
		err2 := g.g.SetDirection(false)
		if err2 != nil {
			fmt.Println(err2)
		}
		err2 = g.g.Unexport()
		if err2 != nil {
			fmt.Println(err2)
		}

		return
	}

	if initial == flg.HIGH() {
		err = g.g.SetValue(true)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err = g.g.SetValue(false)
		if err != nil {
			fmt.Println(err)
		}

	}

	return
}

//Close closes the channel.  The pin cannot be used unless recalled using
func (g GPIOout) Close() error {
	err := g.SetLow()
	if err != nil {
		fmt.Println(err)
	}
	err = g.g.Disable()
	if err != nil {
		fmt.Println(err)
	}
	return g.g.Unexport()

}

//SetHigh sets the gpio out pin high
func (g GPIOout) SetHigh() error {
	return g.g.SetValue(true)
}

//SetLow sets the gpio out pin low.
func (g GPIOout) SetLow() error {
	return g.g.SetValue(false)
}

//Closer closes a pin channel
type Closer interface {
	Close() error
}

//ClosePins is an helper func it will close all the pins passed
func ClosePins(closers ...Closer) []error {
	multierror := make([]error, len(closers))
	var flagerror bool
	for i := range closers {
		err := closers[i].Close()
		if err != nil {
			flagerror = true
			multierror[i] = err
		}

	}
	if flagerror {
		return multierror
	}
	return nil
}
