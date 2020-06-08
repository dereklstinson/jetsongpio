package jetsongpio

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//PinDefer is an interface the implements all of these
type PinDefer interface {
	GPIOGlobal() (int, error)
	LinuxGPIOPinNum() int
	GPIOsysfsDirectory() string
	BoardPinNumber() int
	BCMnum() int
	CVMname() string
	TegraPinName() string
	PWMsysDirectory() string
	PWMid() int
}

//PinDef contains pin definitions
type PinDef struct {
	linuxpinnum        int
	gpiosysfsdirectory string
	boardpinnum        int
	bcmpinnum          int
	cvmpinname         string
	tegrapinname       string
	pwmsysfsdirectory  string
	pwmidinpwmchip     int
}

func (p PinDef) String() string {
	gpioglobal, err := p.GPIOGlobal()
	if err != nil {
		fmt.Println("GLOBALPIN error:", err)
		gpioglobal = -1
	}

	return fmt.Sprintf("PinDef{\nGPIOGlobal: %d\nLinuxGPIOPinNum: %d\nGPIOsysDir: %s\nBoardPinNum: %d\nBCMnum: %d\nCVMname: %s\nTegraPinName: %s\nPWMsysDir: %s\nPWMid: %d\n}\n",
		gpioglobal, p.linuxpinnum, p.gpiosysfsdirectory, p.boardpinnum, p.bcmpinnum, p.cvmpinname, p.tegrapinname, p.pwmsysfsdirectory, p.pwmidinpwmchip)
}
func (m Model) String() string {
	return fmt.Sprintf("Model{\nP1_Revision: %d\nRAM: %s\nREVISION: %s\nTYPE: %s\nMANUFACTURER: %s\nPROCESSOR: %s\n}\n", m.P1Revision, m.RAM, m.Revision, m.Type, m.Manufacturer, m.Processor)
}

//GPIOGlobal is some configuration of the base GPIO controller number
// plus LinuxGPIOPinNum().
//
//Right now I am not sure if there is a difference between LinuxGPIOPinNum() and GPIOGlobal().
//
//The base controller number
func (p PinDef) GPIOGlobal() (int, error) {
	path := p.GPIOsysfsDirectory()
	if path == "" {
		return -1, errors.New("INVALID PATH")
	}
	dircomponents, err := ioutil.ReadDir(path)
	if err != nil {
		return -1, err
	}
	path = path + "/gpio"
	for _, subdir := range dircomponents {

		if strings.Contains(subdir.Name(), "gpiochip") {

			path = path + "/" + subdir.Name() + "/base"
			file, err := os.Open(path)
			if err != nil {
				return -1, err
			}
			defer file.Close()
			filebytes, err := ioutil.ReadAll(file)
			if err != nil {
				return -1, err

			}
			baseval, err := strconv.Atoi(string(bytes.TrimSpace(filebytes)))
			if err != nil {
				return -1, err
			}
			return p.LinuxGPIOPinNum() + baseval, nil
		}

	}
	return -1, errors.New("Unsupported Definition")
}

//LinuxGPIOPinNum returns the linux pin num
func (p PinDef) LinuxGPIOPinNum() int {
	return p.linuxpinnum
}

//GPIOsysfsDirectory returns the sys file system directory
func (p PinDef) GPIOsysfsDirectory() string {
	return p.gpiosysfsdirectory
}

//BoardPinNumber returns the pin number for the developer board
func (p PinDef) BoardPinNumber() int {
	return p.boardpinnum
}

//BCMnum returns the BCM mode pin number
func (p PinDef) BCMnum() int {
	return p.bcmpinnum
}

//CVMname returns the CVM pin name
func (p PinDef) CVMname() string {
	return p.cvmpinname
}

//TegraPinName returns the TEGRA_SOC pin name
func (p PinDef) TegraPinName() string {
	return p.tegrapinname
}

