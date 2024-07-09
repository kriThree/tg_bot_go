package handlers

import (
	"context"
	"english_learn/internal/bot/controller/utils"
	statemanager "english_learn/internal/bot/stateManager"
	"english_learn/internal/domain/models"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var tags = []string{"noun", "verb", "adjective", "adverb"}

func (h *BotHandlers) SaveDefinitionQuery(ctx utils.AppContext) {

	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "Write your definition")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	ctx.State.Operation = statemanager.SAVE_DEFINITION_WAIT_NAME

	h.api.Send(msg)

}

func (h *BotHandlers) SaveDefinitionAddName(ctx utils.AppContext) {
	ctx.State.Creatng.Name = ctx.Update.Message.Text

	ctx.State.Operation = statemanager.SAVE_DEFINITION_WAIT_MEANING

	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "Write mean of your definition")

	h.api.Send(msg)
}

func (h *BotHandlers) SaveDefinitionAddMean(ctx utils.AppContext) {
	ctx.State.Creatng.Mean = ctx.Update.Message.Text

	ctx.State.Operation = statemanager.SAVE_DEFINITION_WAIT_TAG

	btns := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, tag := range tags {
		btns = append(btns, tgbotapi.NewInlineKeyboardButtonData(tag, fmt.Sprintf("%v%s%v", SAVE_DEFINITION_END, utils.EVENT_SEPARATOR, tag)))
	}

	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "Select tag of this pair")

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(btns)

	h.api.Send(msg)
}

func (h *BotHandlers) SaveDefinitionEnd(ctx utils.AppContext) {
	split := strings.Split(ctx.Update.CallbackQuery.Data, utils.EVENT_SEPARATOR)

	if len(split) != 2 {
		return
	}

	ctx.State.Creatng.Tag = split[1]

	_, err := h.hu.AddDefinition(
		context.TODO(),
		ctx.State.Creatng.Name,
		models.Meaning{PartOfSpeach: ctx.State.Creatng.Tag, Value: ctx.State.Creatng.Mean},
		ctx.State.DbId,
	)

	if err != nil {
		return
	}

	_, err = h.api.AnswerCallbackQuery(tgbotapi.CallbackConfig{
		CallbackQueryID: ctx.Update.CallbackQuery.ID,
		Text:            "Definition saved",
	})

	if err != nil {
		return
	}

	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "Definition saved")

	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	ctx.State.Operation = statemanager.BASE

	h.api.Send(msg)

}
