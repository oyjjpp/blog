package store

import (
	"api/internal/pkg/store/cache"
	"api/internal/pkg/store/redis"

	"api/internal/pkg/store/mysql"
)

func Init() error {
	var err error
	if err = mysql.Init(); err != nil {
		return err
	}
	if err = cache.Init(); err != nil {
		return err
	}
	if err = redis.Init(); err != nil {
		return err
	}
	return nil
}
