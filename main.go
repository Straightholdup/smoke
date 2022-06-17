package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math"
	"time"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5448338135:AAHYJKjuL0WmR2t8aVszY3dE3ZanGV8h4oo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	now := time.Now()
	schedule := [2]time.Time{
		time.Date(now.Year(), now.Month(), now.Day(), 11, 00, 00, 00, now.Location()),
		time.Date(now.Year(), now.Month(), now.Day(), 18, 00, 00, 00, now.Location()),
	}
	minRestMinutes := int(math.Abs(schedule[0].Sub(now).Minutes()))
	for _, date := range schedule[1:] {
		restTime := int(math.Abs(date.Sub(now).Minutes()))
		if minRestMinutes > restTime {
			minRestMinutes = restTime
		}
	}
	hours := "hours"
	minutes := "minutes"
	if int(minRestMinutes/60) == 1 {
		hours = "hour"
	}
	if minRestMinutes == 1 {
		minutes = "minutes"
	}

	chatId := int64(678055739)
	msgText := fmt.Sprintf("%d %s and %d %s before smoking", minRestMinutes/60, hours, minRestMinutes%60, minutes)
	if minRestMinutes/60 == 0 {
		msgText = fmt.Sprintf("%d %s before smoking", minRestMinutes%60, minutes)
	}

	msg := tgbotapi.NewMessage(chatId, msgText)
	bot.Send(msg)
}
