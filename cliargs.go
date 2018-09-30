package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Hash    string `short:"H" long:"hash" description:"A file hash" required:"false"`
	File    string `short:"f" long:"file" description:"A file" value-name:"FILE"`
	Config  string `short:"c" long:"config" description:"A config file that contains the API key and sandbox"`
}

func ParseArgs() (*Options, error) {
	opts := &Options{}
	_, err := flags.ParseArgs(opts, os.Args)
	if err != nil {
		return nil, err
	}
	return opts, nil
}
