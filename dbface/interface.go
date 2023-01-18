package dbface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

type (
	Collection interface {
		InsertOne(ctx context.Context, document interface{}, contextopts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	}
)