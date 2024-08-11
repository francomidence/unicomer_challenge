package holidays

import "unicomer_challenge/models"

// SetMockHolidays Used for testing
func SetMockHolidays(mockData []models.Holiday) {
	holidays = mockData
}
