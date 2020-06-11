package main

import (
	"fmt"
	"time"

	gpio ".."
)

func main() {
	//	runtime.SetCPUProfileRate(1000 * 1000 * 1000)
	model, err := gpio.GetPresetDefs()
	if err != nil {
		panic(err)
	}

	pwmpins := model.GetPWMPins()
	pwm32, err := gpio.CreatePWMPin(pwmpins[0])
	if err != nil {
		if !gpio.IsAlreadyExported(err) {
			panic(err)
		} else {
			pwm32.Close()
			pwm32, err = gpio.CreatePWMPin(pwmpins[0])
			if err != nil {
				panic(err)
			}
		}
	}
	defer pwm32.Close()
	err = pwm32.Enable()
	if err != nil {
		fmt.Println(err)
		return
	}
	//min freqency is 25Hz max freqency is 160,000Hz
	for i := 25; i < 1500; i += 25 {

		err = pwm32.SetFrequency(uint(i))
		if err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(50 * time.Microsecond)

		fmt.Printf("Frequency is %d Hz\n", i)

		for j := 0; j < 1000; j++ {
			ratio := (float64(j) / 10000)
			err = pwm32.SetDutyCycleRatio(ratio)

			if err != nil {
				fmt.Println(err)
				break
			}

			time.Sleep(50 * time.Microsecond)
		}

	}

	fmt.Println("Done Goodbye")

}
func valmap(fromValue, fromLow, fromHigh, toLow, toHigh int) (toValue int) {
	return (fromValue-fromLow)*(toHigh-toLow)/(fromHigh-fromLow) + toLow
}
func valmapuint(fromValue, fromLow, fromHigh, toLow, toHigh uint) (toValue uint) {
	return (fromValue-fromLow)*(toHigh-toLow)/(fromHigh-fromLow) + toLow
}
