package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {

	//flag.IntVar(&port, "port", 8081, "websocket port, provides telemetry data to OpenMCT")
	//flag.IntVar(&timerInterval, "timer", 1000, "main timer interval")
	//flag.Usage = func() {
	//	flag.PrintDefaults()
	//}
	//flag.Parse()

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
		<-sc

		//TODO free resources ...

		// close all connections
		CloseAll()

		fmt.Print("\n")
		log.Println("Have a nice day!")
		os.Exit(0)
	}()

	ListenAndServe(8081, 1000)
}
