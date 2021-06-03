// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	govtm "github.com/ysicing/go-vtm"
	"log"
)

func main()  {
	vtm, contactable, err := govtm.NewVirtualTrafficManager(
		"https://172.16.74.115:9070/api",
		"admin",
		"12345678",
		false,
		false,
	)
	if !contactable {
		panic(err)
	}
	status, err  := vtm.GetSystemState()
	if err != nil {
		panic(err.ErrorText)
	}
	log.Println("License: ", *status.State.License)
}