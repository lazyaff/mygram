package controllers

import (
	"final-project/database"
	"final-project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var comment models.Comment
	var user models.User
	var photo models.Photo

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code":    "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Periksa kembali inputan anda",
		})
		return
	}

	// cek foto ada
	if err := db.Where("id = ?", comment.PhotoId).First(&photo).Error; err != nil {
		ctx.JSON(404, gin.H{
			"status":  "error",
			"code":    "404",
			"message": "Foto tidak ditemukan",
		})
		return
	}

	// validasi message tidak kosong
	if comment.Message == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Message tidak boleh kosong",
		})
		return
	}

	// set datetime
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	// set user
	comment.UserId = userID

	// insert data
	if err := db.Create(&comment).Error; err != nil {
		ctx.JSON(500, gin.H{
			"status":  "error",
			"code":    "500",
			"message": err.Error(),
		})
		return
	}

	// response
	ctx.JSON(201, comment)
}

func GetAllUsersComments(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var comments []models.Comment
	var user models.User
	var response []map[string]interface{} = []map[string]interface{}{}

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code":    "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// get all comments
	if err := db.Find(&comments).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// buat response
	for _, comment := range comments {
		var commentUser models.User
		var commentPhoto models.Photo

		// get user
		if err := db.Where("id = ?", comment.UserId).First(&commentUser).Error; err != nil {
			ctx.JSON(400, gin.H{
				"status":  "error",
				"code":    "400",
				"message": err.Error(),
			})
			return
		}

		// get photo
		if err := db.Where("id = ?", comment.PhotoId).First(&commentPhoto).Error; err != nil {
			ctx.JSON(400, gin.H{
				"status":  "error",
				"code":    "400",
				"message": err.Error(),
			})
			return
		}

		response = append(response, map[string]interface{}{
			"id": comment.Id,
			"message": comment.Message,
			"user_id": comment.UserId,
			"photo_id": comment.PhotoId,
			"user": map[string]interface{}{
				"id": commentUser.Id,
				"email": commentUser.Email,
				"username": commentUser.Username,
			},
			"photo": map[string]interface{}{
				"id": commentPhoto.Id,
				"title": commentPhoto.Title,
				"caption": commentPhoto.Caption,
				"photo_url": commentPhoto.PhotoUrl,
				"user_id": commentPhoto.UserId,
			},
		})
	}

	// response
	ctx.JSON(200, response)
}

func GetUserComment(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var commentId = ctx.Param("commentId")
	var comment models.Comment
	var user models.User
	var commentUser models.User
	var commentPhoto models.Photo

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code":    "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// get comment by id
	if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		ctx.JSON(404, gin.H{
			"status":  "error",
			"code":    "404",
			"message": "Komentar tidak ditemukan",
		})
		return
	}

	// get user comment
	if err := db.Where("id = ?", comment.UserId).First(&commentUser).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// get photo comment
	if err := db.Where("id = ?", comment.PhotoId).First(&commentPhoto).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response 
	ctx.JSON(200, gin.H{
		"id": comment.Id,
		"message": comment.Message,
		"user_id": comment.UserId,
		"photo_id": comment.PhotoId,
		"user": map[string]interface{}{
			"id": commentUser.Id,
			"email": commentUser.Email,
			"username": commentUser.Username,
		},
		"photo": map[string]interface{}{
			"id": commentPhoto.Id,
			"title": commentPhoto.Title,
			"caption": commentPhoto.Caption,
			"photo_url": commentPhoto.PhotoUrl,
			"user_id": commentPhoto.UserId,
		},
	})
}

func UpdateComment(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var commentID = ctx.Param("commentId")
	var comment models.Comment
	var user models.User

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code":    "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// cek komentar ada 
	if err := db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
		ctx.JSON(404, gin.H{
			"status":  "error",
			"code":    "404",
			"message": "Komentar tidak ditemukan",
		})
		return
	}

	// bind data
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// validasi message tidak kosong
	if comment.Message == "" {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": "Message tidak boleh kosong",
		})
		return
	}

	// set datetime
	comment.UpdatedAt = time.Now()

	// update comment
	if err := db.Save(&comment).Error; err != nil {
		ctx.JSON(400, gin.H{
			"status":  "error",
			"code":    "400",
			"message": err.Error(),
		})
		return
	}

	// response 
	ctx.JSON(200, comment)
}

func DeleteComment(ctx *gin.Context) {
	// inisialisasi variabel
	var db = database.GetDB()
	var userID = uint(ctx.MustGet("userData").(float64))
	var commentID = ctx.Param("commentId")
	var user models.User
	var comment models.Comment

	// cek user terdaftar
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		ctx.JSON(401, gin.H{
			"status":  "error",
			"code":    "401",
			"message": "Pengguna tidak terdaftar",
		})
		return
	}

	// cek komentar ada
	if err := db.Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error; err != nil {
		ctx.JSON(404, gin.H{
			"status":  "error",
			"code":    "404",
			"message": "Komentar tidak ditemukan",
		})
		return
	}

	// hapus comment
	if err := db.Delete(&comment).Error; err != nil {
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
		"message": "Komentar berhasil di hapus",
	})
}