package controllers

import (
	"final-project/database"
	"final-project/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var photo models.Photo
	var user models.User
	var userID = uint(ctx.MustGet("userData").(float64))

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi title tidak boleh kosong
	if photo.Title == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Title tidak boleh kosong",
		})
		return
	}

	// validasi photo url tidak boleh kosong
	if photo.PhotoUrl == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Photo Url tidak boleh kosong",
		})
		return
	}

	// validasi photo url valid
	if !strings.Contains(photo.PhotoUrl, ".") {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Photo url tidak valid",
		})
		return
	}

	// set datetime
	photo.CreatedAt = time.Now()
	photo.UpdatedAt = time.Now()

	// set user id
	photo.UserId = userID

	// insert data
	if err := db.Create(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response 
	ctx.JSON(201, photo)
}

func GetAllUsersPhotos(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var user models.User
	var photos []models.Photo
	var response []map[string]interface{} = []map[string]interface{}{}

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get all photos
	if err := db.Find(&photos).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// buat response 
	for _, photo := range photos {
		// get user 
		var photoUser models.User

        // Mengambil informasi pengguna berdasarkan ID pengguna foto
        if err := db.First(&photoUser, photo.UserId).Error; err != nil {
            ctx.JSON(400, gin.H{
                "status":  "error",
                "code":    "400",
                "message": err.Error(),
            })
            return
        }

		response = append(response, map[string]interface{}{
			"id":        photo.Id,
			"title":     photo.Title,
			"caption":   photo.Caption,
			"photo_url": photo.PhotoUrl,
			"user_id":   photo.UserId,
			"user": map[string]interface{}{
				"id":       photoUser.Id,
				"email":    photoUser.Email,
				"username": photoUser.Username,
			},
		})
	}

	// response
	ctx.JSON(200, response)
}

func GetUsersPhoto(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var photoId = ctx.Param("photoId")
	var user models.User
	var photoUser models.User
	var photo models.Photo

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get photo by id
	if err := db.Where("id = ?", photoId).First(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get user photo
	if err := db.Where("id = ?", photo.UserId).First(&photoUser).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response 
	ctx.JSON(200, gin.H{
		"id":        photo.Id,
		"title":     photo.Title,
		"caption":   photo.Caption,
		"photo_url": photo.PhotoUrl,
		"user_id":   photo.UserId,
		"user": map[string]interface{}{
			"id":       photoUser.Id,
			"email":    photoUser.Email,
			"username": photoUser.Username,
		},
	})
}

func UpdatePhoto(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var photoId = ctx.Param("photoId")
	var photo models.Photo
	var user models.User

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// cek photo ada
	if err := db.Where("id = ? AND user_id = ?", photoId, userID).First(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi title tidak boleh kosong
	if photo.Title == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Title tidak boleh kosong",
		})
		return
	}

	// validasi photo url tidak boleh kosong
	if photo.PhotoUrl == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Photo Url tidak boleh kosong",
		})
		return
	}

	// validasi photo url valid
	if !strings.Contains(photo.PhotoUrl, ".") {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Photo url tidak valid",
		})
		return
	}

	// set datetime
	photo.UpdatedAt = time.Now()

	// simpan perubahan
	if err := db.Save(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, photo)
}

func DeletePhoto(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var photoId = ctx.Param("photoId")
	var user models.User
	var photo models.Photo

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// cek photo ada
	if err := db.Where("id = ? AND user_id = ?", photoId, userID).First(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// hapus comments
	if err := db.Delete(&photo.Comments, "photo_id = ?", photo.Id).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// hapus photo
	if err := db.Delete(&photo).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response 
	ctx.JSON(200, gin.H{
		"status":  "success",
		"code":    "200",
		"message": "Photo berhasil dihapus",
	})
}