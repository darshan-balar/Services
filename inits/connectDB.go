package inits

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure: "username"`
	Password string `mapstructure: "password"`
	Database string `mapstructure: "database"`
}

func ConnectDB() error {

	var configDB DbConfig

	err := viper.UnmarshalKey(ENV, &configDB)
	if err != nil {
		return err
	}
	connectionString := "host=" + configDB.Host + " user=" + configDB.Username + " dbname=" + configDB.Database + " sslmode=disable" + " password=" + configDB.Password
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}
	return nil
}