//PWMsysDirectory returns the PWM chip sysfs directory
func (p PinDef) PWMsysDirectory() string {
	return p.pwmsysfsdirectory
}

//PWMid returns the ID within the PWM chip.
func (p PinDef) PWMid() int {
	return p.pwmidinpwmchip
}

/*
func copypindef(p PinDef)(c PinDef){
	c.linuxpinnum=p.linuxpinnum
	c.gpiosysfsdirectory=p.gpiosysfsdirectory
	c.boardpinnum=p.boardpinnum
	c.bcmpinnum=p.bcmpinnum
	c.cvmpinname=p.cvmpinname
	c.tegrapinname=p.tegrapinname
	c.pwmsysfsdirectory=p.pwmsysfsdirectory
	c.pwmidinpwmchip=p.pwmidinpwmchip
	return c
}
*/

func checkpluginmanager(compatsprefix ...string) error {
	const pluginpath = "/proc/device-tree/chosen/plugin-manager/ids"
	fi, err := os.Stat(pluginpath)
	if os.IsNotExist(err) {
		if !fi.IsDir() {
			return errors.New("Jetson Plugin Manager folder doesn't exist")
		}

	}
	dircomponents, err := ioutil.ReadDir(pluginpath)
	if err != nil {
		return err
	}
	for _, dirfi := range dircomponents {
		for _, prefix := range compatsprefix {
			if strings.Contains(dirfi.Name(), prefix+"-") {
				return nil
			}
		}

	}
	return errors.New("Jetson.GPIO library not been verified with this board.  Most likely these defs will not work")
}

//GetPresetDefs will scan the jetson and return the predetermined pins.  (Most useful for nvidia development boards)
//
//***Warning*** This will not work for boards newer than the NX boards and for custom boards.  If custom boards please follow the interfaces
//used in GPIO.go file and PWM.go file and to control the pins.
func GetPresetDefs() (GPIOData, error) {
	const compatiblepath = "/proc/device-tree/compatible"
	file, err := os.Open(compatiblepath)
	if err != nil {
		return GPIOData{}, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return GPIOData{}, err
	}
	//gpiodata: 0=NX,1=Xavier,2=TX2, 3=TX1,4=NANO
	for _, compat := range compatsNX {
		if strings.Contains(string(data), compat) {

			return gpiodata[0], checkpluginmanager("3509", "3449")
		}
	}
	for _, compat := range compatsXAVIER {
		if strings.Contains(string(data), compat) {
			return gpiodata[1], checkpluginmanager("2822")
		}
	}

	for _, compat := range compatsTX2 {
		if strings.Contains(string(data), compat) {
			return gpiodata[2], checkpluginmanager("2597")
		}
	}
	for _, compat := range compatsTX1 {
		if strings.Contains(string(data), compat) {
			return gpiodata[3], checkpluginmanager("2597")
		}
	}
	for _, compat := range compatsNANO {
		if strings.Contains(string(data), compat) {
			return gpiodata[4], checkpluginmanager("3449", "3448")
		}
	}

	return GPIOData{}, errors.New("Unsupported Board or version or something I dont know check the data of this")
}

/*
func init() {
	jetsonnxpindefs
}
*/

