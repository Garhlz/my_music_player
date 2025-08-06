package config

import "github.com/spf13/viper"

// Config 存储所有应用配置
type Config struct {
	// 添加数据库相关的具体字段
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	
	// 其他配置
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
}

// LoadConfig 从 .env 文件中读取配置
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env") // viper 会自动寻找 .env 文件
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}