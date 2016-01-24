package main

import (
	"github.com/gophergala2016/sisho"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var (
	create    = kingpin.Command("create", "Create epub file")
	repositry = create.Arg("pathToRepository", "").Required().String()
)

func main() {
	var err error
	kingpin.Version("0.1.0")

	switch kingpin.Parse() {
	case "create":
		err = sisho.Run(*repositry)
	}

	if err != nil {
		log.Println(err)
	}
}
