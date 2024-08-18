package balances

import "gorm.io/gorm"

func (bri *BalancesRepoImpl) SetDbInstance(db *gorm.DB) {
	bri.db = db
}
