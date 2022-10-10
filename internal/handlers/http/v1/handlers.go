package handlers

import (
	"auth-service/config"
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

type Handler struct {
	service *service.Service
	logger  *zap.SugaredLogger
	cfg     *config.Configs
}

func NewHandler(services *service.Service, logger *zap.SugaredLogger, cfg *config.Configs) *Handler {
	return &Handler{
		service: services,
		logger:  logger,
		cfg:     cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)
	router.POST("/sign-out", h.SignOut)

	return router
}

func (h *Handler) SignUp(c *gin.Context) {
	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.logger.Errorf("Error occured while reading request body: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}

	var signUpUserRequest *models.SignUpUser
	err = json.Unmarshal(requestBody, &signUpUserRequest)
	if err != nil {
		h.logger.Errorf("Error occurred while unmarshalling request body: %s", err.Error())
		c.JSON(401, gin.H{
			"error": models.ErrInvalidInput.Error(),
		})
		return
	}

	err = h.service.CreateUser(signUpUserRequest)
	if err != nil {
		h.logger.Errorf("Error occurred while creating new user: %s", err.Error())
		c.JSON(500, gin.H{
			"error": models.ErrInternalServerError,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
