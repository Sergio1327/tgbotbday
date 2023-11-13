package rimport

import (
	"tgbot/internal/repository"
)

type Repository struct {
	Bday   repository.Bday
	Logger repository.Logger
}

type MockRepository struct {
	Bday   *repository.MockBday
	Logger *repository.MockLogger
}
