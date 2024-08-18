package users

import (
	pubsubmock "truenorth/mocks/packages/pubsub"
	"truenorth/services/users_service/repositories/users"
)

func (uci *UsersApiImpl) SetUserMock(repo users.UsersRepo) {
	uci.usersRepository = repo
}

func (uci *UsersApiImpl) SetProducerMock(producer *pubsubmock.Producer) {
	uci.producer = producer
}
