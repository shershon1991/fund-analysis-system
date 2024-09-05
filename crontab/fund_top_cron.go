package crontab

import (
	"fmt"
	"shershon1991/fund-analye-system/global"
	"shershon1991/fund-analye-system/service/crawl/fund"
)

type FundTopCron struct{}

func (c FundTopCron) Run() {
	fmt.Println("基金排行榜-定时任务准备运行!")
	f := &fund.TopCrawlService{}
	// 爬取网页
	f.CrawlHtml()
	// 转换数据
	fundDayTopList := f.ConvertEntity()
	// 入库
	if !f.ExistTopDate() {
		result := global.GvaMysqlClient.Create(fundDayTopList)
		if result.Error != nil {
			global.GvaLogger.Sugar().Errorf("本次任务保存数据失败：%条", result.Error)
			return
		}
		global.GvaLogger.Sugar().Infof("本次任务保存数据：%d条", result.RowsAffected)
	}
	fmt.Println("基金排行榜-定时任务运行结束!")
}
