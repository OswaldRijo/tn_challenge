package operations

import "gorm.io/gorm"

func (ori *OperationsRepoImpl) SetDbInstance(db *gorm.DB) {
	ori.db = db
}
