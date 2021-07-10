package domain

import "time"
import uuid "github.com/satori/go.uuid"

type CreditCard struct {
	ID string
	Name string
	Number string
	ExpirationMonth int32
	ExpirationYear int32
	CVV int32
	Balance float64
	Limit float64
	CreatedAt time.Time
}	

func NewCreditCard() *CreditCard{
	c := &CreditCard{}
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.Now()

	return c
}

func (t *Transaction) ProcessAndValitade(creditCard *CreditCard ){
	if t.Amount + creditCard.Balance > creditCard.Limit {
		t.Status = "Rejected"
	} else {
		t.Status = "Approved"
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}