var jetsonNXpindefs = []PinDef{
	{148, "/sys/devices/2200000.gpio", 7, 4, "GPIO09", "AUD_MCLK", "", -1},
	{140, "/sys/devices/2200000.gpio", 11, 17, "UART1_RTS", "UART1_RTS", "", -1},
	{157, "/sys/devices/2200000.gpio", 12, 18, "I2S0_SCLK", "DAP5_SCLK", "", -1},
	{192, "/sys/devices/2200000.gpio", 13, 27, "SPI1_SCK", "SPI3_SCK", "", -1},
	{20, "/sys/devices/c2f0000.gpio", 15, 22, "GPIO12", "TOUCH_CLK", "", -1},
	{196, "/sys/devices/2200000.gpio", 16, 23, "SPI1_CS1", "SPI3_CS1_N", "", -1},
	{195, "/sys/devices/2200000.gpio", 18, 24, "SPI1_CS0", "SPI3_CS0_N", "", -1},
	{205, "/sys/devices/2200000.gpio", 19, 10, "SPI0_MOSI", "SPI1_MOSI", "", -1},
	{204, "/sys/devices/2200000.gpio", 21, 9, "SPI0_MISO", "SPI1_MISO", "", -1},
	{193, "/sys/devices/2200000.gpio", 22, 25, "SPI1_MISO", "SPI3_MISO", "", -1},
	{203, "/sys/devices/2200000.gpio", 23, 11, "SPI0_SCK", "SPI1_SCK", "", -1},
	{206, "/sys/devices/2200000.gpio", 24, 8, "SPI0_CS0", "SPI1_CS0_N", "", -1},
	{207, "/sys/devices/2200000.gpio", 26, 7, "SPI0_CS1", "SPI1_CS1_N", "", -1},
	{133, "/sys/devices/2200000.gpio", 29, 5, "GPIO01", "SOC_GPIO41", "", -1},
	{134, "/sys/devices/2200000.gpio", 31, 6, "GPIO11", "SOC_GPIO42", "", -1},
	{136, "/sys/devices/2200000.gpio", 32, 12, "GPIO07", "SOC_GPIO44", "/sys/devices/32f0000.pwm", 0},
	{105, "/sys/devices/2200000.gpio", 33, 13, "GPIO13", "SOC_GPIO54", "/sys/devices/3280000.pwm", 0},
	{160, "/sys/devices/2200000.gpio", 35, 19, "I2S0_FS", "DAP5_FS", "", -1},
	{141, "/sys/devices/2200000.gpio", 36, 16, "UART1_CTS", "UART1_CTS", "", -1},
	{194, "/sys/devices/2200000.gpio", 37, 26, "SPI1_MOSI", "SPI3_MOSI", "", -1},
	{159, "/sys/devices/2200000.gpio", 38, 20, "I2S0_DIN", "DAP5_DIN", "", -1},
	{158, "/sys/devices/2200000.gpio", 40, 21, "I2S0_DOUT", "DAP5_DOUT", "", -1},
}

var compatsNX = []string{
	"nvidia,p3509-0000+p3668-0000",
	"nvidia,p3509-0000+p3668-0001",
	"nvidia,p3449-0000+p3668-0000",
	"nvidia,p3449-0000+p3668-0001",
}

