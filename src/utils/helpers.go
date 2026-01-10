package paymentprocessor

import (
	"fmt"
	"strings"
)

type Status string

const (
	NEW      Status = "new"
	PENDING  Status = "pending"
	CANCELED Status = "canceled"
	FAILED   Status = "failed"
	PAID     Status = "paid"
)

func NewStatus(status string) (Status, error) {
	switch strings.ToLower(status) {
	case "new":
		return NEW, nil
	case "pending":
		return PENDING, nil
	case "canceled":
		return CANCELED, nil
	case "failed":
		return FAILED, nil
	case "paid":
		return PAID, nil
	default:
		return "", fmt.Errorf("invalid status: %s", status)
	}
}