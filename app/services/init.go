package services

import "api-main/config/database"

func InitServices() error {
	// Initialize database connection
	if _, err := database.GetDBConnection(); err != nil {
		return err
	}

	// TODO: Add other service initializations as needed
	return nil
}
