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

	go helloworld.App().ListenAndServe(":81", nil)
	// go helloworld.App().ListenAndServeTLS(":81", "server.crt", "server.key", func(err error) {
	// 	if err != nil {
	// 		fmt.Println("Unable to start helloworld server on port 81", err)
	// 	} else {
	// 		fmt.Println("helloworld server is running on port 81")
	// 	}
	// })
	go basicrouting.App().ListenAndServe(":82", nil)
	// go basicrouting.App().ListenAndServeTLS(":82", "server.crt", "server.key", func(err error) {
	// 	if err != nil {
	// 		fmt.Println("Unable to start basicrouting server on port 81", err)
	// 	} else {
	// 		fmt.Println("basicrouting server is running on port 81")
	// 	}
	// })
	go servestatic.App().ListenAndServe(":83", func(err error) {
		if err != nil {
			fmt.Println("Unable to start servestatic server on port 81", err)
		} else {
			fmt.Println("servestatic server is running on port 81")
		}
	})
	// go servestatic.App().ListenAndServeTLS(":83", "server.crt", "server.key", nil)
	go webservice.App().ListenAndServe(":84", func(err error) {
		if err != nil {
			fmt.Println("Unable to start webservice server on port 81", err)
		} else {
			fmt.Println("webservice server is running on port 81")
		}
	})
	// go webservice.App().ListenAndServeTLS(":84", "server.crt", "server.key", nil)

	go func() {
		fmt.Println("press Ctrl+C to exit.")
		<-sig
		fmt.Println("\nCtrl+C pressed. Exiting...")
		done <- true
		os.Exit(0)
	}()

	<-done
}
