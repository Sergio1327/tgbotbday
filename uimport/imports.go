package uimport

import (
	"tgbot/internal/transaction"
	"tgbot/internal/usecase"
	"tgbot/rimport"
	"tgbot/tools/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type UsecaseImports struct {
	SessionManager transaction.SessionManager
	Usecase        Usecase
}

func NewUsecaseImports(
	log *logrus.Logger,
	dblog *logrus.Logger,
	tgBOT *tgbotapi.BotAPI,
	ri rimport.RepositoryImports,
	sessionManager transaction.SessionManager,
) UsecaseImports {

	ui := UsecaseImports{
		SessionManager: sessionManager,

		Usecase: Usecase{
			BDay:   usecase.NewBDay(logger.NewUsecaseLogger(log, "bday"), dblog, tgBOT, ri),
			Logger: usecase.NewLogger(logger.NewUsecaseLogger(log, "logger"), ri),
		},
	}

	return ui
}
