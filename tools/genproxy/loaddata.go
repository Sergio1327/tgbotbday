package genproxy

import (
	"fmt"
	"tgbot/internal/entity/global"
	"tgbot/internal/transaction"

	"github.com/sirupsen/logrus"
)

type LoadDataFunc[T any] func() (T, error)
type ReturnErrFunc func(ts transaction.Session) error

func TsWrapperForUsecase(
	log *logrus.Logger,
	sessionManager transaction.SessionManager,
	f ReturnErrFunc,
	logPrefix string,
) {
	ts := sessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		log.Errorln(logPrefix, fmt.Sprintf("не удалось открыть транзакцию, ошибка: %v", err))
		return
	}
	defer ts.Rollback()
	err := f(ts)
	if err != nil {
		return
	}
	if err := ts.Commit(); err != nil {
		log.Errorln(logPrefix, fmt.Sprintf("не удалось закоммитить транзакцию, ошибка: %v", err))
		return
	}
}

func LoadData[T any](f LoadDataFunc[T], l *logrus.Logger, errMsg ...interface{}) (T, error) {
	data, err := f()
	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		errMsg = append(errMsg, []interface{}{"ошибка:", err}...)
		l.Errorln(errMsg...)
		return data, global.ErrInternalError
	}
}

func LoadDataWithFields[T any](f LoadDataFunc[T], l *logrus.Logger, fields logrus.Fields, errMsg ...interface{}) (T, error) {
	data, err := f()
	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		errMsg = append(errMsg, []interface{}{"ошибка:", err}...)
		l.WithFields(fields).Errorln(errMsg...)
		return data, global.ErrInternalError
	}
}

func LoadDataWithCustomError[T any](f LoadDataFunc[T], noDataError error, l *logrus.Logger, fields logrus.Fields, errMsg ...interface{}) (T, error) {
	data, err := f()
	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, noDataError
	default:
		errMsg = append(errMsg, []interface{}{"ошибка:", err}...)
		l.WithFields(fields).Errorln(errMsg...)
		return data, global.ErrInternalError
	}
}

func LoadDataCanNoData[T any](f LoadDataFunc[T], l *logrus.Logger, errMsg ...interface{}) (T, error) {
	data, err := f()
	switch err {
	case nil, global.ErrNoData:
		return data, nil
	default:
		errMsg = append(errMsg, []interface{}{"ошибка:", err}...)
		l.Errorln(errMsg...)
		return data, global.ErrInternalError
	}
}
