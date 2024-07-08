package test

import (
	"context"
	"english_learn/internal/config"
	log "english_learn/internal/lib/logs"
	user_service "english_learn/internal/service/userService"
	storageLib "english_learn/internal/storage"
	"english_learn/internal/storage/sqlite"
	"errors"
	"log/slog"
	"testing"
)

func TestUserService(t *testing.T) {
	conf := config.MustLoadByPath("../../config/local.yaml")

	log := log.LogInitializer(conf.Env)
	t.Log("test", slog.Any("conf", conf))

	storage, err := sqlite.New("../../storage/storage.db")

	if err != nil {
		panic(err)
	}

	user_service := user_service.New(log, storage.User)

	res, err := user_service.AddUser(context.Background(), 1121213)

	if err != nil {

		if errors.Is(err, storageLib.UserAlreadyAddedErr) {
			t.Log("User already added")
		} else {
			panic(err)
		}

	}

	log.Info(res)

}
