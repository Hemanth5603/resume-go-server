package database

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
	database    *mongo.Database
)

// ConnectMongoDB initializes the MongoDB connection.
// It uses a singleton pattern to ensure only one client is created.
func ConnectMongoDB(databaseURL string) (*mongo.Database, error) {
	var err error
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(databaseURL)

		// Set connection pool settings
		clientOptions.SetMaxPoolSize(100)
		clientOptions.SetMinPoolSize(10)
		clientOptions.SetMaxConnIdleTime(30 * time.Second)

		// Configure TLS for MongoDB Atlas
		tlsConfig := &tls.Config{
			InsecureSkipVerify: false,
		}
		clientOptions.SetTLSConfig(tlsConfig)

		// Set server selection timeout
		clientOptions.SetServerSelectionTimeout(10 * time.Second)
		clientOptions.SetConnectTimeout(10 * time.Second)

		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Connect to MongoDB
		mongoClient, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			return
		}

		// Ping the database to verify connection
		err = mongoClient.Ping(ctx, nil)
		if err != nil {
			return
		}

		// Get database name from connection string or use default
		// MongoDB driver automatically extracts database name from URI
		database = mongoClient.Database("resumedb")
	})

	if err != nil {
		return nil, err
	}

	return database, nil
}

// GetMongoDB returns the existing MongoDB database instance.
// It's recommended to call ConnectMongoDB once at the application start.
func GetMongoDB() *mongo.Database {
	return database
}

// GetMongoClient returns the MongoDB client for advanced operations.
func GetMongoClient() *mongo.Client {
	return mongoClient
}

// DisconnectMongoDB closes the MongoDB connection.
// Should be called when the application shuts down.
func DisconnectMongoDB() error {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return mongoClient.Disconnect(ctx)
	}
	return nil
}

// GetCollection returns a MongoDB collection by name.
func GetCollection(collectionName string) *mongo.Collection {
	if database == nil {
		return nil
	}
	return database.Collection(collectionName)
}
