// Package crontab: 基金股票持仓
package crontab

import (
	"fmt"
	"math"
	"shershon1991/fund-analye-system/global"
	"shershon1991/fund-analye-system/model/dao"
	"shershon1991/fund-analye-system/model/entity"
	"shershon1991/fund-analye-system/service/crawl/fund"
	"sync"
)

type FundStockCron struct {
}

// 声明并发等待组
var wg sync.WaitGroup

// 每次任务抓取总数量
var perTaskTotal = 50

// 定时任务启动入口
func (c FundStockCron) Run() {
	fmt.Println("基金持仓-股票定时任务准备执行....")
	pageSize := 10
	page := int(math.Ceil(float64(perTaskTotal) / float64(pageSize)))
	// 开启协程分组抓取
	// 创建数据通道
	var dataChan = make(chan [][]entity.FundStock, perTaskTotal/pageSize)
	// 记录每次任务对应的基金code
	var fundCodeChannel = make(chan []string, perTaskTotal)
	runWithGoroutine(dataChan, fundCodeChannel, page, pageSize)
	// 读取通道，数据入库
	saveToDb(dataChan, fundCodeChannel)
	defer close(dataChan)
	defer close(fundCodeChannel)
}

// 开启协程分组抓取
func runWithGoroutine(dataChan chan [][]entity.FundStock, fundCodeChannel chan []string, page, pageSize int) {
	// 开启协程抓取
	wg.Add(page)
	for i := 1; i <= page; i++ {
		pageTmp := i
		go func() {
			// 获取对应页数的code
			fundStocks, err := dao.FindNoSyncFundStockByPage(pageTmp, pageSize)
			if err == nil {
				var fundStockList [][]entity.FundStock
				var fundCodes []string
				for _, val := range fundStocks {
					rows := &fund.StockPercentageRowsCrawl{}
					rows.CrawlHtml(val.Code)
					fundCodes = append(fundCodes, val.Code)
					if len(rows.Rows) > 0 {
						convertEntity := rows.ConvertEntity()
						fundStockList = append(fundStockList, convertEntity)
					}
				}
				// 数据存入通道
				dataChan <- fundStockList
				fundCodeChannel <- fundCodes
			}
			// 并发等待组减一
			wg.Done()
		}()
	}
	wg.Wait()
}

// 保存入库
func saveToDb(dataChan chan [][]entity.FundStock, fundCodeChannel chan []string) {
	// 声明基金持仓股票实体列表
	fundStockRows := []entity.FundStock{}
	// 声明股票实体列表
	stockRows := []entity.Stock{}
	// 声明股票实体列表
	checkExistKey := make(map[string]struct{}, perTaskTotal)
	// 遍历
	for fundStockGroup := range dataChan {
		for _, fundStockList := range fundStockGroup {
			for _, fundStock := range fundStockList {
				stockCode := fundStock.StockCode
				fundStockRows = append(fundStockRows, fundStock)
				// 判断是否已经存在
				if _, ok := checkExistKey[stockCode]; !ok {
					stockRows = append(stockRows, entity.Stock{
						Code:         fundStock.StockCode,
						Name:         fundStock.StockName,
						ExchangeCode: fundStock.StockExchange,
					})
					checkExistKey[stockCode] = struct{}{}
				}
			}
		}
	}
	var codeList []string
	for val := range fundCodeChannel {
		for _, c := range val {
			codeList = append(codeList, c)
		}
	}

	if len(fundStockRows) > 0 {
		// 入库
		if save := global.GvaMysqlClient.Create(fundStockRows); save.Error != nil {
			global.GvaLogger.Sugar().Errorf("基金持仓入库失败:%s", save.Error)
		}
	}
	// 批量更新
	if len(codeList) > 0 {
		if up := global.GvaMysqlClient.Model(&entity.FundBasis{}).Where("`code` in ?", codeList).
			Update("sync_stock", 1); up.Error != nil {
			global.GvaLogger.Sugar().Errorf("信息更新失败:%s", up.Error)
		}
	}
	if len(stockRows) > 0 {
		if save := global.GvaMysqlClient.Create(stockRows); save.Error != nil {
			global.GvaLogger.Sugar().Errorf("股票信息入库失败:%s", save.Error)
		}
	}
}
