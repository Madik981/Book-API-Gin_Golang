package handlers

import (
	"net/http"
	"time"

	"Book-API-Gin_Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const tokenTTL = 24 * time.Hour

type AuthHandler struct {
	store     *models.Store
	jwtSecret []byte
}

func NewAuthHandler(store *models.Store, jwtSecret []byte) *AuthHandler {
	return &AuthHandler{store: store, jwtSecret: jwtSecret}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input models.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, created := h.store.CreateUser(input)
	if !created {
		c.JSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}

	c.JSON(http.StatusCreated, models.BuildUserResponse(user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := h.store.AuthenticateUser(input)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := h.buildToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"token_type": "Bearer",
		"expires_in": int(tokenTTL.Seconds()),
	})
}

func (h *AuthHandler) buildToken(user models.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		"iat":      now.Unix(),
		"exp":      now.Add(tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.jwtSecret)
}
