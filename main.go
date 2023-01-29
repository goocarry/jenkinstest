package main

import "github.com/goocarry/jenkinslocal/pkg/logger"

func main() {
	logger := &logger.Logger{}
	logger.Log("hello")
}
