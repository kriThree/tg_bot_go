package middlewares

import (
	"english_learn/internal/bot/controller/utils"
	"english_learn/internal/storage"
	"errors"
	"log/slog"
)

func (m *BotMiddlewares) PutUser(ctx utils.AppContext) error {

	const op = "middlewares.PutUser"

	log := m.l.With(slog.String("op", op))

	user, err := m.mu.UserTgInteraction(ctx, ctx.State.TgID)

	if err != nil {
		if errors.Is(err, storage.UserAlreadyAddedErr) {
			return nil
		}
		return err
	}

	log.Info("User", slog.Any("user", user))

	ctx.State.DbId = user.ID

	return nil
}
