package jetsongpio

type pindef struct {
	linuxpinnum        uint
	gpiosysfsdirectory string
	boardpinnum        uint
	bcmpinnum          uint
	cvmpinname         string
	tegrapinname       string
	pwmsysfsdirectory  string
	pwmidinpwmchip     uint
}

/*
func init() {
	jetsonnxpindefs
}
*/
var jetsonNXpindefs = []pindef{
	{148, "/sys/devices/2200000.gpio", 7, 4, "GPIO09", "AUD_MCLK", "", 0},
	{140, "/sys/devices/2200000.gpio", 11, 17, "UART1_RTS", "UART1_RTS", "", 0},
	{157, "/sys/devices/2200000.gpio", 12, 18, "I2S0_SCLK", "DAP5_SCLK", "", 0},
	{192, "/sys/devices/2200000.gpio", 13, 27, "SPI1_SCK", "SPI3_SCK", "", 0},
	{20, "/sys/devices/c2f0000.gpio", 15, 22, "GPIO12", "TOUCH_CLK", "", 0},
	{196, "/sys/devices/2200000.gpio", 16, 23, "SPI1_CS1", "SPI3_CS1_N", "", 0},
	{195, "/sys/devices/2200000.gpio", 18, 24, "SPI1_CS0", "SPI3_CS0_N", "", 0},
	{205, "/sys/devices/2200000.gpio", 19, 10, "SPI0_MOSI", "SPI1_MOSI", "", 0},
	{204, "/sys/devices/2200000.gpio", 21, 9, "SPI0_MISO", "SPI1_MISO", "", 0},
	{193, "/sys/devices/2200000.gpio", 22, 25, "SPI1_MISO", "SPI3_MISO", "", 0},
	{203, "/sys/devices/2200000.gpio", 23, 11, "SPI0_SCK", "SPI1_SCK", "", 0},
	{206, "/sys/devices/2200000.gpio", 24, 8, "SPI0_CS0", "SPI1_CS0_N", "", 0},
	{207, "/sys/devices/2200000.gpio", 26, 7, "SPI0_CS1", "SPI1_CS1_N", "", 0},
	{133, "/sys/devices/2200000.gpio", 29, 5, "GPIO01", "SOC_GPIO41", "", 0},
	{134, "/sys/devices/2200000.gpio", 31, 6, "GPIO11", "SOC_GPIO42", "", 0},
	{136, "/sys/devices/2200000.gpio", 32, 12, "GPIO07", "SOC_GPIO44", "/sys/devices/32f0000.pwm", 0},
	{105, "/sys/devices/2200000.gpio", 33, 13, "GPIO13", "SOC_GPIO54", "/sys/devices/3280000.pwm", 0},
	{160, "/sys/devices/2200000.gpio", 35, 19, "I2S0_FS", "DAP5_FS", "", 0},
	{141, "/sys/devices/2200000.gpio", 36, 16, "UART1_CTS", "UART1_CTS", "", 0},
	{194, "/sys/devices/2200000.gpio", 37, 26, "SPI1_MOSI", "SPI3_MOSI", "", 0},
	{159, "/sys/devices/2200000.gpio", 38, 20, "I2S0_DIN", "DAP5_DIN", "", 0},
	{158, "/sys/devices/2200000.gpio", 40, 21, "I2S0_DOUT", "DAP5_DOUT", "", 0},
}

var compatsNX = []string{
	"nvidia,p3509-0000+p3668-0000",
	"nvidia,p3509-0000+p3668-0001",
	"nvidia,p3449-0000+p3668-0000",
	"nvidia,p3449-0000+p3668-0001",
}

