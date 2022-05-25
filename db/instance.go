package db

import (
	"soccer-manager/logger"

	"github.com/astaxie/beego/orm"
)

type DBInstance interface {
	GetReadableDB() orm.Ormer
	GetWritableDB() orm.Ormer
	GetTransactionDB() (orm.Ormer, error)
}

type dbInstance struct {
	readableDb *orm.Ormer
	writableDb *orm.Ormer

	instanceConfig PgInstanceConfig
}

const (
	DATABASE_DRIVER            = "postgres"
	READABLE_DATABASE_ALIAS    = "default"
	WRITABLE_DATABASE_ALIAS    = "master"
	TRANSACTION_DATABASE_ALIAS = "transaction_db"
)

func (d *dbInstance) GetReadableDB() orm.Ormer {
	if d.readableDb != nil {
		return *d.readableDb
	}
	logger.Log.Info("fallback registering readable DB")
	db := orm.NewOrm()
	db.Using(READABLE_DATABASE_ALIAS)
	logger.Log.Info("Connected to postgres successfully...")
	d.readableDb = &db
	return db
}

func (d *dbInstance) GetWritableDB() orm.Ormer {
	if d.writableDb != nil {
		return *d.writableDb
	}
	logger.Log.Info("fallback registering writable DB")
	db := orm.NewOrm()
	db.Using(WRITABLE_DATABASE_ALIAS)
	logger.Log.Info("Connected to postgres successfully...")
	d.writableDb = &db
	return db
}

func (d *dbInstance) GetTransactionDB() (orm.Ormer, error) {
	logger.Log.Info("Creating new transaction database connection...")
	if d.instanceConfig.BuildEnv == "dev" {
		orm.Debug = true
	}
	err := orm.RegisterDriver(DATABASE_DRIVER, orm.DRPostgres)
	if err != nil {
		return nil, err
	}
	err = orm.RegisterDataBase(TRANSACTION_DATABASE_ALIAS, DATABASE_DRIVER, d.instanceConfig.ConnURL)
	if err != nil {
		if err.Error() != "DataBase alias name `transaction_db` already registered, cannot reuse" {
			return nil, err
		}
	}

	db := orm.NewOrm()
	db.Using(TRANSACTION_DATABASE_ALIAS)
	logger.Log.Info("New transaction database connection created...")

	return db, nil
}

func (*dbInstance) registerDatabase(databaseAlias, databaseDriver, databaseConnectionUrl string) *orm.Ormer {
	logger.Log.Infof("registering %v DB", databaseAlias)
	err := orm.RegisterDataBase(databaseAlias, databaseDriver, databaseConnectionUrl)
	if err != nil {
		logger.Log.Fatal(err)
	}
	db := orm.NewOrm()
	db.Using(databaseAlias)
	logger.Log.Infof("registered %v DB", databaseAlias)
	return &db
}

func (d *dbInstance) initialize() {
	// Initializing default database
	logger.Log.Info("Connecting to postgres...")
	if d.instanceConfig.BuildEnv == "dev" {
		orm.Debug = true
	}
	err := orm.RegisterDriver(DATABASE_DRIVER, orm.DRPostgres)
	if err != nil {
		logger.Log.Fatal(err)
	}

	d.readableDb = d.registerDatabase(READABLE_DATABASE_ALIAS, DATABASE_DRIVER, d.instanceConfig.ConnURL)
	d.writableDb = d.registerDatabase(WRITABLE_DATABASE_ALIAS, DATABASE_DRIVER, d.instanceConfig.ConnURL)

	logger.Log.Info("Connected to postgres successfully...")
}

func NewDBInstance(instanceConfig PgInstanceConfig) DBInstance {
	i := &dbInstance{
		instanceConfig: instanceConfig,
	}
	i.initialize()
	return i
}
