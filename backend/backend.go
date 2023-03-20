package server

import (
	"github.com/FISCO-BCOS/go-sdk/canal"
	encrypter "github.com/FISCO-BCOS/go-sdk/encryption"
	"github.com/FISCO-BCOS/go-sdk/redis"
	sql "github.com/FISCO-BCOS/go-sdk/sqlController"
)

type Server struct {
	encrypte *encrypter.Encrypter
	sql      *sql.SqlCtr
	redis    *redis.RedisOperator
	canal    *canal.Connector
}
