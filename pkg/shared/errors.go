package shared

import "context"

type contextCloser interface {
	Close(ctx context.Context) error
}

func PanicOnClosureError(ctx context.Context, closer contextCloser) {
	PanicOnErr(closer.Close(ctx))
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
