package http

import "context"

type Alerter interface {
	Alert(ctx context.Context, s string) error
}
