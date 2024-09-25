package pgress

import (
	"strconv"
	"time"

	"github.com/vishal2911/algoTrading/angelOneGo/model"
	"github.com/vishal2911/algoTrading/angelOneGo/util"
	"gorm.io/gorm"
)

func (store PgressStore) setLimitAndPage(filter map[string]string, query *gorm.DB) {
	// Convert limit and page to integers
	limit := filter[model.DataPerPage]
	page := filter[model.DataPerPage]
	if limit == "" {
		limit = "10" // Default limit is 10
	}
	if page == "" {
		page = "1" // Default page is 1
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.SetLimitAndPage,
			"error while while converting limit to int", err)
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.SetLimitAndPage,
			"error while while converting page to int", err)
	}

	// Apply pagination
	offset := (pageInt - 1) * limitInt
	query = query.Limit(limitInt).Offset(offset)

}

func (store PgressStore) setDateRangeFilter(filter map[string]string, query *gorm.DB) {
	// Convert limit and page to integers
	startDatestr := filter[model.StartDate]
	endDatestr := filter[model.EndDate]
	if startDatestr == "" || endDatestr == "" {
		util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.SetDateRangeFilter,
			"Date range not provided", nil)
		return
	}
	// Parse the start and end dates
	startDate, err := time.Parse(model.TimeLayout, startDatestr)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.SetDateRangeFilter,
			"error in parsing start date", err)
		return
	}
	endDate, err := time.Parse(model.TimeLayout, endDatestr)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.SetDateRangeFilter,
			"error in parsing end date", err)
		return
	}
	query.Where("date_column BETWEEN ? AND ?", startDate, endDate)
}
