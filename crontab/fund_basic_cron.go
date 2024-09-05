package crontab

import (
	"fmt"
	"shershon1991/fund-analye-system/service/crawl/fund"
)

type FundBasicCron struct {
	Code string
}

// 抓取详情信息
func (c FundBasicCron) Run() {
	fmt.Println("基金详情-定时任务开始运行")
	// 开始爬取
	fund.BatchBasicCrawl()
	fmt.Println("基金详情-定时任务结束")
}
