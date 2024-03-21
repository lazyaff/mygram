package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"net/mail"
	"strings"
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
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi email tidak kosong
	if user.Email == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak boleh kosong",
		})
		return
	}

	// validasi email unique
	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email sudah terdaftar",
		})
		return
	}

	// validasi email valid
	if _, err := mail.ParseAddress(user.Email); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak valid",
		})
		return
	}

	// validasi username tidak kosong
	if user.Username == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username tidak boleh kosong",
		})
		return
	}

	// validasi username unique
	if err := db.Where("username = ?", user.Username).First(&user).Error; err == nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username sudah terdaftar",
		})
		return
	}

	// validasi password tidak kosong
	if user.Password == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Password tidak boleh kosong",
		})
		return
	}

	// validasi password minimal 6 karakter
	if len(user.Password) < 6 {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Password minimal 6 karakter",
		})
		return
	}

	// validasi age tidak kosong
	if user.Age == 0 {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia tidak boleh kosong",
		})
		return
	}

	// validasi age minimal 8
	if user.Age < 8 {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia minimal 8 tahun",
		})
		return
	}

	// validasi profile image url
	if user.ProfileImageUrl != "" {
		// validasi url harus valid
		if !strings.Contains(user.ProfileImageUrl, ".") {
			ctx.JSON(400, gin.H{
				"status":  "error",
				"code": "400",
				"message": "Url tidak valid",
			})
			return
		}
	}

	// set datetime
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// insert data
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
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
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	passwordInput := user.Password

	// verifikasi emaiil terdaftar
	if err := db.Where("email = ?", user.Email).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code": "401",
			"message": "Email tidak terdaftar",
		})
		return
	}

	// verifikasi kata sandi
	if pass := helpers.CheckPassHash(passwordInput, user.Password); !pass {
		ctx.JSON(401, gin.H{
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

// update user information
func UpdateUser(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	userID := ctx.MustGet("userData").(float64)
	var user models.User
	var cekUser models.User
	
	// verifikasi user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code": "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi email tidak kosong
	if user.Email == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak boleh kosong",
		})
		return
	}

	// validasi email unique
	if err := db.Where("email = ? AND id != ?", user.Email, user.Id).First(&cekUser).Error; err == nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email sudah terdaftar",
		})
		return
	}

	// validasi email valid
	if _, err := mail.ParseAddress(user.Email); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Email tidak valid",
		})
		return
	}

	// validasi username tidak kosong
	if user.Username == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username tidak boleh kosong",
		})
		return
	}

	// validasi username unique
	if err := db.Where("username = ? AND id != ?", user.Username, user.Id).First(&cekUser).Error; err == nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Username sudah terdaftar",
		})
		return
	}

	// validasi age tidak kosong
	if user.Age == 0 {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia tidak boleh kosong",
		})
		return
	}

	// validasi age minimal 8
	if user.Age < 8 {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Usia minimal 8 tahun",
		})
		return
	}

	// validasi profile image url
	if user.ProfileImageUrl != "" {
		// validasi url harus valid
		if !strings.Contains(user.ProfileImageUrl, ".") {
			ctx.JSON(400, gin.H{
				"status":  "error",
				"code": "400",
				"message": "Url tidak valid",
			})
			return
		}
	}

	// set datetime
	user.UpdatedAt = time.Now()

	// update user
	if err := db.Save(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Gagal update user",
		})
		return
	}

	// response
	ctx.JSON(200, gin.H{
		"id": user.Id,
		"email": user.Email,
		"username": user.Username,
		"age": user.Age,
		"profile_image_url": user.ProfileImageUrl,
	})
}

// delete user
func DeleteUser(ctx *gin.Context) {
	// inisialisasi variabel
	db := database.GetDB()
	userID := ctx.MustGet("userData").(float64)
	var user models.User

	// verifikasi user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code": "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// hapus foto
	if err := db.Delete(&user.Photos, "user_id = ?", user.Id).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Gagal hapus foto",
		})
		return
	}

	// hapus komentar
	if err := db.Delete(&user.Comments, "user_id = ?", user.Id).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Gagal hapus komentar",
		})
		return
	}

	// hapus media sosial
	if err := db.Delete(&user.SocialMedias, "user_id = ?", user.Id).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Gagal hapus media sosial",
		})
		return
	}

	// hapus user
	if err := db.Delete(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Gagal hapus user",
		})
		return
	}

	// response
	ctx.JSON(200, gin.H{
		"status":  "success",
		"code": "200",
		"message": "Pengguna berhasil dihapus",
	})
}