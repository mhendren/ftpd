package Connection

type Status struct {
	Connected          bool
	Remote             string
	Authenticated      bool
	Anonymous          bool
	User               string
	CurrentPath        string
	AcceptableCommands map[string]bool
}

func (cs *Status) Connect() {
	cs.Connected = true
}

func (cs *Status) Disconnect() {
	cs.Connected = false
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
