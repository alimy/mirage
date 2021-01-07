// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by GNU General Public License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	. "github.com/alimy/mir/v2/core"
	. "github.com/alimy/mir/v2/engine"

	_ "github.com/alimy/mirage/mirc/routes"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialDebugMode),
		GeneratorName(GeneratorGin),
		SinkPath("auto"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
