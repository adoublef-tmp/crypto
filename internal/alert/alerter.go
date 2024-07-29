package alert

import (
	"context"
	"errors"
	"log"
)

type Alerter struct{}

func (a Alerter) Alert(ctx context.Context, s string) error {
	if s == "" {
		return errors.New("empty input")
	}
	log.Printf("alerting %s", s)
	return nil
}
