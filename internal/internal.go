package internal

import (
	"context"
	"log"
	"math/rand"

	"yola/internal/entdata"
	"yola/internal/entdata/migrate"

	"entgo.io/ent/dialect"
)

func NewEntClient() *entdata.Client {
	client, err := entdata.Open(dialect.SQLite, "assets/files/yola.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed connecting to sqlite: %v", err)
	}
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func Shuffle[T any](data []T) []T {
	length := len(data)
	result := make([]T, length)
	perm := rand.Perm(length)
	for i, v := range perm {
		result[v] = data[i]
	}
	return data
}
