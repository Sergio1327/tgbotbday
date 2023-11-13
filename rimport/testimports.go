package rimport

import (
	"tgbot/internal/repository"
	"tgbot/internal/transaction"

	"go.uber.org/mock/gomock"
)

type TestRepositoryImports struct {
	SessionManager *transaction.MockSessionManager
	MockRepository MockRepository
	ctrl           *gomock.Controller
}

func NewTestRepositoryImports(
	ctrl *gomock.Controller,
) TestRepositoryImports {

	return TestRepositoryImports{
		ctrl: ctrl,

		SessionManager: transaction.NewMockSessionManager(ctrl),
		MockRepository: MockRepository{
			Bday:   repository.NewMockBday(ctrl),
		},
	}
}

func (t *TestRepositoryImports) MockSession() *transaction.MockSession {
	ts := transaction.NewMockSession(t.ctrl)

	ts.EXPECT().Start().Return(nil).AnyTimes()
	ts.EXPECT().Rollback().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) MockSessionWithCommit() *transaction.MockSession {
	ts := t.MockSession()

	ts.EXPECT().Commit().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) RepositoryImports() RepositoryImports {
	return RepositoryImports{
		SessionManager: t.SessionManager,
		Repository: Repository{
			Logger: t.MockRepository.Logger,
			Bday:   t.MockRepository.Bday,
		},
	}
}
