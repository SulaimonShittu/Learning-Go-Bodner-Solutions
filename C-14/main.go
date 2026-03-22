package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

func main() {

}

func middlewareGenerator(runTime time.Duration) func(http.Handler) http.Handler {
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
	var x, y int
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				count++
				x = rand.IntN(100_000_000)
				y = rand.IntN(100_000_000)
				if x+y == 1234 {
					cancelCase(errors.New("number reached"))
					return
				}
			}
		}
	}()
	<-ctx.Done()
	fmt.Printf("The sum is: %v, the number of iterations: %v, reason for end: %v", x+y, count, ctx.Err())
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel = LevelFromContext(ctx)
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, "level", level)
}

func LevelFromContext(ctx context.Context) Level {
	return ctx.Value("level").(Level)
}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var level Level = Level(r.URL.Query().Get("log_level"))
		r = r.WithContext(context.WithValue(context.Background(), "level", level))
		h.ServeHTTP(w, r)
	})
}
