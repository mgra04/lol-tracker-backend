package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
    Client *mongo.Client
    DB     *mongo.Database
}

func NewMongoClient(uri string) (*Repository, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOpts := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        return nil, err
    }

    // Check the connection
    if err := client.Ping(ctx, nil); err != nil {
        return nil, err
    }

    db := client.Database("lol_tracker") // nazwę DB możesz dać w config

    return &Repository{
        Client: client,
        DB:     db,
    }, nil
}

func (r *Repository) Disconnect(ctx context.Context) error {
    return r.Client.Disconnect(ctx)
}
