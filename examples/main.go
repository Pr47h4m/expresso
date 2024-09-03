package main

import (
	basicrouting "expresso_example/basic-routing"
	helloworld "expresso_example/hello-world"
	servestatic "expresso_example/serve-static"
	webservice "expresso_example/web-service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go helloworld.App().ListenAndServe(":81")
	// go helloworld.App().ListenAndServeTLS(":81", "server.crt", "server.key")
	go basicrouting.App().ListenAndServe(":82")
	// go basicrouting.App().ListenAndServeTLS(":82", "server.crt", "server.key")
	go servestatic.App().ListenAndServe(":83")
	// go servestatic.App().ListenAndServeTLS(":83", "server.crt", "server.key")
	go webservice.App().ListenAndServe(":84")
	// go webservice.App().ListenAndServeTLS(":84", "server.crt", "server.key")

	go func() {
		<-sig
		fmt.Println("\nCtrl+C pressed. Exiting...")
		done <- true
		os.Exit(0)
	}()

	<-done
}
