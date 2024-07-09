package utils

import (
	"context"
	statemanager "english_learn/internal/bot/stateManager"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const EVENT_SEPARATOR = "|"

type Handler func(appCtx AppContext)

type AppContext struct {
	Ctx    context.Context
	Update tgbotapi.Update
	State  *statemanager.UserState
	value  map[interface{}]interface{}
	done   chan struct{}
}

func (m AppContext) Deadline() (deadline time.Time, ok bool) {
	ok = true
	deadline = time.Now().Add(10 * time.Minute)
	return
}

func (m AppContext) Done() <-chan struct{} {
	return m.done
}

func (m AppContext) Err() error {
	return nil
}

func (m AppContext) Value(key interface{}) interface{} {
	if _, ok := m.value[key]; !ok {
		return nil
	}
	return m.value[key]
}

func NewAppContext(ctx context.Context, update tgbotapi.Update, state *statemanager.UserState) AppContext {
	return AppContext{
		Ctx:    ctx,
		Update: update,
		State:  state,
	}
}
