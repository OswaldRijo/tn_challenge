package users

import "gorm.io/gorm"

func (uri *UsersRepoImpl) SetDbInstance(db *gorm.DB) {
	uri.db = db
}
