package uimport

import (
	"tgbot/internal/usecase"
)

type Usecase struct {
	Logger *usecase.LoggerUsecase
	BDay   *usecase.BdayUsecase
}
