package repository

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/valonekowd/clean-architecture/adapter/repository/transaction"
	"github.com/valonekowd/clean-architecture/usecase/interfaces/gateway"
)

type Config struct {
	sqlDB   *sqlx.DB
	mongoDB *mongo.Database
	// es      *elasticsearch.Client
	// rdb     *redis.Client
}

type ConfigOption func(*Config)

func WithSQLDB(db *sqlx.DB) ConfigOption {
	return func(c *Config) {
		c.sqlDB = db
	}
}

func New(options ...ConfigOption) gateway.DataSource {
	c := &Config{}

	for _, o := range options {
		o(c)
	}

	var transactionRepository gateway.TransactionDataSource
	{
		transactionRepository = &transaction.GuardRepository{}
		transactionRepository = transaction.NewSQLRepository(c.sqlDB, transactionRepository)
		transactionRepository = transaction.NewMongoRepository(c.mongoDB, transactionRepository)
	}

	return gateway.DataSource{
		Transaction: transactionRepository,
	}
}
