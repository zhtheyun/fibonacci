//package main is the entry point for fibonacci webservice system.
//Fibonacci webservice system allow user generate fibonacci number through restful API.
package main

import "github.com/zhtheyun/fibonacci/cmd"

//BuildDate records the date for the binaries
var BuildDate string

//Version records the binaries version
var Version string

func main() {
	cmd.Execute(Version, BuildDate)
}
