package bot

import (
	"english_learn/internal/bot/controller"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	l     *slog.Logger
	token string
	contr *controller.Controller
}

func (b *Bot) Run() error {

	const op = "bot.Run"

	log := b.l.With(slog.String("op", op))

	log.Info("Starting bot")

	bot, err := tgbotapi.NewBotAPI(b.token)

	if err != nil {
		return err
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		return err
	}

	go b.contr.Handle(updates, bot)

	log.Info("Rub run bot")

	select {}

	return nil
}

func New(token string, l *slog.Logger, controller *controller.Controller) *Bot {
	return &Bot{token: token, l: l, contr: controller}
}
