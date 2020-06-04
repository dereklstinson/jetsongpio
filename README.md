# JetsonGPIO
Jetson GPIO for Go

Well, right now all I have is the intern subpackage made, and it finally works without super user permissions.  YAY!  I would like to eventually build this
around I think it is called /dev/mem.   It will make accessing the pins super fast.  

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