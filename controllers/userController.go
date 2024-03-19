package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"net/mail"
	"time"

	"github.com/gin-gonic/gin"
)

func Tes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

// register user
func RegisterUser(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	var user models.User

	// bind data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi email tidak kosong
	if user.Email == "" {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak boleh kosong",
		})
		return
	}

	// validasi email unique
	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email sudah terdaftar",
		})
		return
	}

	// validasi email valid
	if _, err := mail.ParseAddress(user.Email); err != nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak valid",
		})
		return
	}

	// validasi username tidak kosong
	if user.Username == "" {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username tidak boleh kosong",
		})
		return
	}

	// validasi username unique
	if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username sudah terdaftar",
		})
		return
	}

	// validasi password tidak kosong
	if user.Password == "" {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Password tidak boleh kosong",
		})
		return
	}

	// validasi password minimal 6 karakter
	if len(user.Password) < 6 {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Password minimal 6 karakter",
		})
		return
	}

	// validasi age tidak kosong
	if user.Age == 0 {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia tidak boleh kosong",
		})
		return
	}

	// validasi age minimal 8
	if user.Age <= 8 {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia harus diatas 8 tahun",
		})
		return
	}

	// enkripsi password
	

	// set datetime
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// insert data
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(201, gin.H{
		"id": user.Id,
		"email": user.Email,
		"username": user.Username,
		"age": user.Age,
		"profile_image_url": user.ProfileImageUrl,
	})
}


// login user
func LoginUser(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	var user models.User

	// bind data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	passwordInput := user.Password

	// verifikasi emaiil terdaftar
	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "401",
			"message": "Email tidak terdaftar",
		})
		return
	}

	// verifikasi kata sandi
	if pass := helpers.CheckPassHash(passwordInput, user.Password); !pass {
		ctx.JSON(200, gin.H{
			"status":  "error",
			"code": "401",
			"message": "Kata sandi salah",
		})
		return
	}

	// generate token
	token := helpers.GenerateToken(user.Id, user.Email)

	// response
	ctx.JSON(200, gin.H{
		"token": token,
	})
}