var jetsonXAVIERpindefs = []PinDef{
	{134, "/sys/devices/2200000.gpio", 7, 4, "MCLK05", "SOC_GPIO42", "", -1},
	{140, "/sys/devices/2200000.gpio", 11, 17, "UART1_RTS", "UART1_RTS", "", -1},
	{63, "/sys/devices/2200000.gpio", 12, 18, "I2S2_CLK", "DAP2_SCLK", "", -1},
	{136, "/sys/devices/2200000.gpio", 13, 27, "PWM01", "SOC_GPIO44", "/sys/devices/32f0000.pwm", 0},
	// Older versions of L4T don"t enable this PWM controller in DT, so this PWM
	// channel may not be available.
	{105, "/sys/devices/2200000.gpio", 15, 22, "GPIO27", "SOC_GPIO54", "/sys/devices/3280000.pwm", 0},
	{8, "/sys/devices/c2f0000.gpio", 16, 23, "GPIO8", "CAN1_STB", "", -1},
	{56, "/sys/devices/2200000.gpio", 18, 24, "GPIO35", "SOC_GPIO12", "/sys/devices/32c0000.pwm", 0},
	{205, "/sys/devices/2200000.gpio", 19, 10, "SPI1_MOSI", "SPI1_MOSI", "", -1},
	{204, "/sys/devices/2200000.gpio", 21, 9, "SPI1_MISO", "SPI1_MISO", "", -1},
	{129, "/sys/devices/2200000.gpio", 22, 25, "GPIO17", "SOC_GPIO21", "", -1},
	{203, "/sys/devices/2200000.gpio", 23, 11, "SPI1_CLK", "SPI1_SCK", "", -1},
	{206, "/sys/devices/2200000.gpio", 24, 8, "SPI1_CS0_N", "SPI1_CS0_N", "", -1},
	{207, "/sys/devices/2200000.gpio", 26, 7, "SPI1_CS1_N", "SPI1_CS1_N", "", -1},
	{3, "/sys/devices/c2f0000.gpio", 29, 5, "CAN0_DIN", "CAN0_DIN", "", -1},
	{2, "/sys/devices/c2f0000.gpio", 31, 6, "CAN0_DOUT", "CAN0_DOUT", "", -1},
	{9, "/sys/devices/c2f0000.gpio", 32, 12, "GPIO9", "CAN1_EN", "", -1},
	{0, "/sys/devices/c2f0000.gpio", 33, 13, "CAN1_DOUT", "CAN1_DOUT", "", -1},
	{66, "/sys/devices/2200000.gpio", 35, 19, "I2S2_FS", "DAP2_FS", "", -1},
	//Input-only (due to base board)
	{141, "/sys/devices/2200000.gpio", 36, 16, "UART1_CTS", "UART1_CTS", "", -1},
	{1, "/sys/devices/c2f0000.gpio", 37, 26, "CAN1_DIN", "CAN1_DIN", "", -1},
	{65, "/sys/devices/2200000.gpio", 38, 20, "I2S2_DIN", "DAP2_DIN", "", -1},
	{64, "/sys/devices/2200000.gpio", 40, 21, "I2S2_DOUT", "DAP2_DOUT", "", -1},
}
var compatsXAVIER = []string{
	"nvidia,p2972-0000",
	"nvidia,p2972-0006",
	"nvidia,jetson-xavier",
}

