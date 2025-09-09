package database

import (
	"fmt"
	"sync"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/spf13/viper"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/database")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	fmt.Printf("Loaded database config: host=%s, port=%d, username=%s\n",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.username"))

	required := []string{"database.host", "database.port", "database.username", "database.password"}
	for _, field := range required {
		if !viper.IsSet(field) {
			panic(fmt.Errorf("missing required config field: %s", field))
		}
	}
}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/items?charset=utf8mb4&parseTime=True&loc=Local",
			viper.GetString("database.username"),
			viper.GetString("database.password"),
			viper.GetString("database.host"),
			viper.GetInt("database.port"))

		if viper.GetString("database.host") == "" {
			panic("database host cannot be empty")
		}

		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic(fmt.Errorf("failed to connect to database: %w", err))
		}
	})
	return db
}
