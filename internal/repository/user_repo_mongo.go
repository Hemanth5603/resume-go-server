package repository

import (
	"context"
	"time"

	"github.com/Hemanth5603/resume-go-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

// NewUserRepositoryMongo creates a new MongoDB user repository
func NewUserRepositoryMongo(db *mongo.Database) *UserRepositoryMongo {
	collection := db.Collection("users")
	
	// Create indexes for unique fields
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// Create unique index on username
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	
	// Create unique index on email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	
	// Create indexes
	collection.Indexes().CreateMany(ctx, []mongo.IndexModel{usernameIndex, emailIndex})
	
	return &UserRepositoryMongo{
		collection: collection,
	}
}

// CreateUser inserts a new user into MongoDB
func (r *UserRepositoryMongo) CreateUser(user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	
	// Generate new ObjectID if not set
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}
	
	// Insert user
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

// GetUserByID retrieves a user by their ID
func (r *UserRepositoryMongo) GetUserByID(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	var user model.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// GetUserByEmail retrieves a user by their email
func (r *UserRepositoryMongo) GetUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// GetUserByUsername retrieves a user by their username
func (r *UserRepositoryMongo) GetUserByUsername(username string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// UpdateUser updates an existing user
func (r *UserRepositoryMongo) UpdateUser(id string, user *model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	
	// Update timestamp
	user.UpdatedAt = time.Now()
	
	update := bson.M{
		"$set": bson.M{
			"username":          user.Username,
			"name":              user.Name,
			"email":             user.Email,
			"profile_image_url": user.ProfileImageURL,
			"updated_at":        user.UpdatedAt,
		},
	}
	
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return nil, err
	}
	
	return r.GetUserByID(id)
}

// DeleteUser deletes a user by their ID
func (r *UserRepositoryMongo) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	
	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// ListUsers retrieves all users with pagination
func (r *UserRepositoryMongo) ListUsers(page, limit int64) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	skip := (page - 1) * limit
	opts := options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.D{{Key: "created_at", Value: -1}})
	
	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	
	var users []*model.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	
	return users, nil
}

// CountUsers returns the total number of users
func (r *UserRepositoryMongo) CountUsers() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	return r.collection.CountDocuments(ctx, bson.M{})
}

