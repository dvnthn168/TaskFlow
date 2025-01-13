package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DBSource       string `mapstructure:"DB_SOURCE"`
	HTTPServerAddr string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddr string `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisAddr      string `mapstructure:"REDIS_ADDR"`
	RedisPassword  string `mapstructure:"REDIS_PASSWORD"`

	AccountEmail      string `mapstructure:"ACCOUNT_EMAIL"`
	PasswordEmail     string `mapstructure:"PASSWORD_EMAIL"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`

	KeyVerifyUser  string `mapstructure:"KEY_VERIFY_USER"`
	KeyRecoverUser string `mapstructure:"KEY_RECOVER_USER"`

	GateWayServerAddr string `mapstructure:"GATEWAY_SERVER_ADDRESS"`

	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	TokenSymmetricKeyCustomer string        `mapstructure:"TOKEN_SYMMETRIC_KEY_CUSTOMER"`
	AesKey                    string        `mapstructure:"AES_KEY"`
	CloudinaryUrl             string        `mapstructure:"CLOUDINARY_URL"`
	CloudinaryName            string        `mapstructure:"CLOUDINARY_NAME"`
	CloudinarySecret          string        `mapstructure:"CLOUDINARY_SECRET_KEY"`
	CloudinaryKey             string        `mapstructure:"CLOUDINARY_KEY"`
}

func LoadConfig(path string) (_ *Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, err
}
