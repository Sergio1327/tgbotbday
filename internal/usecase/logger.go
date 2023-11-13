package usecase

import (
	"tgbot/internal/entity/log"
	"tgbot/rimport"

	"github.com/sirupsen/logrus"
)

type LoggerUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
}

func NewLogger(log *logrus.Logger,
	ri rimport.RepositoryImports,
) *LoggerUsecase {
	return &LoggerUsecase{
		log:               log,
		RepositoryImports: ri,
	}
}

func (u *LoggerUsecase) SpecialFields() []string {
	return []string{"c_id", "se_id", "oper_login"}
}

func (u *LoggerUsecase) SaveLog(row log.Row) error {
	ts := u.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		u.log.Errorln("не удается стартовать транзакцию", err)
		return err
	}
	defer ts.Rollback()

	if err := u.Repository.Logger.SaveLog(ts, row.Message); err != nil {
		u.log.Errorln("не удается сохранить данные в детали лога", err)
		return err
	}

	if err := ts.Commit(); err != nil {
		u.log.Errorln("не удается зафиксировать транзакцию", err)
		return err
	}

	return nil
}
