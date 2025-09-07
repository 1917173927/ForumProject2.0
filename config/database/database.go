package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/database")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// Debug log loaded config
	fmt.Printf("Loaded database config: host=%s, port=%d, username=%s\n",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.username"))

	// Validate required config fields
	required := []string{"database.host", "database.port", "database.username", "database.password"}
	for _, field := range required {
		if !viper.IsSet(field) {
			panic(fmt.Errorf("missing required config field: %s", field))
		}
	}
}

func GetDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/items?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetInt("database.port"))

	if viper.GetString("database.host") == "" {
		return nil, fmt.Errorf("database host cannot be empty")
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
}
