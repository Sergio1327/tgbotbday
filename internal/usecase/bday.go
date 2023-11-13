package usecase

import (
	"fmt"
	"tgbot/internal/entity/bot"
	"tgbot/internal/entity/global"
	"tgbot/rimport"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type BdayUsecase struct {
	bot   *tgbotapi.BotAPI
	log   *logrus.Logger
	dblog *logrus.Logger
	rimport.RepositoryImports
}

func NewBDay(log, dblog *logrus.Logger, bot *tgbotapi.BotAPI, ri rimport.RepositoryImports) *BdayUsecase {
	return &BdayUsecase{
		log:               log,
		dblog:             dblog,
		bot:               bot,
		RepositoryImports: ri,
	}
}

func (u *BdayUsecase) LoadBDays(currentTime time.Time) error {
	ts := u.SessionManager.CreateSession()
	ts.Start()
	defer ts.Rollback()

	lastWeek := currentTime.AddDate(0, 0, -7)
	UserList, err := u.Repository.Bday.FindAllBdayPeople(ts, currentTime, lastWeek)
	switch err {
	case nil:
		if err != nil {
			msg := tgbotapi.NewMessage(bot.DevChatID, err.Error())
			u.bot.Send(msg)
		}
	case global.ErrNoData:
		u.log.Debugln("ни у кого не было др")
		return err
	default:
		u.log.Errorln("не удалось извлечь тех у кого было др, ошибка:", err)
		return global.ErrInternalError
	}

	u.log.Infoln(UserList)

	msgText := "За прошедшую неделю др было у:"
	for i, v := range UserList {
		msgText += fmt.Sprintf("\n%d) %s \nдата др : %s \n\n", i+1, v.Name, v.BirthDate.Format("2006-01-02"))
	}

	msg := tgbotapi.NewMessage(bot.TestChatID, msgText)
	_, err = u.bot.Send(msg)
	if err != nil {
		msg := tgbotapi.NewMessage(bot.DevChatID, err.Error())
		u.bot.Send(msg)
		return global.ErrInternalError
	}

	for _, v := range UserList {
		v.SetNextYear()
		fmt.Println(v)
		if err := u.Repository.Bday.UpdateBday(ts, v.ID, v.BirthDate); err != nil {
			u.log.Errorln("не удалось обновить дату др, причина:", err)
			return global.ErrInternalError
		}
	}

	ts.Commit()
	return nil
}