var jetsonTX2pindefs = []PinDef{
	{76, "/sys/devices/2200000.gpio", 7, 4, "AUDIO_MCLK", "AUD_MCLK", "", -1},
	// Output-only (due to base board)
	{146, "/sys/devices/2200000.gpio", 11, 17, "UART0_RTS", "UART1_RTS", "", -1},
	{72, "/sys/devices/2200000.gpio", 12, 18, "I2S0_CLK", "DAP1_SCLK", "", -1},
	{77, "/sys/devices/2200000.gpio", 13, 27, "GPIO20_AUD_INT", "GPIO_AUD0", "", -1},
	{15, "/sys/devices/3160000.i2c/i2c-0/0-0074", 15, 22, "GPIO_EXP_P17", "GPIO_EXP_P17", "", -1},
	// Input-only (due to module):
	{40, "/sys/devices/c2f0000.gpio", 16, 23, "AO_DMIC_IN_DAT", "CAN_GPIO0", "", -1},
	{161, "/sys/devices/2200000.gpio", 18, 24, "GPIO16_MDM_WAKE_AP", "GPIO_MDM2", "", -1},
	{109, "/sys/devices/2200000.gpio", 19, 10, "SPI1_MOSI", "GPIO_CAM6", "", -1},
	{108, "/sys/devices/2200000.gpio", 21, 9, "SPI1_MISO", "GPIO_CAM5", "", -1},
	{14, "/sys/devices/3160000.i2c/i2c-0/0-0074", 22, 25, "GPIO_EXP_P16", "GPIO_EXP_P16", "", -1},
	{107, "/sys/devices/2200000.gpio", 23, 11, "SPI1_CLK", "GPIO_CAM4", "", -1},
	{110, "/sys/devices/2200000.gpio", 24, 8, "SPI1_CS0", "GPIO_CAM7", "", -1},
	{0, "", 26, 7, "SPI1_CS1", "", "", -1},
	{78, "/sys/devices/2200000.gpio", 29, 5, "GPIO19_AUD_RST", "GPIO_AUD1", "", -1},
	{42, "/sys/devices/c2f0000.gpio", 31, 6, "GPIO9_MOTION_INT", "CAN_GPIO2", "", -1},
	// Output-only (due to module):
	{41, "/sys/devices/c2f0000.gpio", 32, 12, "AO_DMIC_IN_CLK", "CAN_GPIO1", "", -1},
	{69, "/sys/devices/2200000.gpio", 33, 13, "GPIO11_AP_WAKE_BT", "GPIO_PQ5", "", -1},
	{75, "/sys/devices/2200000.gpio", 35, 19, "I2S0_LRCLK", "DAP1_FS", "", -1},
	// Input-only (due to base board) IF NVIDIA debug card NOT plugged in
	// Output-only (due to base board) IF NVIDIA debug card plugged in
	{147, "/sys/devices/2200000.gpio", 36, 16, "UART0_CTS", "UART1_CTS", "", -1},
	{68, "/sys/devices/2200000.gpio", 37, 26, "GPIO8_ALS_PROX_INT", "GPIO_PQ4", "", -1},
	{74, "/sys/devices/2200000.gpio", 38, 20, "I2S0_SDIN", "DAP1_DIN", "", -1},
	{73, "/sys/devices/2200000.gpio", 40, 21, "I2S0_SDOUT", "DAP1_DOUT", "", -1},
}
var compatsTX2 = []string{
	"nvidia,p2771-0000",
	"nvidia,p2771-0888",
	"nvidia,p3489-0000",
	"nvidia,lightning",
	"nvidia,quill",
	"nvidia,storm",
}
var jetsonTX1pindefs = []PinDef{
	{216, "/sys/devices/6000d000.gpio", 7, 4, "AUDIO_MCLK", "AUD_MCLK", "", -1},
	// Output-only (due to base board)
	{162, "/sys/devices/6000d000.gpio", 11, 17, "UART0_RTS", "UART1_RTS", "", -1},
	{11, "/sys/devices/6000d000.gpio", 12, 18, "I2S0_CLK", "DAP1_SCLK", "", -1},
	{38, "/sys/devices/6000d000.gpio", 13, 27, "GPIO20_AUD_INT", "GPIO_PE6", "", -1},
	{15, "/sys/devices/7000c400.i2c/i2c-1/1-0074", 15, 22, "GPIO_EXP_P17", "GPIO_EXP_P17", "", -1},
	{37, "/sys/devices/6000d000.gpio", 16, 23, "AO_DMIC_IN_DAT", "DMIC3_DAT", "", -1},
	{184, "/sys/devices/6000d000.gpio", 18, 24, "GPIO16_MDM_WAKE_AP", "MODEM_WAKE_AP", "", -1},
	{16, "/sys/devices/6000d000.gpio", 19, 10, "SPI1_MOSI", "SPI1_MOSI", "", -1},
	{17, "/sys/devices/6000d000.gpio", 21, 9, "SPI1_MISO", "SPI1_MISO", "", -1},
	{14, "/sys/devices/7000c400.i2c/i2c-1/1-0074", 22, 25, "GPIO_EXP_P16", "GPIO_EXP_P16", "", -1},
	{18, "/sys/devices/6000d000.gpio", 23, 11, "SPI1_CLK", "SPI1_SCK", "", -1},
	{19, "/sys/devices/6000d000.gpio", 24, 8, "SPI1_CS0", "SPI1_CS0", "", -1},
	{20, "/sys/devices/6000d000.gpio", 26, 7, "SPI1_CS1", "SPI1_CS1", "", -1},
	{219, "/sys/devices/6000d000.gpio", 29, 5, "GPIO19_AUD_RST", "GPIO_X1_AUD", "", -1},
	{186, "/sys/devices/6000d000.gpio", 31, 6, "GPIO9_MOTION_INT", "MOTION_INT", "", -1},
	{36, "/sys/devices/6000d000.gpio", 32, 12, "AO_DMIC_IN_CLK", "DMIC3_CLK", "", -1},
	{63, "/sys/devices/6000d000.gpio", 33, 13, "GPIO11_AP_WAKE_BT", "AP_WAKE_NFC", "", -1},
	{8, "/sys/devices/6000d000.gpio", 35, 19, "I2S0_LRCLK", "DAP1_FS", "", -1},
	// Input-only (due to base board) IF NVIDIA debug card NOT plugged in
	// Input-only (due to base board) (always reads fixed value) IF NVIDIA debug card plugged in
	{163, "/sys/devices/6000d000.gpio", 36, 16, "UART0_CTS", "UART1_CTS", "", -1},
	{187, "/sys/devices/6000d000.gpio", 37, 26, "GPIO8_ALS_PROX_INT", "ALS_PROX_INT", "", -1},
	{9, "/sys/devices/6000d000.gpio", 38, 20, "I2S0_SDIN", "DAP1_DIN", "", -1},
	{10, "/sys/devices/6000d000.gpio", 40, 21, "I2S0_SDOUT", "DAP1_DOUT", "", -1},
}
var compatsTX1 = []string{
	"nvidia,p2371-2180",
	"nvidia,jetson-cv",
}
var jetsonNANOpindefs = []PinDef{
	{216, "/sys/devices/6000d000.gpio", 7, 4, "GPIO9", "AUD_MCLK", "", -1},
	{50, "/sys/devices/6000d000.gpio", 11, 17, "UART1_RTS", "UART2_RTS", "", -1},
	{79, "/sys/devices/6000d000.gpio", 12, 18, "I2S0_SCLK", "DAP4_SCLK", "", -1},
	{14, "/sys/devices/6000d000.gpio", 13, 27, "SPI1_SCK", "SPI2_SCK", "", -1},
	{194, "/sys/devices/6000d000.gpio", 15, 22, "GPIO12", "LCD_TE", "", -1},
	{232, "/sys/devices/6000d000.gpio", 16, 23, "SPI1_CS1", "SPI2_CS1", "", -1},
	{15, "/sys/devices/6000d000.gpio", 18, 24, "SPI1_CS0", "SPI2_CS0", "", -1},
	{16, "/sys/devices/6000d000.gpio", 19, 10, "SPI0_MOSI", "SPI1_MOSI", "", -1},
	{17, "/sys/devices/6000d000.gpio", 21, 9, "SPI0_MISO", "SPI1_MISO", "", -1},
	{13, "/sys/devices/6000d000.gpio", 22, 25, "SPI1_MISO", "SPI2_MISO", "", -1},
	{18, "/sys/devices/6000d000.gpio", 23, 11, "SPI0_SCK", "SPI1_SCK", "", -1},
	{19, "/sys/devices/6000d000.gpio", 24, 8, "SPI0_CS0", "SPI1_CS0", "", -1},
	{20, "/sys/devices/6000d000.gpio", 26, 7, "SPI0_CS1", "SPI1_CS1", "", -1},
	{149, "/sys/devices/6000d000.gpio", 29, 5, "GPIO01", "CAM_AF_EN", "", -1},
	{200, "/sys/devices/6000d000.gpio", 31, 6, "GPIO11", "GPIO_PZ0", "", -1},
	// Older versions of L4T have a DT bug which instantiates a bogus device
	// which prevents this library from using this PWM channel.
	{168, "/sys/devices/6000d000.gpio", 32, 12, "GPIO07", "LCD_BL_PW", "/sys/devices/7000a000.pwm", 0},
	{38, "/sys/devices/6000d000.gpio", 33, 13, "GPIO13", "GPIO_PE6", "/sys/devices/7000a000.pwm", 2},
	{76, "/sys/devices/6000d000.gpio", 35, 19, "I2S0_FS", "DAP4_FS", "", -1},
	{51, "/sys/devices/6000d000.gpio", 36, 16, "UART1_CTS", "UART2_CTS", "", -1},
	{12, "/sys/devices/6000d000.gpio", 37, 26, "SPI1_MOSI", "SPI2_MOSI", "", -1},
	{77, "/sys/devices/6000d000.gpio", 38, 20, "I2S0_DIN", "DAP4_DIN", "", -1},
	{78, "/sys/devices/6000d000.gpio", 40, 21, "I2S0_DOUT", "DAP4_DOUT", "", -1},
}
var compatsNANO = []string{
	"nvidia,p3450-0000",
	"nvidia,p3450-0002",
	"nvidia,jetson-nano",
}

