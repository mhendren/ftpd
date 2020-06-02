package Configuration

type FTPConfig struct {
	AllowAnonymous bool
	AllowNoLogin   bool
	RootPath       string
	BasePort       uint16
}
