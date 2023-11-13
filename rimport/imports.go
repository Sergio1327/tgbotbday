package rimport

import (
	"tgbot/internal/repository/postgresql"
	"tgbot/internal/transaction"
)

type RepositoryImports struct {
	SessionManager transaction.SessionManager
	Repository     Repository
}

func NewRepositoryImports(
	sessionManager transaction.SessionManager,
) RepositoryImports {

	return RepositoryImports{
		SessionManager: sessionManager,
		Repository: Repository{
			Bday:   postgresql.NewBDay(),
		},
	}
}
