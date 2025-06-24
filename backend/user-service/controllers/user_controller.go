package controllers

import (
	"net/http"
	"strconv"
	"user-service/database"
	"user-service/models"
	"user-service/session"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendaftarkan pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Registrasi berhasil",
		"user_id":    user.ID,
		"user_email": user.Email,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format login tidak valid"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	session.SetLoggedIn(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"message":    "Login berhasil",
		"user_id":    user.ID,
		"user_email": user.Email,
	})
}

func Me(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	if !session.IsLoggedIn(user.ID) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Belum login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Kamu sedang login",
		"user_id":    user.ID,
		"user_email": user.Email,
	})
}
