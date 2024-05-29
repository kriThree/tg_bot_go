package bot

import (
	"english_learn/internal/bot/handlers"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	l     *slog.Logger
	token string
	du    handlers.DefinitionUscase
}

func (b *Bot) Run() error {
	log := b.l.With(slog.String("method", "Run"))

	log.Info("Starting bot")

	bot, err := tgbotapi.NewBotAPI(b.token)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		panic(err)
	}

	log.Info("Rub run bot")

	router := createRouter(handlers.New(b.l, b.du, bot))

	for update := range updates {
		router(update)
	}

	return nil
}

func New(token string, l *slog.Logger, du handlers.DefinitionUscase) *Bot {
	return &Bot{token: token, l: l, du: du}
}
