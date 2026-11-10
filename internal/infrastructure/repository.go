package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/maadco/go-crud-test/internal/domain"
)

type InMemoryCustomerRepository struct {
	mu        sync.RWMutex
	customers map[string]*domain.Customer
}

func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[string]*domain.Customer),
	}
}

func (r *InMemoryCustomerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.customers[string(customer.ID)] = customer
	return nil
}

func (r *InMemoryCustomerRepository) FindByID(ctx context.Context, id string) (*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	c, ok := r.customers[id]
	if !ok {
		return nil, errors.New("customer not found")
	}
	return c, nil
}

func (r *InMemoryCustomerRepository) FindByUniqueKey(ctx context.Context, firstName, lastName string, dob string) (*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}
	for _, c := range r.customers {
		if c.FirstName == firstName && c.LastName == lastName && c.DateOfBirth.Format("2006-01-02") == parsedDob.Format("2006-01-02") {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *InMemoryCustomerRepository) FindByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, c := range r.customers {
		if string(c.Email) == email {
			return c, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *InMemoryCustomerRepository) FindAll(ctx context.Context) ([]*domain.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*domain.Customer, 0, len(r.customers))
	for _, c := range r.customers {
		list = append(list, c)
	}
	return list, nil
}

func (r *InMemoryCustomerRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.customers[id]; !ok {
		return errors.New("customer not found")
	}
	delete(r.customers, id)
	return nil
}

type EventRecord struct {
	ID          uint `gorm:"primaryKey"`
	EventID     string
	AggregateID string
	EventType   string
	Payload     string
	OccurredAt  time.Time
}

type InMemoryEventStore struct {
	mu     sync.Mutex
	events []domain.DomainEvent
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		events: make([]domain.DomainEvent, 0),
	}
}

func (s *InMemoryEventStore) SaveEvent(ctx context.Context, event domain.DomainEvent) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	_, _ = json.Marshal(event)
	return nil
}

func (s *InMemoryEventStore) GetEvents(ctx context.Context, aggregateID string) ([]domain.DomainEvent, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []domain.DomainEvent
	for _, e := range s.events {
		if e.AggregateID() == aggregateID {
			result = append(result, e)
		}
	}
	return result, nil
}
