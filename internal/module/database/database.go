package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Mongo struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  context.CancelFunc
}

func (m *Mongo) Connect(uri string) error {
	var err error
	m.Context, m.Cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer m.Cancel()
	if m.Client, err = mongo.Connect(options.Client().ApplyURI(uri)); err != nil {
		return fmt.Errorf("failed to connect: %s", err.Error())
	}

	if err = m.Client.Ping(m.Context, readpref.Primary()); err != nil {
		return fmt.Errorf("failed to connect: %s", err.Error())
	}
	return nil
}
