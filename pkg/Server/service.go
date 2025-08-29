package Server

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	Create(req *SubscriptionRequest) (Subscription, error)
	GetAll() ([]Subscription, error)
	Update(id string, req *SubscriptionRequest) (Subscription, error)
	Delete(id string) error
	GetByFilter(req *ToFilter) (string, error)
	// GetByID(id string) (bool, error)
	// EndDate(date string) (string, error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(r RepositoryInterface) ServiceInterface {
	return &Service{repo: r}
}

func (s *Service) Create(req *SubscriptionRequest) (Subscription, error) {
	endDate, err := s.EndDate(req.StartDate)
	if err != nil {
		return Subscription{}, err
	}
	subsc := Subscription{
		ID:          uuid.New().String(),
		Price:       req.Price,
		ServiceName: req.ServiceName,
		StartDate:   req.StartDate,
		UserID:      req.UserID,
		EndDate:     endDate,
	}

	if err := s.repo.Create(&subsc); err != nil {
		return Subscription{}, err
	}
	return subsc, nil

}

func (s *Service) GetAll() ([]Subscription, error) {
	return s.repo.GetAll()
}

func (s *Service) Update(id string, req *SubscriptionRequest) (Subscription, error) {
	exist, err := s.GetByID(id)
	if !exist || err != nil {
		return Subscription{}, err
	}
	endDate, err := s.EndDate(req.StartDate)
	if err != nil {
		return Subscription{}, err
	}
	newSubsc := Subscription{
		ID:          id,
		Price:       req.Price,
		ServiceName: req.ServiceName,
		StartDate:   req.StartDate,
		UserID:      req.UserID,
		EndDate:     endDate,
	}
	if err := s.repo.Update(newSubsc); err != nil {
		return Subscription{}, err
	}
	return newSubsc, nil
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *Service) GetByFilter(req *ToFilter) (string, error) {
	prices, err := s.repo.GetByFilter(req)
	if err != nil {
		return "", err
	}
	var price int
	for _, p := range prices {
		price += p
	}
	result := strconv.Itoa(price)
	return result, nil
}

func (s *Service) EndDate(date string) (string, error) {
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	endDate := startDate.AddDate(0, 3, 0)
	result := endDate.Format("2006-01-02")
	return result, nil
}
func (s *Service) GetByID(id string) (bool, error) {
	return s.repo.GetByID(id)
}
