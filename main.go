/*
* Project: IshalaChan
* Date: Friday, February 23, 2024
* File: main.go
* Author: Ieudru Volteal
* - - - - -
* Copyright © 2024 Ieudru Volteal
* All Rights Reserved
 */

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sess, prefix, err := initialize()
	if err != nil {
		log.Fatal(err)
	}

	handler(sess, prefix)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
