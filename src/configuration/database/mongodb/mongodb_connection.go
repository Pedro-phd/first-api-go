package mongodb

import (
	"context"
	"os"

	"github.com/pedro-phd/first-api-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func InitConnection() {

	logger.Info("Init Database Connection", zap.String("journey", "InitConnection"))

	dburl := os.Getenv("DB_URL")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))

	if err != nil {
		logger.Error("Database Connection failed", err, zap.String("journey", "InitConnection"))
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Database ping failed", err, zap.String("journey", "InitConnection"))
		panic(err)
	}

	logger.Info("Database Connection successfully", zap.String("journey", "InitConnection"))

}
