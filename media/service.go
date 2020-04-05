package media

type Service interface {
	FindByTicketId(string) []string
}

type DefaultService struct{}

func (m DefaultService) FindByTicketId(tid string) []string {
	return []string{"Media 1", "Media 2", "Media 3"}
}

func NewMediaService() DefaultService {
	return DefaultService{}
}