//Model is the model of the tegra chip
type Model struct {
	P1Revision   int
	RAM          string
	Revision     string
	Type         string
	Manufacturer string
	Processor    string
}

//GPIOData contains pin and chip data for jetson chip
type GPIOData struct {
	defs  []PinDef
	model Model
}

//GetPins returns the pindefs
func (g GPIOData) GetPins() []PinDef {
	return g.defs
}

//Model returns the model
func (g GPIOData) Model() Model {
	return g.model
}

//GetBoardPins gets the PinDefs associated with Board Pins passed
//If one or more is not supported function will return no defs and an error.
//
//Not all board pins are supported because they might be ground,power or some other reason.
func (g GPIOData) GetBoardPins(pins ...int) ([]PinDef, error) {
	pdefs := make([]PinDef, 0)
	for _, pin := range pins {
		var flag bool
		for _, pd := range g.defs {
			if pd.BoardPinNumber() == pin {
				pdefs = append(pdefs, pd)
				flag = true
				break
			}
		}
		if !flag {
			return nil, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)
		}

	}
	return pdefs, nil
}

//GetBoardPin gets a board pin
func (g GPIOData) GetBoardPin(pin int) (PinDef, error) {

	for _, pd := range g.defs {
		if pd.BoardPinNumber() == pin {
			return pd, nil

		}
	}

	return PinDef{}, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)

}

