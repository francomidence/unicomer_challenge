package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"unicomer_challenge/models"
	"unicomer_challenge/services/holidays"
)

func TestInitializeHolidayData(t *testing.T) {
	testCases := []struct {
		TestDescription string
		TestFunction    func(t *testing.T)
	}{
		{
			TestDescription: "Happy path - initialize holiday data",
			TestFunction: func(t *testing.T) {
				c, _ := gin.CreateTestContext(nil) // Create a gin context for testing

				service := holidays.NewHolidaysService()

				// Initialize the holiday data
				err := service.InitializeHolidayData(c)

				// Assertions to verify expected outcomes
				assert.Nil(t, err)

				// Now check that holidays were initialized
				holidaysData := service.GetHolidays(c)
				assert.NotEmpty(t, holidaysData)
			},
		},
	}

	runTestCases(t, testCases)
}

func TestGetFilteredHolidays(t *testing.T) {
	c, _ := gin.CreateTestContext(nil) // Create a gin context for testing

	service := holidays.NewHolidaysService()

	// Mock holiday data for testing
	mockHolidays := []models.Holiday{
		{Date: "2024-01-01", Title: "Año Nuevo", Type: "Civil", Inalienable: true, Extra: "Civil e Irrenunciable"},
		{Date: "2024-03-29", Title: "Viernes Santo", Type: "Religioso", Inalienable: false, Extra: "Religioso"},
		{Date: "2024-05-01", Title: "Día Nacional del Trabajo", Type: "Civil", Inalienable: true, Extra: "Civil e Irrenunciable"},
		{Date: "2024-12-25", Title: "Navidad", Type: "Religioso", Inalienable: true, Extra: "Religioso e Irrenunciable"},
	}

	// Manually set the holidays to the mock data
	holidays.SetMockHolidays(mockHolidays)

	testCases := []struct {
		TestDescription string
		FilterType      string
		StartDate       string
		EndDate         string
		ExpectedCount   int
	}{
		{
			TestDescription: "Filter by type Civil",
			FilterType:      "Civil",
			StartDate:       "",
			EndDate:         "",
			ExpectedCount:   2,
		},
		{
			TestDescription: "Filter by type Religioso",
			FilterType:      "Religioso",
			StartDate:       "",
			EndDate:         "",
			ExpectedCount:   2,
		},
		{
			TestDescription: "Filter by date range",
			FilterType:      "",
			StartDate:       "2024-01-01",
			EndDate:         "2024-05-01",
			ExpectedCount:   3,
		},
		{
			TestDescription: "Filter by type and date range",
			FilterType:      "Religioso",
			StartDate:       "2024-01-01",
			EndDate:         "2024-12-31",
			ExpectedCount:   2,
		},
	}

	runTestCasesWithArgs(t, testCases, func(tc struct {
		TestDescription string
		FilterType      string
		StartDate       string
		EndDate         string
		ExpectedCount   int
	}) {
		startDate, _ := time.Parse("2006-01-02", tc.StartDate)
		endDate, _ := time.Parse("2006-01-02", tc.EndDate)

		filteredHolidays := service.GetFilteredHolidays(c, tc.FilterType, startDate, endDate)

		assert.Equal(t, tc.ExpectedCount, len(filteredHolidays))
	})
}

func runTestCases(t *testing.T, testCases []struct {
	TestDescription string
	TestFunction    func(t *testing.T)
}) {
	for _, testCase := range testCases {
		t.Run(testCase.TestDescription, func(t *testing.T) {
			testCase.TestFunction(t)
		})
	}
}

func runTestCasesWithArgs(t *testing.T, testCases []struct {
	TestDescription string
	FilterType      string
	StartDate       string
	EndDate         string
	ExpectedCount   int
}, testFunc func(tc struct {
	TestDescription string
	FilterType      string
	StartDate       string
	EndDate         string
	ExpectedCount   int
})) {
	for _, tc := range testCases {
		t.Run(tc.TestDescription, func(t *testing.T) {
			testFunc(tc)
		})
	}
}
