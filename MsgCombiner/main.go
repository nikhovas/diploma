package main

import (
	"MsgCombiner/Application"
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/nikhovas/diploma/proto/data/actionEvent"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/sync/semaphore"
	"log"
)

func stateManager(messages []string, botService string, groupId string, userId string) {

}

func main() {
	const maxWorkers = 2000

	var app Application.Application

	if err := app.SetUpCoordinator(); err != nil {
		log.Fatal(err)
	}

	if err := app.SetUpAmqp(); err != nil {
		log.Fatal(err)
	}

	sem := semaphore.NewWeighted(int64(maxWorkers))
	ctx := context.TODO()

	msgs, _ := amqpChannel.Consume(amqpQueue.Name, "", true, false, false, false, nil)

	for {
		err := sem.Acquire(ctx, 1)
		if err != nil {
			return
		}

		select {
		case d := <-msgs:
			var ae actions.ActionEvent
			err := jsonpb.Unmarshal(bytes.NewReader(d.Body), &ae)
			if err != nil {
				continue
			}

			msgText := answer.Answer
			if msgText == "" {
				msgText = gotrans.T("There is no answer for your question yet. The question was sent to our support team")
				answer.Operator = true
				err = tm.operatorChannel.Publish("", tm.OperatorQueue, false, false,
					amqp.Publishing{
						ContentType: "application/json",
						Body: []byte(fmt.Sprintf(`{"queue": "%s", "metadata": {"chatId": %d, "message": %d}, "question": "%s", "badAnswer": ""}`,
							tm.QuestionAnswerChannel, answer.Metadata.ChatId, answer.Metadata.Message, answer.Question)),
					})
			} else if answer.Operator {
				msgText = "*" + gotrans.T("Operator answer") + "*\n" + msgText
			}

			answerMsg := tgbotapi.NewMessage(answer.Metadata.ChatId, msgText)
			answerMsg.ReplyToMessageID = answer.Metadata.Message
			if !answer.Operator {
				inlineText := fmt.Sprintf("BAD:%d", answer.Metadata.Message)
				answerMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(gotrans.T("Bad answer"), inlineText)))
			}
			answerMsg.ParseMode = tgbotapi.ModeMarkdown
			_, _ = tm.bot.Send(answerMsg)
		case _ = <-stop:
			break loop1
		}
	}
}
