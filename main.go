package main

import (
	"log"
	"tgbot/external"
	"tgbot/internal/entity/bot"
	"tgbot/internal/transaction"
	"tgbot/rimport"
	"tgbot/tools/logger"
	"tgbot/tools/pgdb"
	"tgbot/uimport"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	tgbot, err := tgbotapi.NewBotAPI(bot.BotToken)
	if err != nil {
		log.Fatal("не удалось подключиться к боту, причина:", err)
	}

	log := logger.NewNoFileLogger("tgbot1")
	dblog := logger.NewNoFileLogger("tgbot")

	db := pgdb.ConnectToDB()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sm := transaction.NewSQLSessionManager(db)
	ri := rimport.NewRepositoryImports(sm)

	ui := uimport.NewUsecaseImports(log, dblog, tgbot, ri, sm)

	cron := external.NewBOT(log, tgbot, ui)
	cron.RunBOT(sm)
}
