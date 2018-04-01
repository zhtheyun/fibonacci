package main

import "github.com/zhtheyun/fibonacci/cmd"

//BuildDate records the date for the binaries
var BuildDate string

//Version records the binaries version
var Version string

func main() {
	cmd.Execute(Version, BuildDate)
}
