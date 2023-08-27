package increasing

type UserCounterIncreaserService struct{}

func NewUserCounterIncreaserService() UserCounterIncreaserService {
	return UserCounterIncreaserService{}
}

func (s UserCounterIncreaserService) Increase(id string) error {
	return nil
}
