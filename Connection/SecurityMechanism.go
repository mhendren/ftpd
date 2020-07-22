package Connection

type SecurityMechanism interface {
	SupportedProtections() []DataChannelProtectionLevel
	Name() string
}
