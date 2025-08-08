// my-music-go/internal/config/config.go
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// 告诉 Viper 自动读取环境变量（保持，但我们会通过 BindEnv 增强）
	viper.AutomaticEnv()

	// === 新增：显式绑定环境变量 ===
	// 即使没有 .env 文件，也能确保这些环境变量被读取并绑定到 Viper
	// 绑定到 Viper 内部的配置树中，这样 Unmarshal 才能找到它们
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("SERVER_ADDRESS")
	viper.BindEnv("JWT_SECRET")
	// === 结束新增 ===

	fmt.Println("--- Viper 调试信息 ---")
	fmt.Printf("Viper 配置路径: %s, 已使用的配置文件: %s\n", path, viper.ConfigFileUsed())
	fmt.Printf("从 OS 环境读取 DB_HOST: %s\n", os.Getenv("DB_HOST"))
	fmt.Printf("从 OS 环境读取 DB_USER: %s\n", os.Getenv("DB_USER"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(".env 文件未找到，将继续使用环境变量。")
		} else {
			fmt.Printf("读取 .env 配置文件时出错: %v\n", err)
			return
		}
	} else {
		fmt.Println(".env 文件已找到并成功读取。")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("解析配置到结构体时出错: %v\n", err)
		return
	}

	fmt.Println("配置加载成功！")
	fmt.Printf("DB Host: %s, DB User: %s\n", config.DBHost, config.DBUser)
	fmt.Printf("DB Password: %s, DB Name: %s, DB Port: %s\n", config.DBPassword, config.DBName, config.DBPort)
	fmt.Println("--- Viper 调试结束 ---")

	return config, nil
}
