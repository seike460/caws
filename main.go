package main

import (
	"log"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	caws := newCaws()

	// Set Service List
	caws.LList.AddItem(" S3 ", " Objects Storage ", 0, func() {
		setS3Service(caws)
	})
	caws.LList.AddItem(" Quit ", " Press to exit ", 0, func() {
		caws.App.Stop()
	})

	if err := caws.App.SetRoot(caws.Flex, true).Run(); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
