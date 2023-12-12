package mongodb

import (
	"context"
	"os"

	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	logger.Info("Init Database Connection", zap.String("journey", "InitConnection"))

	dburl := os.Getenv("DB_URL")
	dbname := os.Getenv("DB_NAME")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))

	if err != nil {
		logger.Error("Database Connection failed", err, zap.String("journey", "InitConnection"))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Database ping failed", err, zap.String("journey", "InitConnection"))
		return nil, err
	}

	logger.Info("Database Connection successfully", zap.String("journey", "InitConnection"))
	return client.Database(dbname), nil
}