//GetLinuxPin gets the linux pin
func (g GPIOData) GetLinuxPin(pin int) (PinDef, error) {
	for _, pd := range g.defs {
		if pd.LinuxGPIOPinNum() == pin {
			return pd, nil

		}
	}

	return PinDef{}, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)
}

//GetLinuxPins returns an array of PinDefs that are associated with the pins passed.
//If one or more is not supported function will return no defs and an error.
func (g GPIOData) GetLinuxPins(pins ...int) ([]PinDef, error) {
	pdefs := make([]PinDef, 0)
	for _, pin := range pins {
		var flag bool
		for _, pd := range g.defs {
			if pd.LinuxGPIOPinNum() == pin {
				pdefs = append(pdefs, pd)
				flag = true
				break
			}
		}
		if !flag {
			return nil, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)
		}

	}
	return pdefs, nil
}

//GetBCMpin is like GetBoardPin except using BCM
func (g GPIOData) GetBCMpin(pin int) (PinDef, error) {
	for _, pd := range g.defs {
		if pd.BCMnum() == pin {
			return pd, nil

		}
	}

	return PinDef{}, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)
}

//GetBCMpins is like GetBoardPins except using BCM
func (g GPIOData) GetBCMpins(pins ...int) ([]PinDef, error) {
	pdefs := make([]PinDef, 0)
	for _, pin := range pins {
		var flag bool
		for _, pd := range g.defs {
			if pd.BCMnum() == pin {
				pdefs = append(pdefs, pd)
				flag = true
				break
			}
		}
		if !flag {
			return nil, fmt.Errorf("pin %d: NOT SUPPORTED ", pin)
		}

	}
	return pdefs, nil
}

