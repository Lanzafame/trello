// Copyright 2016 The trello client Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

t is a CLI that allows to add tasks into trello.

	usage: t [flags] title [description]
	  -c string
		config file (default: ~/.trc)
	  -debug
		debug mode

It also allows to specify, within the title, several "labels" (@label1
@label2), as well as one "board" (#board) and "list" (^List). E.g.:

	t "Add examples @dev @home ^Today #GTD" "Add examples in the documentation"

The configuration file has the following format:

	{
		"key": "KEY",
		"token": "TOKEN",
		"default_board": "BOARD_NAME",
		"default_list": "LIST_NAME"
	}

*/
package main

import (
	"flag"
	"log"
	"os/user"
	"path/filepath"
)

type config struct {
	Key          string `json:"key"`
	Token        string `json:"token"`
	DefaultBoard string `json:"default_board"`
	DefaultList  string `json:"default_list"`
}

var (
	cfgFile = flag.String("c", "", "config file (default: ~/.trc)")
	debug   = flag.Bool("debug", false, "debug mode")

	cfg config
)

func main() {
	//flag.Usage = usage
	flag.Parse()

	path := *cfgFile
	if *cfgFile == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatalln(err)
		}
		path = filepath.Join(usr.HomeDir, ".trc")
	}

	if cfg, err := parseConfig(path); err != nil {
		log.Fatalln(err)
	}

}
