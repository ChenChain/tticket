package util

import "context"

var DEFAULT_RETRY_TIMES = 3

func Retry(ctx context.Context, f func(ctx context.Context) error) error {
	var err error
	for i := 0; i < DEFAULT_RETRY_TIMES; i++ {
		err = f(ctx)
		if err == nil {
			return nil
		}
	}
	return err
}
