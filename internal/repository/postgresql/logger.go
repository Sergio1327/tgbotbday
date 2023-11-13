package postgresql

import (
	"tgbot/internal/repository"
	"tgbot/internal/transaction"
)

type loggerRepository struct {
}

func NewLogger() repository.Logger {
	return &loggerRepository{}
}

func (r *loggerRepository) SaveLog(ts transaction.Session, logtext string) error {
	sqlQuery := `insert into log_table (log_text) values (?)`

	_, err := SqlxTx(ts).Exec(sqlQuery, logtext)

	return err
}
