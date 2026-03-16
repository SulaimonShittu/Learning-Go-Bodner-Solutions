package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

func main() {

}

func middleware(runTime time.Duration) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(context.Background(), runTime)
			defer cancel()
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
	}
}

func random() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx, cancelCase := context.WithCancelCause(ctx)
	defer cancelCase(nil)

	count := 0
	sum := 0
	go func() {
		for {
			x := rand.IntN(100_000_000)
			y := rand.IntN(100_000_000)
			sum = x + y
			count++
		}
	}()

}
