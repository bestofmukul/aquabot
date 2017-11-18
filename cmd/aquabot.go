package main

import (
	"flag"
	"log"
	"regexp"
	"time"

	"github.com/jasonlvhit/gocron"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var (
	timeFormat = regexp.MustCompile(`([01][0-9]|2[0-3]):[0-5][0-9]`)
)

func main() {
	days := flag.Uint64("days", 1, "Frequency at which sprinklers will start")
	sprinklerTimer := flag.Duration("sprinkle", 15*time.Second, "Run sprinklers for this much time in seconds")
	daytime := flag.String("time", "10:00", "Time(format hh:mm) at which sprinklers will run")
	flag.Parse()

	if !timeFormat.MatchString(*daytime) {
		log.Fatalf("Invalid time provided. Must be in hh:mm format. Found: %s", *daytime)
	}

	pi := raspi.NewAdaptor()
	relay := gpio.NewRelayDriver(pi, "11")

	task := func() {
		log.Println("Starting sprinklers")
		relay.On()

		// turning relay off will stop current flow and motor will stop
		stopTimer := time.NewTimer(*sprinklerTimer)
		for {
			select {
			case <-stopTimer.C:
				log.Println("Stopping sprinklers")
				relay.Off()
				return
			}
		}

	}

	// schedule sprinklers to run at pre-defined frequency
	work := func() {
		relay.Off()
		sched := gocron.NewScheduler()
		sched.Every(*days).Day().At(*daytime).Do(task)
		<-sched.Start()
	}

	robot := gobot.NewRobot("AquaBot",
		[]gobot.Connection{pi},
		[]gobot.Device{relay},
		work,
	)

	robot.Start()
}
