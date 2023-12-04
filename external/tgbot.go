package external

import (
	"tgbot/internal/transaction"
	"tgbot/tools/genproxy"
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

func (e *BOT) logPrefix() string {
	return "[cron_external]"
}

func (e *BOT) RunBOT(sm transaction.SessionManager) {
	e.log.Infoln("бот запущен")
	currentTime := time.Now()

	genproxy.TsWrapperForUsecase(e.log, e.SessionManager, func(ts transaction.Session) error {
		return e.Usecase.BDay.LoadBDays(ts,currentTime)
	}, e.logPrefix())
}
