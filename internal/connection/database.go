package connection

import(
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"log"
	"shellrean.id/Go-RestAPI/internal/config"
)

func GetDatabase(conf config.Database) *sql.DB{
	dsn:= fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
	conf.User,
	conf.Pass,
	conf.Host,
	conf.Port,
	conf.Name,
	conf.Tz,
)
db, err := sql.Open("mysql", dsn)
if err != nil {
	log.Fatal("failed to open connection: ", err.Error())
}

err= db.Ping()
if err != nil {
	log.Fatal("failed to ping connection:  ", err.Error())
}

return db
}