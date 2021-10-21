package global

import (
	"strconv"
	"time"
	"todoapp/models"

	"github.com/go-redis/redis/v7"
)

var Redisclient *redis.Client

func CreateRedisConnection() {
	dsn := EnviromentVariable("REDIS_DSN")
	Redisclient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := Redisclient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func CreateAuth(userid int, td *models.TokenDetails) error {
	atexpire := time.Unix(td.AtExpires, 0)
	rtexpire := time.Unix(td.RtExpires, 0)
	now := time.Now()

	err := Redisclient.Set(td.AccessUuid, strconv.Itoa(userid), atexpire.Sub(now)).Err()
	if err != nil {
		return err
	}

	err = Redisclient.Set(td.RefreshUuid, strconv.Itoa(userid), rtexpire.Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}

func CheckAuth(accessDetails *models.ExtractedTokenData) (int, error) {
	userIdstring, err := Redisclient.Get(accessDetails.Uuid).Result()
	if err != nil {
		return 0, err
	}

	userId, _ := strconv.ParseInt(userIdstring, 10, 64)

	return int(userId), nil
}

func CheckRefreshToken(refreshTokenUuid string) error {
	_, err := Redisclient.Get(refreshTokenUuid).Result()
	if err != nil {
		return err
	}

	return nil
}

func DeleteAuth(AccessUuid, RefreshUuid string) (int64, error) {
	_, err := Redisclient.Del(AccessUuid).Result()
	if err != nil {
		return 0, err
	}

	_, err = Redisclient.Del(RefreshUuid).Result()
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func DeleteRefreshToken(RefreshUuid string) (int64, error) {
	_, err := Redisclient.Del(RefreshUuid).Result()
	if err != nil {
		return 0, err
	}
	return 0, nil
}
