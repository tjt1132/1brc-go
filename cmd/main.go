package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"

	"github.com/tjt1132/1brc-go/aggregator"
)

var cpuProfile = flag.String("cp", "", "write cpu profile to `file`")
var memProfile = flag.String("mp", "", "write memory profile to `file`")
var exeProfile = flag.String("ep", "", "write trace execution to `file`")
var filepath = flag.String("f", "", "path to measurements `file`")

func main() {
	flag.Parse()

	if *exeProfile != "" {
		f, err := os.Create(*exeProfile)
		if err != nil {
			log.Fatal("could not create trace execution profile: ", err)
		}
		defer f.Close()
		trace.Start(f)
		defer trace.Stop()
	}

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	aggregator.Do(*filepath)

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
