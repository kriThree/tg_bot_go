package test

import (
	"context"
	user_service "english_learn/internal/service/userService"
	storageLib "english_learn/internal/storage"
	"english_learn/internal/storage/sqlite"
	"errors"
	"log/slog"
	"testing"
)

// func TestMain(t *testing.T) {
// 	conf := config.MustLoadByPath("../../config/local.yaml")
// 	t.Log("test", slog.Any("conf", conf))
// 	TestUserService(t)
// }

func TestUserService(t *testing.T) {
	storage, err := sqlite.New("../../storage/storage.db")
	if err != nil {
		panic(err)
	}

	user_service := user_service.New(slog.Default(), storage.User)

	res, err := user_service.AddUser(context.Background(), 1121213)

	if err != nil {

		if errors.Is(err, storageLib.UserAlreadyAddedErr) {
			t.Log("User already added")
		} else {
			panic(err)
		}

	}
	t.Log("res", slog.Any("res", res))
}
