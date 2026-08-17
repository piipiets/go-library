package connection

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var (
	DBConnections *sql.DB
	err           error
)

func Initiator() {
	dbEngine := viper.GetString("DB_ENGINE")
	dsn := viper.GetString("DATABASE_URL")

	DBConnections, err = sql.Open(dbEngine, dsn)

	// check connection
	err = DBConnections.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
