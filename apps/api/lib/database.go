package lib

import (
	"context"
	"fmt"

	"github.com/kkrishguptaa/whitestone-capital/apps/api/ent"

	_ "github.com/lib/pq"
)

var Client *ent.Client

func InitDatabase() {
	var err error

	Client, err = ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Env.DATABASE_HOST,
		Env.DATABASE_PORT,
		Env.DATABASE_USER,
		Env.DATABASE_PASSWORD,
		Env.DATABASE_NAME,
	))

	if err != nil {
		panic(err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}
