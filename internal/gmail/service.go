package gmail

type Service struct {
	config *GmailConfig
}

func NewService(config *GmailConfig) *Service {
	return &Service{config}
}
