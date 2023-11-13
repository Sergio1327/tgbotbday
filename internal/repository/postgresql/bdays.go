package postgresql

import (
	"tgbot/internal/entity/user"
	"tgbot/internal/repository"
	"tgbot/internal/transaction"
	"tgbot/tools/gensql"
	"time"
)

type bdayRepository struct{}

func NewBDay() repository.Bday {
	return &bdayRepository{}
}

func (r *bdayRepository) FindAllBdayPeople(ts transaction.Session, sundayTime, lastWeektime time.Time) ([]user.User, error) {
	query := `
	SELECT id, name, bday
	FROM Kids
	WHERE bday BETWEEN $1 and $2`

	return gensql.Select[user.User](SqlxTx(ts), query, lastWeektime, sundayTime)
}

func (r *bdayRepository) UpdateBday(ts transaction.Session, id int, newBdayTime time.Time) error {
	query := `
	UPDATE Kids
	SET bday = $1
	WHERE id = $2`

	_, err := SqlxTx(ts).Exec(query, newBdayTime, id)
	return err

}
