package common

import (
	"fmt"
	"strconv"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
)


type configuration struct {
	DbUser, DbPwd, Database, DbHost      string
}

var ServiceConfig configuration

func initConfig() {
	ServiceConfig.DbUser = "ticket_management"
	ServiceConfig.DbPwd = "ticket_management"
	ServiceConfig.Database = "ticket_management"
	ServiceConfig.DbHost = "localhost"
}

var (
	Db         *gorm.DB
	Log        *logrus.Logger
)


func createDb() {
	if Db == nil {
		var err error
		dns := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
			ServiceConfig.DbHost, ServiceConfig.DbUser, ServiceConfig.Database, ServiceConfig.DbPwd)
		Db, err = gorm.Open("postgres", dns)
		if err != nil {
			Log.Panicf("unable to connect to the %s database: %s", ServiceConfig.Database, err.Error())
		}
		Log.Debugf("Successfully connected to database '%s'", ServiceConfig.Database)
	}
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	n, err := migrate.Exec(Db.DB(), "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err.Error())
	}
	Log.Debugf("Applied %d migrations", n)
}

func CreateLog() {
	if Log == nil {
		Log = logrus.New()
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&nested.Formatter{
			HideKeys:    false,
			FieldsOrder: []string{"handler", "issue"},
			NoColors:    true,
		})
	}
}

func ParseUID(url string) (uint, error) {
	s := strings.Split(url, "/")
	sub := s[len(s)-1]
	id, err := strconv.ParseInt(sub, 10, 32)
	if err != nil {
		return 0, nil
	}
	return uint(id), nil
}
