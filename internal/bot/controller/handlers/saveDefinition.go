package handlers

import (
	"context"
	statemanager "english_learn/internal/bot/stateManager"
	"english_learn/internal/domain/models"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var tags = []string{"noun", "verb", "adjective", "adverb"}

func (h *BotHandlers) SaveDefinitionQuery(ctx context.Context, update tgbotapi.Update, state *statemanager.UserState) {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Write your definition")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	state.Operation = statemanager.SAVE_DEFINITION_WAIT_NAME

	h.api.Send(msg)

}

func (h *BotHandlers) SaveDefinitionAddName(ctx context.Context, update tgbotapi.Update, state *statemanager.UserState) {
	state.Creatng.Name = update.Message.Text

	state.Operation = statemanager.SAVE_DEFINITION_WAIT_MEANING

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Write mean of your definition")

	h.api.Send(msg)
}

func (h *BotHandlers) SaveDefinitionAddMean(ctx context.Context, update tgbotapi.Update, state *statemanager.UserState) {
	state.Creatng.Mean = update.Message.Text

	btns := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, tag := range tags {
		btns = append(btns, tgbotapi.NewInlineKeyboardButtonData(tag, fmt.Sprintf("%v_%v", SAVE_DEFINITION_END, tag)))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Select tag of this pair")

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns)

	h.api.Send(msg)
}

func (h *BotHandlers) SaveDefinitionEnd(ctx context.Context, update tgbotapi.Update, state *statemanager.UserState) {

	_, err := h.hu.AddDefinition(context.TODO(), state.Creatng.Name, models.Meaning{PartOfSpeach: state.Creatng.Tag, Value: state.Creatng.Mean})

	if err != nil {

		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Definition saved")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	state.Operation = statemanager.BASE

	h.api.Send(msg)

}
