package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configsFolderPathEnv = "RLMP_CONFIGS_FOLDER_PATH"
	configNameEnv        = "RLMP_CONFIG_NAME"
	dotEnvFilePathEnv    = "RLMP_DOTENV_PATH"
)

func Get() *Config {
	var (
		once sync.Once
		cfg  *Config
	)

	once.Do(func() {
		if err := initConfigParser(); err != nil {
			logrus.Fatal(err)
		}

		if err := initDotEnvParser(); err != nil {
			logrus.Fatal(err)
		}

		cfg = newConfig()
	})

	return cfg
}

func initConfigParser() error {
	configsFolderPath := os.Getenv(configsFolderPathEnv)
	if configsFolderPath == "" {
		return fmt.Errorf("empty configs folder path environment variable: %s", configsFolderPathEnv)
	}

	configName := os.Getenv(configNameEnv)
	if configName == "" {
		return fmt.Errorf("empty config name environment variable: %s", configNameEnv)
	}

	viper.AddConfigPath(configsFolderPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while reading config: %w", err)
	}

	return nil
}

func initDotEnvParser() error {
	dotEnvFilePath := os.Getenv(dotEnvFilePathEnv)
	if dotEnvFilePath == "" {
		return fmt.Errorf("empty .env file path environment variable: %s", dotEnvFilePath)
	}

	if err := godotenv.Load(dotEnvFilePath); err != nil {
		return fmt.Errorf("error loading env variables from [%s]: %w", dotEnvFilePath, err)
	}

	return nil
}

type Config struct {
	Server        *Server
	DB            *DB
	Cache         *Cache
	Authorization *Authorization
	Parsing       *Parsing
}

func newConfig() *Config {
	return &Config{
		Server:        newServer(),
		DB:            newDB(),
		Cache:         newCache(),
		Authorization: newAuth(),
		Parsing:       newParsing(),
	}
}

type Server struct {
	Port           string
	MaxHeaderBytes int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

func newServer() *Server {
	return &Server{
		Port:           viper.GetString("server.port"),
		MaxHeaderBytes: viper.GetInt("server.max_header_bytes"),
		ReadTimeout:    viper.GetDuration("server.read_timeout") * time.Second,
		WriteTimeout:   viper.GetDuration("server.write_timeout") * time.Second,
	}
}

type DB struct {
	Driver         string
	Host           string
	Port           string
	Username       string
	Password       string
	DBName         string
	SSLMode        string
	MigrationsPath string
}

func newDB() *DB {
	return &DB{
		Driver:         viper.GetString("db.driver"),
		Host:           viper.GetString("db.host"),
		Port:           viper.GetString("db.port"),
		Username:       viper.GetString("db.username"),
		Password:       os.Getenv("DB_PASSWORD"),
		DBName:         viper.GetString("db.dbname"),
		SSLMode:        viper.GetString("db.sslmode"),
		MigrationsPath: viper.GetString("db.migrations_path"),
	}
}

type Cache struct {
	Address  string
	Password string
	DB       int
}

func newCache() *Cache {
	return &Cache{
		Address:  viper.GetString("cache.address"),
		Password: viper.GetString("cache.password"),
		DB:       viper.GetInt("cache.db"),
	}
}

type JWTToken struct {
	TTL        time.Duration
	SigningKey []byte
}

type Authorization struct {
	AccessToken  JWTToken
	RefreshToken JWTToken
}

func newAuth() *Authorization {
	return &Authorization{
		AccessToken: JWTToken{
			TTL:        time.Minute * viper.GetDuration("auth.access_token.ttl"),
			SigningKey: []byte(viper.GetString("auth.access_token.signing_key")),
		},
		RefreshToken: JWTToken{
			TTL:        time.Minute * viper.GetDuration("auth.refresh_token.ttl"),
			SigningKey: []byte(viper.GetString("auth.refresh_token.signing_key")),
		},
	}
}

type Parsing struct {
	RatingListTTL       time.Duration
	ReadTimeout         time.Duration
	ReadBufferSize      int
	MaxResponseBodySize int
	MaxConnsPerHost     int
}

func newParsing() *Parsing {
	return &Parsing{
		RatingListTTL:       time.Minute * viper.GetDuration("parsing.rating_list_ttl"),
		ReadTimeout:         viper.GetDuration("parsing.read_timeout") * time.Second,
		ReadBufferSize:      viper.GetInt("parsing.read_buffer_size"),
		MaxResponseBodySize: viper.GetInt("parsing.max_response_body_size"),
		MaxConnsPerHost:     viper.GetInt("parsing.max_conns_per_host"),
	}
}
