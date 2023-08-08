package routesHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryougi-shiky/COMP90018_Backend/models"
	"github.com/ryougi-shiky/COMP90018_Backend/services"
	"net/http"
)

// memo used when create, update, delete
type Memo struct {
	MemoID  string `json:"memoId"`
	UserID  string `json:"userId" binding:"required"`
	Title   string `json:"title" biding:"required"`
	Content string `json:"content" binding:"required"`
}

// userId, when reading all memos of this user
type GetMemoRequest struct {
	UserID string `json:"userId" binding:"required"`
}

func CreateMemo(memoService services.MemoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get json from client
		var newMemo Memo
		if err := context.ShouldBindJSON(&newMemo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse userid string to uuid
		UseridUuid, err := uuid.Parse(newMemo.UserID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// match json to memo object
		memo := models.Memo{
			MemoID:  uuid.New(),
			UserID:  UseridUuid,
			Title:   newMemo.Title,
			Content: newMemo.Content,
		}

		// Creating new memo
		if err := memoService.CreateMemo(&memo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "New memo created!"})
	}
}

func UpdateMemo(memoService services.MemoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var updateMemo Memo
		if err := context.ShouldBindJSON(&updateMemo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse userid string to uuid
		UseridUuid, err := uuid.Parse(updateMemo.UserID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		MemoidUuid, err := uuid.Parse(updateMemo.MemoID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// match json to memo object
		memo := models.Memo{
			MemoID:  MemoidUuid,
			UserID:  UseridUuid,
			Title:   updateMemo.Title,
			Content: updateMemo.Content,
		}

		// Updating the memo
		if err := memoService.UpdateMemo(&memo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "The memo updated!"})
	}
}

func DeleteMemo(memoService services.MemoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var delMemo Memo
		if err := context.ShouldBindJSON(&delMemo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse userid string to uuid
		UseridUuid, err := uuid.Parse(delMemo.UserID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		MemoidUuid, err := uuid.Parse(delMemo.MemoID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// match json to memo object
		memo := models.Memo{
			MemoID:  MemoidUuid,
			UserID:  UseridUuid,
			Title:   delMemo.Title,
			Content: delMemo.Content,
		}

		// Updating the memo
		if err := memoService.DeleteMemo(&memo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "The memo deleted!"})
	}
}

func GetUserMemo(memoService services.MemoService) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user GetMemoRequest
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse userid string to uuid
		UseridUuid, err := uuid.Parse(user.UserID)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Retrieving the memo
		memos, err := memoService.GetUserMemo(&UseridUuid)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// return a list of memos
		context.JSON(http.StatusOK, gin.H{"memos": memos})
	}
}
