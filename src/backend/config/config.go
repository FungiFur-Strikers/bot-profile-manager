package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI    string
	MongoDBName   string
	ServerAddress string
}

func Load() (*Config, error) {
	// .env ファイルが存在する場合にのみ読み込む
	_ = godotenv.Load()

	config := &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"), // デフォルト値を設定
		MongoDBURI:    getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBName:   getEnv("MONGODB_NAME", "botprofiles"),
	}

	return config, nil
}

// getEnv は環境変数の値を取得し、設定されていない場合はデフォルト値を返す
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
