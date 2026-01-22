package utils

import (
	"authservice/models"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

var ctx=context.Background()
func AddToMessageQueue(username, email string) error {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	payload := models.UserCreationResponse{
		To:	  "assis.mohanty.98@gmail.com",
		Subject: "Welcome to AuthService",
		Body: models.CreateRequestTypeWithOutPassowrd{
			Username: username,
			Email:    email,
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return rdb.LPush(ctx, "email_queue", data).Err()
}
