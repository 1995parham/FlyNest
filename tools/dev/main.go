package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"./test"

	"github.com/elahejalalpour/beehive-netctrl/controller"

	"../../discovery"
	"../../openflow"

	bh "github.com/kandoo/beehive"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	h := bh.NewHive()
	openflow.StartOpenFlow(h)
	controller.RegisterNOMController(h)
	discovery.RegisterDiscovery(h)

	// Register a switch:
	// switching.RegisterSwitch(h, bh.Persistent(1))
	// or a hub:
	// switching.RegisterHub(h, bh.NonTransactional())

	test.StartTest(h)

	h.Start()
}
