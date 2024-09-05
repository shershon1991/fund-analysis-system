// Package dao: 基金基本信息
package dao

import (
	"gorm.io/gorm"
	"shershon1991/fund-analye-system/global"
	"shershon1991/fund-analye-system/model/entity"
)

// 统计没有获取持仓股票的基金数量
func CountNoSyncFundStock() int64 {
	var num int64
	global.GvaMysqlClient.Model(&entity.FundBasis{}).Where("sync_stock = ?", 0).Count(&num)
	return num
}

// 分页获取没有持仓股票的基金code
func FindNoSyncFundStockByPage(page, pageSize int) ([]entity.FundBasis, error) {
	limit := (page - 1) * pageSize
	fbs := []entity.FundBasis{}
	find := global.GvaMysqlClient.Select("`code`").Where("sync_stock = ?", 0).
		Limit(pageSize).Offset(limit).Find(&fbs)
	if find.Error != nil && find.Error == gorm.ErrRecordNotFound {
		return fbs, nil
	}
	return fbs, find.Error
}
