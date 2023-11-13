package external

import (
	"fmt"
	"tgbot/internal/transaction"
	"tgbot/uimport"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type BOT struct {
	tgBOT *tgbotapi.BotAPI
	uimport.UsecaseImports
	log *logrus.Logger
}

func NewBOT(log *logrus.Logger, tgbot *tgbotapi.BotAPI, ui uimport.UsecaseImports) *BOT {
	return &BOT{
		log:            log,
		UsecaseImports: ui,
		tgBOT:          tgbot,
	}
}

func (e *BOT) RunBOT(sm transaction.SessionManager) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	e.log.Infoln("бот запущен")

	for {
		currentTime := time.Now()
		if currentTime.Weekday() == time.Monday && currentTime.Hour() == 16 && currentTime.Minute() == 50 {
			fmt.Println("tets")
			e.Usecase.BDay.LoadBDays(currentTime)
			time.Sleep(time.Minute)
		}
	}

}
