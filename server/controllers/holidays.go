package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"unicomer_challenge/services/holidays"
)

type HolidaysController interface {
	GetHolidays(c *gin.Context)
}

type holidaysController struct {
	holidaysService *holidays.HolidaysService
}

func NewHolidaysController(holidaysService *holidays.HolidaysService) HolidaysController {
	return &holidaysController{
		holidaysService: holidaysService,
	}
}

func (ctrl *holidaysController) GetHolidays(c *gin.Context) {
	// Initialize the data if not already done
	if len(ctrl.holidaysService.GetHolidays(c)) == 0 {
		if err := ctrl.holidaysService.InitializeHolidayData(c); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize holiday data"})
			return
		}
	}

	filterType := c.Query("type")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)

	holidays := ctrl.holidaysService.GetFilteredHolidays(c, filterType, startDate, endDate)

	if c.GetHeader("Accept") == "application/xml" {
		c.XML(http.StatusOK, holidays)
	} else {
		c.JSON(http.StatusOK, holidays)
	}
}
