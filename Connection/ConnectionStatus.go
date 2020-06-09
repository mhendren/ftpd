package Connection

import (
	"FTPserver/Replies"
	"fmt"
	"net"
	"os"
)

type TransferMode int

const (
	Active int = iota
	Passive
)

func (tm TransferMode) String() string {
	return [...]string{"Active", "Passive"}[tm]
}

type Status struct {
	Connected          bool
	CommandConnection  net.Conn
	DataConnected      bool
	DataConnection     net.Conn
	Remote             string
	Authenticated      bool
	Anonymous          bool
	User               string
	CurrentPath        string
	AcceptableCommands map[string]bool
	TypeCode           string
	FormCode           string
	LocalByteSize      uint8
	Mode               TransferMode
	RenameFrom         string
}

func (cs *Status) Connect() {
	cs.Connected = true
}

func (cs *Status) Disconnect() {
	cs.Connected = false
}

func (cs *Status) DataConnect(conn net.Conn) {
	cs.DataConnected = true
	cs.DataConnection = conn
}

func (cs *Status) DataDisconnect() {
	if cs.DataConnection != nil {
		cs.DataConnection.Close()
	}
	cs.DataConnected = false
	cs.DataConnection = nil
}

func (cs *Status) IsConnected() bool {
	return cs.Connected
}

func (cs *Status) SetRemote(addr string) {
	cs.Remote = addr
}

func (cs *Status) SetAuthenticated(isAuth bool) {
	cs.Authenticated = isAuth
}

func (cs *Status) SetAnonymous(isAnonymous bool) {
	cs.Anonymous = isAnonymous
}

func (cs *Status) SetUser(username string) {
	cs.User = username
}

func (cs *Status) SetPath(path string) {
	cs.CurrentPath = path
}

func (cs *Status) SetCommands(commands map[string]bool) {
	cs.AcceptableCommands = commands
}

func (cs *Status) CanUseCommand(command string) bool {
	if cs.AcceptableCommands == nil {
		return true
	}
	_, ok := cs.AcceptableCommands[command]
	return ok
}

func (cs Status) SendFTPReply(reply Replies.FTPReply) {
	if reply.Code < 0 {
		_, _ = fmt.Fprintf(os.Stderr, "SendFTPReply: No action requested (communication handled by Command)")
		return
	}
	_, err := fmt.Fprint(cs.CommandConnection, reply)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "SendFTPReply: %v\n", err)
	}
}
