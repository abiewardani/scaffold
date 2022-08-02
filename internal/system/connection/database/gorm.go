package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/abiewardani/scaffold/internal/system/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gormInstance struct {
	master, slave *gorm.DB
}

// Master initialize DB for master data
func (g *gormInstance) Master() *gorm.DB {
	return g.master
}

func (g *gormInstance) Slave() *gorm.DB {
	return g.slave
}

func (g *gormInstance) Close() {}

// GormDatabase abstraction
type GormDatabase interface {
	Master() *gorm.DB
	Slave() *gorm.DB
	Close()
}

// InitGorm ...
func InitGorm(config *config.Config) GormDatabase {
	inst := new(gormInstance)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		Logger:                 dbLogger,
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	}

	dsnMaster := fmt.Sprintf("host=%s user=%s "+
		"password=%s port=%v dbname=%s sslmode=disable",
		config.DbMasterHost, config.DbMasterUser, config.DbMasterPassword, config.DbMasterPort, config.DbMasterName)

	dbMaster, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsnMaster,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), gormConfig)
	if err != nil {
		panic("Database Connection Failed")
	}

	sqlDB, _ := dbMaster.DB()
	if err = sqlDB.Ping(); err != nil {
		panic(err.Error())
	}

	inst.master = dbMaster

	dsnSlave := fmt.Sprintf("host=%s user=%s "+
		"password=%s port=%v dbname=%s sslmode=disable",
		config.DbSlaveHost, config.DbSlaveUser, config.DbSlavePassword, config.DbSlavePort, config.DbSlaveName)

	dbSlave, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsnSlave,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), gormConfig)
	if err != nil {
		panic("Database Connection Failed")
	}

	sqlDBSlave, _ := dbSlave.DB()
	if err = sqlDBSlave.Ping(); err != nil {
		panic(err.Error())
	}

	inst.slave = dbSlave

	return inst
}
