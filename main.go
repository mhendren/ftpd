package main

import "FTPserver/Server"

func main() {
	server := Server.SetupConnection()
	Server.Loop(server)
}
