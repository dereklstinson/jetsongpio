# JetsonGPIO
Jetson GPIO for Go

This is not working yet.

## Prep Steps

Must Set User Permissions

```
sudo groupadd -f -r gpio
sudo usermod -a -G gpio your_user_name
```

Copy 99-gpio.rules into rules.d directory.  Need to be in directory package was installed

``` 
sudo cp 99-gpio.rules /etc/udev/rules.d/
```

Reboot or reload udev rules

```
sudo udevadm control --reload-rules && sudo udevadm trigger
```