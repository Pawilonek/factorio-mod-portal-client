package http

import "context"

type Client interface {
	Get(context.Context, string) (string, error)
}
