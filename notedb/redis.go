package notedb

import (
	"log"
	"math"

	redis "gopkg.in/redis.v5"
)

type RedisDB struct {
	master         string
	masterPassword string
	slave          string
	slavePassword  string
	appVersion     string
}

const (
	redisKey string = "note"
	okString string = "ok"
)

var _ NoteDB = RedisDB{}

func NewRedisDB(config map[string]string, appVersion string) RedisDB {
	if _, exists := config["master"]; !exists {
		config["master"] = "redis-master:6379"
	}

	if _, exists := config["masterPassword"]; !exists {
		config["masterPassword"] = ""
	}

	if _, exists := config["slave"]; !exists {
		config["slave"] = "redis-slave:6379"
	}

	if _, exists := config["slavePassword"]; !exists {
		config["slavePassword"] = ""
	}

	return RedisDB{
		master:         config["master"],
		masterPassword: config["masterPassword"],
		slave:          config["slave"],
		slavePassword:  config["slavePassword"],
		appVersion:     appVersion,
	}
}

func createRedisClient(addr, password string) *(redis.Client) {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0, // use default DB
	})
}

func (redisDB RedisDB) GetAllNotes() ([]string, error) {
	cmd := createRedisClient(redisDB.slave, redisDB.slavePassword).LRange(redisKey, 0, math.MaxInt64)

	// Fallback to read from master
	if cmd.Err() != nil {
		log.Println("Fallback using Redis Master")
		cmd = createRedisClient(redisDB.master, redisDB.masterPassword).LRange(redisKey, 0, math.MaxInt64)
	}
	return cmd.Val(), cmd.Err()
}

func (redisDB RedisDB) SaveNote(note string) error {
	return createRedisClient(redisDB.master, redisDB.masterPassword).RPush(redisKey, note).Err()
}

func (redisDB RedisDB) DeleteNote(note string) error {
	return createRedisClient(redisDB.master, redisDB.masterPassword).LRem(redisKey, 1, note).Err()
}
