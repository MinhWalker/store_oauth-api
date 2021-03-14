package db

import (
	"github.com/MinhWalker/store_oauth-api/src/clients/cassandra"
	"github.com/MinhWalker/store_oauth-api/src/domain/access_token"
	"github.com/MinhWalker/store_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(string) (*access_token.AccessToken, *errors.RestErr) {
	session , err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// TODO: implement get access token from CassandraDB
	return nil, errors.NewInternalServerError("database connection not implemented yet!")
}