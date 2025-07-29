package service

import (
	"errors"
	"log"
	"time"

	"github.com/Hemanth5603/resume-go-server/internal/model"
	"github.com/Hemanth5603/resume-go-server/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(req *model.SubscriptionRequest) (*model.Subscription, error)
}

type subscriptionServiceImpl struct {
	subscriptionRepo *repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo *repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionServiceImpl{
		subscriptionRepo: subscriptionRepo,
	}
}

func (s *subscriptionServiceImpl) CreateSubscription(req *model.SubscriptionRequest) (*model.Subscription, error) {
	// timezone, err := getTimezoneOfUser(req.UserID)
	// if err != nil {
	// 	return &model.Subscription{}, err
	// }

	timezone := "Asia/Kolkata"
	log.Println(timezone)
	subscriptionID := generateSubscriptionID()
	log.Println(subscriptionID)
	token, err := generateSubscriptionToken(req.UserID, req.Plan, timezone, subscriptionID)
	log.Println(token)
	if err != nil {
		log.Println(err)
		return &model.Subscription{}, err
	}

	subscription := &model.Subscription{
		ID:       subscriptionID,
		UserID:   req.UserID,
		Plan:     req.Plan,
		Token:    token,
		TimeZone: timezone,
	}
	log.Println(*subscription)
	_ = s.subscriptionRepo.CreateSubscriptionTable()
	return s.subscriptionRepo.CreateSubscription(subscription)
}

func (s *subscriptionServiceImpl) VerifySubscription(userID string) (bool, error) {
	subscription, err := s.subscriptionRepo.GetLatestSubscriptionOfUser(userID)
	if err != nil {
		return false, err
	}
	ok, err := validateSubscriptionToken(subscription.Token)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, errors.New("token expired")
	}
	return ok, nil
}

func getTimezoneOfUser(userID string) (string, error) {
	//to be written
	return "", nil
}

func getUTCExpiry(tz string, plan int) (time.Time, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, err
	}

	now := time.Now().In(loc)
	expiryLocal := now.AddDate(0, plan, 0)
	expiryLocal = time.Date(
		expiryLocal.Year(), expiryLocal.Month(), expiryLocal.Day(),
		23, 59, 59, 0, loc,
	)

	return expiryLocal.UTC(), nil
}

func createJWTTokenForSubscription(userID string, expiryUTC time.Time, plan int, subscriptionID string) (string, error) {
	jwtSecret := []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30") //jwt secret from the env
	claims := jwt.MapClaims{
		"sub":             userID,
		"exp":             expiryUTC.UTC(),
		"iat":             time.Now().Unix(),
		"plan":            plan,
		"subscription_id": subscriptionID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func generateSubscriptionToken(userID string, plan int, timezone string, subscriptionID string) (string, error) {
	expiryUTC, err := getUTCExpiry(timezone, plan)
	if err != nil {
		return "", err
	}

	token, err := createJWTTokenForSubscription(userID, expiryUTC, plan, subscriptionID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateSubscriptionID() string {
	return "sub_" + uuid.New().String()
}

func validateSubscriptionToken(tokenStr string) (bool, error) {
	jwtSecret := []byte("") //jwt secret from env variables
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return false, errors.New("invalid token")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return false, errors.New("token expired")
		}
	}

	return true, nil
}
