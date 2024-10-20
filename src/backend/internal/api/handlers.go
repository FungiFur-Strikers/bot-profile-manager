package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bot-profile-manager/internal/models"
	"bot-profile-manager/internal/service"
)

type Server struct {
	profileService *service.ProfileService
}

func NewServer(profileService *service.ProfileService) *Server {
	return &Server{profileService: profileService}
}

func (s *Server) GetBotBotIdProfile(c *gin.Context, botId string) {
	profile, err := s.profileService.GetProfile(c.Request.Context(), botId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (s *Server) PutBotBotIdProfile(c *gin.Context, botId string) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile.BotID = botId

	err := s.profileService.UpsertProfile(c.Request.Context(), &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}
