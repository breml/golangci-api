package paymentprovider

import "context"

type Provider interface {
	Name() string

	SetBaseURL(url string) error

	CreateCustomer(ctx context.Context, email string, token string) (*Customer, error)

	GetSubscription(ctx context.Context, cust string, sub string) (*Subscription, error)
	GetSubscriptions(ctx context.Context, cust string) ([]Subscription, error)
	CreateSubscription(ctx context.Context, cust string, seats int) (*Subscription, error)
	UpdateSubscription(ctx context.Context, cust string, sub string, payload SubscriptionUpdatePayload) (*Subscription, error)
	DeleteSubscription(ctx context.Context, cust string, sub string) error

	GetEvent(ctx context.Context, eventID string) (*Event, error)
}
