package Connection

type SecurityTLS struct {
}

func (sm SecurityTLS) SupportedProtections() []DataChannelProtectionLevel {
	return []DataChannelProtectionLevel{
		DataChannelProtectionLevel(Clear),
		DataChannelProtectionLevel(Private),
	}
}

func (sm SecurityTLS) Name() string {
	return "TLS"
}