var jetsonXAVIERpindefs = []pindef{
	{134, "/sys/devices/2200000.gpio", 7, 4, "MCLK05", "SOC_GPIO42", "", 0},
	{140, "/sys/devices/2200000.gpio", 11, 17, "UART1_RTS", "UART1_RTS", "", 0},
	{63, "/sys/devices/2200000.gpio", 12, 18, "I2S2_CLK", "DAP2_SCLK", "", 0},
	{136, "/sys/devices/2200000.gpio", 13, 27, "PWM01", "SOC_GPIO44", "/sys/devices/32f0000.pwm", 0},
	// Older versions of L4T don"t enable this PWM controller in DT, so this PWM
	// channel may not be available.
	{105, "/sys/devices/2200000.gpio", 15, 22, "GPIO27", "SOC_GPIO54", "/sys/devices/3280000.pwm", 0},
	{8, "/sys/devices/c2f0000.gpio", 16, 23, "GPIO8", "CAN1_STB", "", 0},
	{56, "/sys/devices/2200000.gpio", 18, 24, "GPIO35", "SOC_GPIO12", "/sys/devices/32c0000.pwm", 0},
	{205, "/sys/devices/2200000.gpio", 19, 10, "SPI1_MOSI", "SPI1_MOSI", "", 0},
	{204, "/sys/devices/2200000.gpio", 21, 9, "SPI1_MISO", "SPI1_MISO", "", 0},
	{129, "/sys/devices/2200000.gpio", 22, 25, "GPIO17", "SOC_GPIO21", "", 0},
	{203, "/sys/devices/2200000.gpio", 23, 11, "SPI1_CLK", "SPI1_SCK", "", 0},
	{206, "/sys/devices/2200000.gpio", 24, 8, "SPI1_CS0_N", "SPI1_CS0_N", "", 0},
	{207, "/sys/devices/2200000.gpio", 26, 7, "SPI1_CS1_N", "SPI1_CS1_N", "", 0},
	{3, "/sys/devices/c2f0000.gpio", 29, 5, "CAN0_DIN", "CAN0_DIN", "", 0},
	{2, "/sys/devices/c2f0000.gpio", 31, 6, "CAN0_DOUT", "CAN0_DOUT", "", 0},
	{9, "/sys/devices/c2f0000.gpio", 32, 12, "GPIO9", "CAN1_EN", "", 0},
	{0, "/sys/devices/c2f0000.gpio", 33, 13, "CAN1_DOUT", "CAN1_DOUT", "", 0},
	{66, "/sys/devices/2200000.gpio", 35, 19, "I2S2_FS", "DAP2_FS", "", 0},
	//Input-only (due to base board)
	{141, "/sys/devices/2200000.gpio", 36, 16, "UART1_CTS", "UART1_CTS", "", 0},
	{1, "/sys/devices/c2f0000.gpio", 37, 26, "CAN1_DIN", "CAN1_DIN", "", 0},
	{65, "/sys/devices/2200000.gpio", 38, 20, "I2S2_DIN", "DAP2_DIN", "", 0},
	{64, "/sys/devices/2200000.gpio", 40, 21, "I2S2_DOUT", "DAP2_DOUT", "", 0},
}
var compatsXAVIER = []string{
	"nvidia,p2972-0000",
	"nvidia,p2972-0006",
	"nvidia,jetson-xavier",
}

