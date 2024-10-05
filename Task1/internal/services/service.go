package services

type Service struct {
	client Client
}

func New(client Client) *Service {
	return &Service{
		client: client,
	}
}
