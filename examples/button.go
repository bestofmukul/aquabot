package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	pi := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(pi, "3")
	led := gpio.NewLedDriver(pi, "11")

	work := func() {

		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
			led.Off()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
			led.On()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{pi},
		[]gobot.Device{button, led},
		work,
	)

	robot.Start()
}
