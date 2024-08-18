package records

import "gorm.io/gorm"

func (rri *RecordsRepoImpl) SetDbInstance(db *gorm.DB) {
	rri.db = db
}
