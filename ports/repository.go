package ports

import (
	"context"
	"github.com/maadco/go-crud-test/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *domain.Customer) error
	FindByID(ctx context.Context, id string) (*domain.Customer, error)
	FindByUniqueKey(ctx context.Context, firstName, lastName string, dob string) (*domain.Customer, error)
	FindByEmail(ctx context.Context, email string) (*domain.Customer, error)
	FindAll(ctx context.Context) ([]*domain.Customer, error)
	Delete(ctx context.Context, id string) error
}

type EventStore interface {
	SaveEvent(ctx context.Context, event domain.DomainEvent) error
	GetEvents(ctx context.Context, aggregateID string) ([]domain.DomainEvent, error)
}
