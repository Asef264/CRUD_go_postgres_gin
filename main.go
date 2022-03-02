package main

import (
	"sample/handller"
	"sample/repository"
)

func main() {
	repository.Connect()
	handller.Run()
}