//GetCVMpin is like GetBoardPin except using BCM
func (g GPIOData) GetCVMpin(pin string) (PinDef, error) {
	for _, pd := range g.defs {
		if pd.CVMname() == pin {
			return pd, nil

		}
	}

	return PinDef{}, fmt.Errorf("pin %s: NOT SUPPORTED ", pin)
}

//GetCVMpins is like GetBoardPins except using BCM
func (g GPIOData) GetCVMpins(pins ...string) ([]PinDef, error) {
	pdefs := make([]PinDef, 0)
	for _, pin := range pins {
		var flag bool
		for _, pd := range g.defs {
			if pd.CVMname() == pin {
				pdefs = append(pdefs, pd)
				flag = true
				break
			}
		}
		if !flag {
			return nil, fmt.Errorf("pin %s: NOT SUPPORTED ", pin)
		}

	}
	return pdefs, nil
}

//GetTegraNamepin is like GetBoardPin except using BCM
func (g GPIOData) GetTegraNamepin(pin string) (PinDef, error) {
	for _, pd := range g.defs {
		if pd.TegraPinName() == pin {
			return pd, nil

		}
	}

	return PinDef{}, fmt.Errorf("pin %s: NOT SUPPORTED ", pin)
}

//GetTegraNamepins is like GetBoardPins except using BCM
func (g GPIOData) GetTegraNamepins(pins ...string) ([]PinDef, error) {
	pdefs := make([]PinDef, 0)
	for _, pin := range pins {
		var flag bool
		for _, pd := range g.defs {
			if pd.TegraPinName() == pin {
				pdefs = append(pdefs, pd)
				flag = true
				break
			}
		}
		if !flag {
			return nil, fmt.Errorf("pin %s: NOT SUPPORTED ", pin)
		}

	}
	return pdefs, nil
}

//GetPWMPins returns the pindefs that can be PWM pins.  It returns nil if there isn't any.
//mulltiple PWM pins might have the same pwm pin number but with a different directory.
func (g GPIOData) GetPWMPins() []PinDef {
	pwm := make([]PinDef, 0)
	for _, pd := range g.defs {
		if pd.PWMid() != -1 {
			pdcopy := pd
			pwm = append(pwm, pdcopy)

		}
	}
	if len(pwm) == 0 {
		return nil
	}
	return pwm
}

var gpiodata = []GPIOData{

	GPIOData{
		defs: jetsonNXpindefs,
		model: Model{P1Revision: 1,
			RAM:          "16384M",
			Revision:     "Unknown",
			Type:         "Jetson NX",
			Manufacturer: "NVIDIA",
			Processor:    "ARM Carmel",
		},
	},
	GPIOData{
		defs: jetsonXAVIERpindefs,
		model: Model{P1Revision: 1,
			RAM:          "16384M",
			Revision:     "Unknown",
			Type:         "Jetson Xavier",
			Manufacturer: "NVIDIA",
			Processor:    "ARM Carmel",
		},
	},
	GPIOData{
		defs: jetsonTX2pindefs,
		model: Model{P1Revision: 1,
			RAM:          "8192M",
			Revision:     "Unknown",
			Type:         "Jetson TX2",
			Manufacturer: "NVIDIA",
			Processor:    "ARM A57 + Denver",
		},
	},
	GPIOData{
		defs: jetsonTX1pindefs,
		model: Model{P1Revision: 1,
			RAM:          "4096M",
			Revision:     "Unknown",
			Type:         "Jetson TX1",
			Manufacturer: "NVIDIA",
			Processor:    "ARM A57",
		},
	},
	GPIOData{
		defs: jetsonNANOpindefs,
		model: Model{P1Revision: 1,
			RAM:          "4096M",
			Revision:     "Unknown",
			Type:         "Jetson Nano",
			Manufacturer: "NVIDIA",
			Processor:    "ARM A57",
		},
	},
}
