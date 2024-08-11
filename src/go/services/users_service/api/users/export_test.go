package users

import "truenorth/services/users_service/repositories/users"

func (uci *UsersApiImpl) SetUserMock(repo users.UsersRepo) {
	uci.usersRepository = repo
}
