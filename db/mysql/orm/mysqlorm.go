package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysql/log"
	"time"
)

// Engine is the main struct of geeorm, manages all db sessions and transactions.
type Engine struct {
	db *sql.DB
}

// NewEngine create a instance of Engine
// connect database and ping it to test whether it's alive
func NewEngine(driver, dataSourceName string) (e *Engine, err error) {
	mysql, err := sql.Open(driver, dataSourceName)
	if err != nil {
		log.Error("数据库注册失败")
	}

	err = mysql.Ping()
	if err != nil {
		log.Error("Ping数据库注册失败", err)
	}
	//defer mysql.Close()

	mysql.SetMaxOpenConns(256)
	mysql.SetMaxIdleConns(100)
	mysql.SetConnMaxLifetime(time.Second * 30)

	_, err = fmt.Println("数据库注册success")

	e = &Engine{db: mysql}
	log.Info("Connect database success")
	return
}

// Close database connection
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// NewSession creates a new session for next operations
func (engine *Engine) NewSession() *Session {
	return New(engine.db)
}
