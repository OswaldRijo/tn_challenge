package users

import "truenorth/services/users_service/api/users"

func (uci *UsersControllerImpl) SetUserApi(api users.UserApi) {
	uci.usersApi = api
}
