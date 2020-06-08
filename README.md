# JetsonGPIO
Jetson GPIO for Go

Use this at your own risk.  If some of the pins are mixed up in the software then bad things could happen.  Not saying that I would do it on purpose, but there is always the chance.  The defs for the pins were pulled from NVIDIA/Jetson-gpio so I think it should be ok for the development boards.  That being said not much is going to be tested, because I am not paid for this.  

Pull requests are welcomed.  So if a bug is found please submit a pull request.  Events are not done yet. I am going to use go channels for that.  I will have examples later.

## Prep Steps

Must Set User Permissions

```
sudo groupadd -f -r gpio
sudo usermod -a -G gpio your_user_name
```

go to setup folder in jetsongpio 

run setup_gpio.sh

```
./setup_gpio.sh
```

Reboot or reload udev rules

```
sudo udevadm control --reload-rules && sudo udevadm trigger
```