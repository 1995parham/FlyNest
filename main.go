package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"

	"./test"

	"github.com/1995parham/flynest/controller"
	"github.com/1995parham/flynest/discovery"
	"github.com/1995parham/flynest/intent"
	"github.com/1995parham/flynest/openflow"

	bh "github.com/kandoo/beehive"
)

var cpu_profile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpu_profile != "" {
		f, err := os.Create(*cpu_profile)
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
	intent.RegisterIntent(h)

	// Register a switch:
	// switching.RegisterSwitch(h, bh.Persistent(1))
	// or a hub:
	// switching.RegisterHub(h, bh.NonTransactional())

	test.StartTest(h)

	h.Start()
}
