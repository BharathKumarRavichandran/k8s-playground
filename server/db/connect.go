package db

import (
	"github.com/gocql/gocql"

	"github.com/BharathKumarRavichandran/k8s-playground/server/utils"
)

var Session *gocql.Session

func Init(config *utils.Config) {
	// Connect to database
	if err := open(config); err != nil {
		panic("Could not connect to database")
	}
}

func open(config *utils.Config) error {

	var err error

	cluster := gocql.NewCluster(config.DB_HOST)
	cluster.Port = config.DB_PORT
	cluster.Keyspace = config.DB_KEYSPACE
	cluster.Consistency = gocql.Quorum // Use gocql.All if you want to achieve strong consistency (W->All nodes)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.DB_USERNAME, Password: config.DB_PASSWORD}

	Session, err = cluster.CreateSession()
	if err != nil {
		utils.Logger.Fatal(err.Error())
		return err
	}
	utils.Logger.Info("Cassandra init done")

	utils.Logger.Infof("Database connection successful to %s:%d", config.DB_HOST, config.DB_PORT)
	return err
}

func Close() {
	utils.Logger.Error("Database connection closed")
	Session.Close()
}
