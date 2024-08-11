package holidays

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
	"unicomer_challenge/models"
)

var (
	//Ensures holiday data is initialized only once
	once sync.Once
	//List of holidays
	holidays []models.Holiday
)

const (
	url = "https://api.victorsanmartin.com/feriados/en.json"
)

type HolidaysService struct {
	// Here we can call another service or a database repository
}

func NewHolidaysService() *HolidaysService {
	return &HolidaysService{}
}

func (s *HolidaysService) InitializeHolidayData(c *gin.Context) error {
	requestID := c.GetString("RequestID")
	logrus.Infof("Request ID %s: Initializing holiday data", requestID)
	//Ensures data is only fetched once even if there are multiple calls
	once.Do(func() {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var holidaysResponse models.HolidaysResponse
		if err := json.NewDecoder(resp.Body).Decode(&holidaysResponse); err != nil {
			panic(err)
		}

		holidays = holidaysResponse.Data
	})
	logrus.Infof("Request ID %s: Successfully initialized holiday data with %d holidays", requestID, len(holidays))
	return nil
}

func (s *HolidaysService) GetFilteredHolidays(c *gin.Context, filterType string, startDate, endDate time.Time) []models.Holiday {
	requestID := c.GetString("RequestID")

	var filtered []models.Holiday
	for _, holiday := range holidays {
		//Parsed holiday date - must follow this format
		holidayDate, _ := time.Parse("2006-01-02", holiday.Date)

		// Apply filters
		if (filterType == "" || holiday.Type == filterType) &&
			(startDate.IsZero() || !holidayDate.Before(startDate)) && // Start date is inclusive
			(endDate.IsZero() || !holidayDate.After(endDate)) { // End date is inclusive
			filtered = append(filtered, holiday)
		}
	}
	logrus.Infof("Request ID %s: Filtering holidays by type '%s' and date range '%v' to '%v', found %d results", requestID, filterType, startDate, endDate, len(filtered))
	return filtered
}

func (s *HolidaysService) GetHolidays(c *gin.Context) []models.Holiday {
	requestID := c.GetString("RequestID")
	logrus.Infof("Request ID %s: Retrieving all holidays", requestID)
	return holidays
}
