package payment

import (
	"errors"
	"fmt"
	"time"
)

// Middleware decorates a service.
type Middleware func(Service) Service

type Service interface {
	Authorise(total float32) (Authorisation, error) // GET /paymentAuth
	Health() []Health                               // GET /health
}

type Authorisation struct {
	Authorised bool   `json:"authorised"`
	Message    string `json:"message"`
}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// NewFixedService returns a simple implementation of the Service interface,
// fixed over a predefined set of socks and tags. In a real service you'd
// probably construct this with a database handle to your socks DB, etc.
func NewAuthorisationService(declineOverAmount float32, authorizer string) Service {
	return &service{
		declineOverAmount: declineOverAmount,
		authorizer: authorizer,
	}
}

type service struct {
	declineOverAmount float32
	authorizer string
}

func (s *service) Authorise(amount float32) (Authorisation, error) {
	if amount == 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if amount < 0 {
		return Authorisation{}, ErrInvalidPaymentAmount
	}
	if s.authorizer == "" {
		return Authorisation{}, ErrNoAuthorizer
	}
	authorised := false
	message := "Payment declined"
	if amount <= s.declineOverAmount {
		authorised = true
		message = fmt.Sprintf("Payment authorised by %s", s.authorizer)
	} else {
		message = fmt.Sprintf("Payment declined: amount exceeds %.2f", s.declineOverAmount)
	}
	return Authorisation{
		Authorised: authorised,
		Message:    message,
	}, nil
}

func (s *service) Health() []Health {
	var health []Health
	app := Health{"payment", "OK", time.Now().String()}
	health = append(health, app)
	return health
}

var ErrInvalidPaymentAmount = errors.New("Invalid payment amount")
var ErrNoAuthorizer = errors.New("No authorizer supplied for payment. Did you set the AUTH_BY environment variable?")
