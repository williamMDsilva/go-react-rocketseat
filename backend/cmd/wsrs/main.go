package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/williamMDsilva/str-go-back-end/internal/api"
	"github.com/williamMDsilva/str-go-back-end/internal/store/pgstore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ctx := context.Background()

	pool, err := pgxpool.New(ctx,
		fmt.Sprintf(
			"user=%s password=%s host=%s dbname=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("POSTGRES_DB"),
		),
	)
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		if err := http.ListenAndServe(port, handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	fmt.Print(fmt.Sprintf("Server is runnig on %s ... \nCTRL + C to quit\n", port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