var jetsonTX2pindefs = []pindef{
	{76, "/sys/devices/2200000.gpio", 7, 4, "AUDIO_MCLK", "AUD_MCLK", "", 0},
	// Output-only (due to base board)
	{146, "/sys/devices/2200000.gpio", 11, 17, "UART0_RTS", "UART1_RTS", "", 0},
	{72, "/sys/devices/2200000.gpio", 12, 18, "I2S0_CLK", "DAP1_SCLK", "", 0},
	{77, "/sys/devices/2200000.gpio", 13, 27, "GPIO20_AUD_INT", "GPIO_AUD0", "", 0},
	{15, "/sys/devices/3160000.i2c/i2c-0/0-0074", 15, 22, "GPIO_EXP_P17", "GPIO_EXP_P17", "", 0},
	// Input-only (due to module):
	{40, "/sys/devices/c2f0000.gpio", 16, 23, "AO_DMIC_IN_DAT", "CAN_GPIO0", "", 0},
	{161, "/sys/devices/2200000.gpio", 18, 24, "GPIO16_MDM_WAKE_AP", "GPIO_MDM2", "", 0},
	{109, "/sys/devices/2200000.gpio", 19, 10, "SPI1_MOSI", "GPIO_CAM6", "", 0},
	{108, "/sys/devices/2200000.gpio", 21, 9, "SPI1_MISO", "GPIO_CAM5", "", 0},
	{14, "/sys/devices/3160000.i2c/i2c-0/0-0074", 22, 25, "GPIO_EXP_P16", "GPIO_EXP_P16", "", 0},
	{107, "/sys/devices/2200000.gpio", 23, 11, "SPI1_CLK", "GPIO_CAM4", "", 0},
	{110, "/sys/devices/2200000.gpio", 24, 8, "SPI1_CS0", "GPIO_CAM7", "", 0},
	{0, "", 26, 7, "SPI1_CS1", "", "", 0},
	{78, "/sys/devices/2200000.gpio", 29, 5, "GPIO19_AUD_RST", "GPIO_AUD1", "", 0},
	{42, "/sys/devices/c2f0000.gpio", 31, 6, "GPIO9_MOTION_INT", "CAN_GPIO2", "", 0},
	// Output-only (due to module):
	{41, "/sys/devices/c2f0000.gpio", 32, 12, "AO_DMIC_IN_CLK", "CAN_GPIO1", "", 0},
	{69, "/sys/devices/2200000.gpio", 33, 13, "GPIO11_AP_WAKE_BT", "GPIO_PQ5", "", 0},
	{75, "/sys/devices/2200000.gpio", 35, 19, "I2S0_LRCLK", "DAP1_FS", "", 0},
	// Input-only (due to base board) IF NVIDIA debug card NOT plugged in
	// Output-only (due to base board) IF NVIDIA debug card plugged in
	{147, "/sys/devices/2200000.gpio", 36, 16, "UART0_CTS", "UART1_CTS", "", 0},
	{68, "/sys/devices/2200000.gpio", 37, 26, "GPIO8_ALS_PROX_INT", "GPIO_PQ4", "", 0},
	{74, "/sys/devices/2200000.gpio", 38, 20, "I2S0_SDIN", "DAP1_DIN", "", 0},
	{73, "/sys/devices/2200000.gpio", 40, 21, "I2S0_SDOUT", "DAP1_DOUT", "", 0},
}
var compatsTX2 = []string{
	"nvidia,p2771-0000",
	"nvidia,p2771-0888",
	"nvidia,p3489-0000",
	"nvidia,lightning",
	"nvidia,quill",
	"nvidia,storm",
}
var jetsonTX1pindefs = []pindef{
	{216, "/sys/devices/6000d000.gpio", 7, 4, "AUDIO_MCLK", "AUD_MCLK", "", 0},
	// Output-only (due to base board)
	{162, "/sys/devices/6000d000.gpio", 11, 17, "UART0_RTS", "UART1_RTS", "", 0},
	{11, "/sys/devices/6000d000.gpio", 12, 18, "I2S0_CLK", "DAP1_SCLK", "", 0},
	{38, "/sys/devices/6000d000.gpio", 13, 27, "GPIO20_AUD_INT", "GPIO_PE6", "", 0},
	{15, "/sys/devices/7000c400.i2c/i2c-1/1-0074", 15, 22, "GPIO_EXP_P17", "GPIO_EXP_P17", "", 0},
	{37, "/sys/devices/6000d000.gpio", 16, 23, "AO_DMIC_IN_DAT", "DMIC3_DAT", "", 0},
	{184, "/sys/devices/6000d000.gpio", 18, 24, "GPIO16_MDM_WAKE_AP", "MODEM_WAKE_AP", "", 0},
	{16, "/sys/devices/6000d000.gpio", 19, 10, "SPI1_MOSI", "SPI1_MOSI", "", 0},
	{17, "/sys/devices/6000d000.gpio", 21, 9, "SPI1_MISO", "SPI1_MISO", "", 0},
	{14, "/sys/devices/7000c400.i2c/i2c-1/1-0074", 22, 25, "GPIO_EXP_P16", "GPIO_EXP_P16", "", 0},
	{18, "/sys/devices/6000d000.gpio", 23, 11, "SPI1_CLK", "SPI1_SCK", "", 0},
	{19, "/sys/devices/6000d000.gpio", 24, 8, "SPI1_CS0", "SPI1_CS0", "", 0},
	{20, "/sys/devices/6000d000.gpio", 26, 7, "SPI1_CS1", "SPI1_CS1", "", 0},
	{219, "/sys/devices/6000d000.gpio", 29, 5, "GPIO19_AUD_RST", "GPIO_X1_AUD", "", 0},
	{186, "/sys/devices/6000d000.gpio", 31, 6, "GPIO9_MOTION_INT", "MOTION_INT", "", 0},
	{36, "/sys/devices/6000d000.gpio", 32, 12, "AO_DMIC_IN_CLK", "DMIC3_CLK", "", 0},
	{63, "/sys/devices/6000d000.gpio", 33, 13, "GPIO11_AP_WAKE_BT", "AP_WAKE_NFC", "", 0},
	{8, "/sys/devices/6000d000.gpio", 35, 19, "I2S0_LRCLK", "DAP1_FS", "", 0},
	// Input-only (due to base board) IF NVIDIA debug card NOT plugged in
	// Input-only (due to base board) (always reads fixed value) IF NVIDIA debug card plugged in
	{163, "/sys/devices/6000d000.gpio", 36, 16, "UART0_CTS", "UART1_CTS", "", 0},
	{187, "/sys/devices/6000d000.gpio", 37, 26, "GPIO8_ALS_PROX_INT", "ALS_PROX_INT", "", 0},
	{9, "/sys/devices/6000d000.gpio", 38, 20, "I2S0_SDIN", "DAP1_DIN", "", 0},
	{10, "/sys/devices/6000d000.gpio", 40, 21, "I2S0_SDOUT", "DAP1_DOUT", "", 0},
}
var compatsTX1 = []string{
	"nvidia,p2371-2180",
	"nvidia,jetson-cv",
}
var jetsonNANOpindefs = []pindef{
	{216, "/sys/devices/6000d000.gpio", 7, 4, "GPIO9", "AUD_MCLK", "", 0},
	{50, "/sys/devices/6000d000.gpio", 11, 17, "UART1_RTS", "UART2_RTS", "", 0},
	{79, "/sys/devices/6000d000.gpio", 12, 18, "I2S0_SCLK", "DAP4_SCLK", "", 0},
	{14, "/sys/devices/6000d000.gpio", 13, 27, "SPI1_SCK", "SPI2_SCK", "", 0},
	{194, "/sys/devices/6000d000.gpio", 15, 22, "GPIO12", "LCD_TE", "", 0},
	{232, "/sys/devices/6000d000.gpio", 16, 23, "SPI1_CS1", "SPI2_CS1", "", 0},
	{15, "/sys/devices/6000d000.gpio", 18, 24, "SPI1_CS0", "SPI2_CS0", "", 0},
	{16, "/sys/devices/6000d000.gpio", 19, 10, "SPI0_MOSI", "SPI1_MOSI", "", 0},
	{17, "/sys/devices/6000d000.gpio", 21, 9, "SPI0_MISO", "SPI1_MISO", "", 0},
	{13, "/sys/devices/6000d000.gpio", 22, 25, "SPI1_MISO", "SPI2_MISO", "", 0},
	{18, "/sys/devices/6000d000.gpio", 23, 11, "SPI0_SCK", "SPI1_SCK", "", 0},
	{19, "/sys/devices/6000d000.gpio", 24, 8, "SPI0_CS0", "SPI1_CS0", "", 0},
	{20, "/sys/devices/6000d000.gpio", 26, 7, "SPI0_CS1", "SPI1_CS1", "", 0},
	{149, "/sys/devices/6000d000.gpio", 29, 5, "GPIO01", "CAM_AF_EN", "", 0},
	{200, "/sys/devices/6000d000.gpio", 31, 6, "GPIO11", "GPIO_PZ0", "", 0},
	// Older versions of L4T have a DT bug which instantiates a bogus device
	// which prevents this library from using this PWM channel.
	{168, "/sys/devices/6000d000.gpio", 32, 12, "GPIO07", "LCD_BL_PW", "/sys/devices/7000a000.pwm", 0},
	{38, "/sys/devices/6000d000.gpio", 33, 13, "GPIO13", "GPIO_PE6", "/sys/devices/7000a000.pwm", 2},
	{76, "/sys/devices/6000d000.gpio", 35, 19, "I2S0_FS", "DAP4_FS", "", 0},
	{51, "/sys/devices/6000d000.gpio", 36, 16, "UART1_CTS", "UART2_CTS", "", 0},
	{12, "/sys/devices/6000d000.gpio", 37, 26, "SPI1_MOSI", "SPI2_MOSI", "", 0},
	{77, "/sys/devices/6000d000.gpio", 38, 20, "I2S0_DIN", "DAP4_DIN", "", 0},
	{78, "/sys/devices/6000d000.gpio", 40, 21, "I2S0_DOUT", "DAP4_DOUT", "", 0},
}
var compatsNANO = []string{
	"nvidia,p3450-0000",
	"nvidia,p3450-0002",
	"nvidia,jetson-nano",
}
