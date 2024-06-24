package configs

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	lock    = &sync.Mutex{}
	configs *Configs
)

func GetInstance() *Configs {
	lock.Lock()
	defer lock.Unlock()

	configs = &Configs{
		App: AppConfig{
			Name:          getEnv("APP_NAME", "go-skeleton"),
			Env:           getEnv("APP_ENV", "dev"),
			Port:          getEnv("APP_PORT", "8000"),
			BasicUsername: getEnv("BASIC_USERNAME", "user1"),
			BasicPassword: getEnv("BASIC_PASSWORD", "Secret123"),
		},
		DB: DbConfig{
			Host:        getEnv("DB_HOST", "127.0.0.1"),
			Port:        getEnv("DB_PORT", "5432"),
			DbName:      getEnv("DB_NAME", "my_db"),
			Username:    getEnv("DB_USER", "postgres"),
			Password:    getEnv("DB_PASS", ""),
			DbIsMigrate: getEnv("DB_IS_MIGRATE", "true") == "true",
		},
		Redis: RedisConfig{
			Host:                getEnv("REDIS_HOST", "127.0.0.1"),
			Port:                GetRedisPort(),
			Password:            getEnv("REDIS_PASSWORD", ""),
			Database:            GetRedisDatabase(),
			Master:              getEnv("REDIS_MASTER", "master"),
			PoolSize:            GetRedisPoolSize(),
			PoolTimeout:         GetRedisPoolTimeout(),
			MinIdleConn:         GetRedisMinIdleConn(),
			DefaultCacheTimeOut: GetRedisDefaultCacheTimeOut(),
		},
	}

	return configs
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(value)
	}
	return strings.TrimSpace(fallback)
}

func BackEndUrl() string {
	backEndUrl := ``
	switch os.Getenv(`ENVIRONMENT`) {
	case `dev`:
		backEndUrl = os.Getenv(`BACK_END_DEV_URL`)
	case `stg`:
		backEndUrl = os.Getenv(`BACK_END_STG_URL`)
	case `prd`:
		backEndUrl = os.Getenv(`BACK_END_PRD_URL`)
	}
	return backEndUrl
}

func ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func GetRedisPort() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_HOST", "6379"))
	if err != nil {
		return 6379
	}

	return intVar
}
func GetRedisDatabase() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_DATABASE", "0"))
	if err != nil {
		return 0
	}

	return intVar
}

func GetRedisPoolSize() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_POOL_SIZE", "128"))
	if err != nil {
		return 128
	}

	return intVar
}

func GetRedisPoolTimeout() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_POOL_TIMEOUT", "10"))
	if err != nil {
		return 10
	}

	return intVar
}

func GetRedisMinIdleConn() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_MIN_IDLE_CONN", "4"))
	if err != nil {
		return 4
	}

	return intVar
}
func GetRedisDefaultCacheTimeOut() time.Duration {
	intVar, err := strconv.Atoi(getEnv("CACHE_DEFAULT_TIMEOUT", "5"))
	if err != nil {
		return 5 * time.Minute
	}

	return time.Duration(intVar) * time.Minute
}
