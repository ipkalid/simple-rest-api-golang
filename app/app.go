package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	client *redis.Client
}

func NewApp() *App {
	app := &App{
		client: loadClient(),
	}
	app.loadRouter()
	return app
}

func loadClient() *redis.Client {

	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	return client
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{Addr: ":3000", Handler: a.router}
	err := a.client.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis %w", err)
	}
	defer func() {
		if err := a.client.Close(); err != nil {
			fmt.Println("failed to close Redis %w", err)
		}
	}()

	ch := make(chan error, 1)
	fmt.Println("app start at http://localhost:3000")
	go func() {

		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %w", err)
		}
		close(ch)
	}()
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
