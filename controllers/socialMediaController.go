package controllers

import (
	"final-project/database"
	"final-project/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(ctx *gin.Context){
	// inisialisasi variabel
	var db = database.GetDB()
	userID := ctx.MustGet("userData").(float64)
	var socialMedia models.SocialMedia
	var user models.User

	// get user by id
	if err := db.First(&user, uint(userID)).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": err.Error(),
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi name tidak boleh kosong
	if socialMedia.Name == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Name tidak boleh kosong",
		})
		return
	}

	// validasi social_media_url tidak boleh kosong
	if socialMedia.SocialMediaUrl == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Social Media Url tidak boleh kosong",
		})
		return
	}

	// validasi social_media_url valid
	if !strings.Contains(socialMedia.SocialMediaUrl, ".") {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Social media url tidak valid",
		})
		return
	}

	// set datetime
	socialMedia.CreatedAt = time.Now()
	socialMedia.UpdatedAt = time.Now()

	// set user id
	socialMedia.UserId = uint(userID)

	// insert data
	if err := db.Create(&socialMedia).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(201, socialMedia)
}

func GetAllSocialMedias(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64)) 
	var socialMedias []models.SocialMedia
	var user models.User
	var response []map[string]interface{} = []map[string]interface{}{}

	// get user by id
	if err := db.First(&user, userID).Error; err != nil { 
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get all social medias
	if err := db.Where("user_id = ?", userID).Find(&socialMedias).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// masukkan informasi user kedalam setiap social media
	for _, socialMedia := range socialMedias {
		response = append(response, map[string]interface{}{ 
			"id":              socialMedia.Id,
			"name":            socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaUrl,
			"user_id":         socialMedia.UserId,
			"user": map[string]interface{}{ 
				"id":       user.Id,
				"email":    user.Email,
				"username": user.Username,
			},
		})
	}

	// response
	ctx.JSON(200, response)
}

func GetSocialMedia(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var socialMediaId = ctx.Param("socialMediaId")
	var socialMedia models.SocialMedia
	var user models.User

	// get user by id
	if err := db.First(&user, userID).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get social media by social media id dan user id
	if err := db.Where("id = ? AND user_id = ?", socialMediaId, userID).First(&socialMedia).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, gin.H{
		"id":              socialMedia.Id,
		"name":            socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":         socialMedia.UserId,
		"user": map[string]interface{}{ 
			"id":       user.Id,
			"email":    user.Email,
			"username": user.Username,
		},
	})
}

func UpdateSocialMedia(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var socialMediaId = ctx.Param("socialMediaId")
	var socialMedia models.SocialMedia
	var user models.User	

	// get user by id
	if err := db.First(&user, userID).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get social media by social media id dan user id
	if err := db.Where("id = ? AND user_id = ?", socialMediaId, userID).First(&socialMedia).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// validasi name tidak boleh kosong
	if socialMedia.Name == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Name tidak boleh kosong",
		})
		return
	}

	// validasi social_media_url tidak boleh kosong
	if socialMedia.SocialMediaUrl == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Social Media Url tidak boleh kosong",
		})
		return
	}

	// validasi social_media_url valid
	if !strings.Contains(socialMedia.SocialMediaUrl, ".") {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code": "400",
			"message": "Social media url tidak valid",
		})
		return
	}

	// set datetime
	socialMedia.UpdatedAt = time.Now()

	// update social media
	if err := db.Model(&socialMedia).Updates(&socialMedia).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(200, socialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var socialMediaId = ctx.Param("socialMediaId")
	var socialMedia models.SocialMedia
	var user models.User

	// get user by id
	if err := db.First(&user, userID).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get social media by social media id dan user id
	if err := db.Where("id = ? AND user_id = ?", socialMediaId, userID).First(&socialMedia).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// hapus social media
	if err := db.Delete(&socialMedia).Error; err != nil {
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
		"message": "Social media berhasil dihapus",
	})
}
