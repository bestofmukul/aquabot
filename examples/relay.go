package main

import (
	"flag"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// Toggles relay state(on/off).
// Note: Electro-mechanical relays are not meant for high speed switching, use long time.
func main() {
	toggleTimer := flag.Duration("timer", 15*time.Second, "Toggle time in seconds")
	flag.Parse()

	pi := raspi.NewAdaptor()
	relay := gpio.NewRelayDriver(pi, "11")

	task := func() {
		gobot.Every(*toggleTimer, func() {
			log.Println("Toggling relay")
			relay.Toggle()
		})

	}

	robot := gobot.NewRobot("AquaBot",
		[]gobot.Connection{pi},
		[]gobot.Device{relay},
		task,
	)

	robot.Start()
}
