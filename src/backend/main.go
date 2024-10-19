package main

import (
	"net/http"

	"bot-profile-manager/api"

	"github.com/gin-gonic/gin"
)

// ServerInterfaceを実装する構造体
type Server struct {
	// 必要なフィールドがあればここに追加
}

// GetBotBotIdProfileの実装
func (s *Server) GetBotBotIdProfile(c *gin.Context, botId string) {
	// ここにプロフィール取得のロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "プロフィールを取得しました", "botId": botId})
}

// PutBotBotIdProfileの実装
func (s *Server) PutBotBotIdProfile(c *gin.Context, botId string) {
	// ここにプロフィール更新のロジックを実装
	c.JSON(http.StatusOK, gin.H{"message": "プロフィールを更新しました", "botId": botId})
}

func main() {
	r := gin.Default()
	server := &Server{}
	api.RegisterHandlers(r, server)
	r.Run() // listen and serve on 0.0.0.0:8080
}
