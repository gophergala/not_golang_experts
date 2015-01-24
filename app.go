package main

import (
	"github.com/gophergala/not_golang_experts/worker"
)

func main() {
	stopped := make(chan bool, 1)
	worker.StartObserving(stopped)

	<-stopped
}
