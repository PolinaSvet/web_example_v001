package main

import cmd "tracker/cmd/web"

func init() {
	cmd.InitVariable()
}

func main() {
	cmd.Handler()
}
