package repository

import (
	"tgbot/internal/entity/user"
	"tgbot/internal/transaction"
	"time"
)

type Bday interface {
	FindAllBdayPeople(ts transaction.Session, sundayTime, lastWeekTime time.Time) ([]user.User, error)
	UpdateBday(ts transaction.Session, id int, newBdayTime time.Time) error
}

type Logger interface {
	SaveLog(ts transaction.Session, logtext string) error
}
