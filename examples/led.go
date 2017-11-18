package main

import (
	"flag"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// Toggles LED state(on/off) at duration provided as cmdline arg.
func main() {
	toggleTimer := flag.Duration("timer", 5*time.Second, "Toggle time in seconds")
	flag.Parse()

	pi := raspi.NewAdaptor()
	led := gpio.NewLedDriver(pi, "11")

	work := func() {
		gobot.Every(*toggleTimer, func() {
			log.Println("Toggling LED")
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{pi},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
