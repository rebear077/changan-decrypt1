package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisOperator struct {
	rdb *redis.Client
}
type Rediser interface {
	Set(ctx context.Context, key string, values ...interface{}) error
	Get(ctx context.Context, key, field string) (string, error)
	GetAll(ctx context.Context, key string) (map[string]string, error)
	GetKeys(ctx context.Context, key string) ([]string, error)
	GetHashLen(ctx context.Context, key string) (int64, error)
	MultipleGet(ctx context.Context, key string, fields ...string) ([]interface{}, error)
	MultipleSet(ctx context.Context, key string, values map[string]interface{}) error
	Delete(ctx context.Context, key string, fields ...string) error
	Exists(ctx context.Context, key, field string) (bool, error)
	Scan(ctx context.Context, condition string) int
}

// 初始化一个redis客户端
func NewRedisOperator() *RedisOperator {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	return &RedisOperator{
		rdb: rdb,
	}
}

// 调用Hset方法，根据key和field字段设置，field字段的值(单个值)
// HSet accepts values in following formats:
//
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
//
//     Playing struct With "redis" tag.
//     type MyHash struct { Key1 string `redis:"key1"`; Key2 int `redis:"key2"` }
//
//   - HSet("myhash", MyHash{"value1", "value2"})
//
//     For struct, can be a structure pointer type, we only parse the field whose tag is redis.
//     if you don't want the field to be read, you can use the `redis:"-"` flag to ignore it,
//     or you don't need to set the redis tag.
//     For the type of structure field, we only support simple data types:
//     string, int/uint(8,16,32,64), float(32,64), time.Time(to RFC3339Nano), time.Duration(to Nanoseconds ),
//     if you are other more complex or custom data types, please implement the encoding.BinaryMarshaler interface.
//
// Note that it requires Redis v4 for multiple field/value pairs support.
func (operator *RedisOperator) Set(ctx context.Context, key string, values ...interface{}) error {
	err := operator.rdb.HSet(ctx, key, values).Err()
	return err
}

// 调用HGet方法，根据key和field字段，查询field字段的值
func (operator *RedisOperator) Get(ctx context.Context, key, field string) (string, error) {
	data, err := operator.rdb.HGet(ctx, key, field).Result()
	switch {
	case err != nil:
		return "", err
	default:
		return data, nil
	}
}

// HGetAll - 根据key查询所有字段和值
func (operator *RedisOperator) GetAll(ctx context.Context, key string) (map[string]string, error) {
	data, err := operator.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// HKeys-根据key返回所有字段名
func (operator *RedisOperator) GetKeys(ctx context.Context, key string) ([]string, error) {
	data, err := operator.rdb.HKeys(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// HLen-根据key，查询hash的字段数量
func (operator *RedisOperator) GetHashLen(ctx context.Context, key string) (int64, error) {
	len, err := operator.rdb.HLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return len, nil
}

// HMGet-根据key和多个字段名，批量查询多个hash字段值
func (operator *RedisOperator) MultipleGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	data, err := operator.rdb.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// HMSet-根据key和多个字段名和字段值，批量设置hash字段值
func (operator *RedisOperator) MultipleSet(ctx context.Context, key string, values map[string]interface{}) error {
	err := operator.rdb.HMSet(ctx, key, values).Err()
	return err
}

// HDel-根据key和字段名，删除hash字段，支持批量删除hash字段
func (operator *RedisOperator) Delete(ctx context.Context, key string, fields ...string) error {
	err := operator.rdb.HDel(ctx, key, fields...).Err()
	return err

}

// HExists-检测hash字段名是否存在。
func (operator *RedisOperator) Exists(ctx context.Context, key, field string) (bool, error) {
	res, err := operator.rdb.HExists(ctx, key, field).Result()
	return res, err
}

// 扫描当前redis数据库中包含的主键数量
func (operator *RedisOperator) Scan(ctx context.Context, condition string) int {
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		//*扫描所有key，每次20条
		keys, cursor, err = operator.rdb.Scan(ctx, cursor, condition, 20).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		logrus.Printf("\nfound %d keys\n", n)
		// var value []string
		// for _, key := range keys {
		// 	value, err = operator.rdb.HKeys(ctx, key).Result()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Printf("%v %v\n", key, value)
		// }
		if cursor == 0 {
			break
		}
	}
	return n
}
