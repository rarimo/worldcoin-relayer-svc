package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/rarify-protocol/relayer-svc/internal/config"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	configCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

// CtxConfig adds config provider instance to ctx.
func CtxConfig(cfg config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, cfg)
	}
}

// Config returns the config provider instance stored in ctx.
func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}