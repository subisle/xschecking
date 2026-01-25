package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Redis     RedisConfig     `mapstructure:"redis"`
	JWT       JWTConfig       `mapstructure:"jwt"`
	Captcha   CaptchaConfig   `mapstructure:"captcha"`
	Email     EmailConfig     `mapstructure:"email"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
	Risk      RiskConfig      `mapstructure:"risk"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"dbname"`
	Charset      string `mapstructure:"charset"`
	ParseTime    bool   `mapstructure:"parse_time"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
	Issuer      string `mapstructure:"issuer"`
}

type CaptchaConfig struct {
	LuckyColaAPIURL    string `mapstructure:"luckycola_api_url"`
	LuckyColaAppID     string `mapstructure:"luckycola_app_id"`
	LuckyColaAppSecret string `mapstructure:"luckycola_app_secret"`
	ExpireMinutes      int    `mapstructure:"expire_minutes"`
}

type EmailConfig struct {
	SMTPHost          string `mapstructure:"smtp_host"`
	SMTPPort          int    `mapstructure:"smtp_port"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	FromName          string `mapstructure:"from_name"`
	CodeExpireMinutes int    `mapstructure:"code_expire_minutes"`
	CodeLength        int    `mapstructure:"code_length"`
}

type RateLimitConfig struct {
	LoginMaxAttempts       int `mapstructure:"login_max_attempts"`
	LoginBlockMinutes      int `mapstructure:"login_block_minutes"`
	RegisterMaxPerIPPerDay int `mapstructure:"register_max_per_ip_per_day"`
	EmailCodeMaxPerHour    int `mapstructure:"email_code_max_per_hour"`
}

type RiskConfig struct {
	Enable               bool `mapstructure:"enable"`
	MaxLoginFailures     int  `mapstructure:"max_login_failures"`
	AccountFreezeMinutes int  `mapstructure:"account_freeze_minutes"`
}

var GlobalConfig *Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	GlobalConfig = &Config{}
	if err := viper.Unmarshal(GlobalConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	return nil
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	return GlobalConfig
}
