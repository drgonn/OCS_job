package handlers

import (
	"net/http"
	"word-card-app/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// 获取所有单词卡
func GetWordCards(c *gin.Context) {
	wordCards, err := models.GetAllWordCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get word cards"})
		return
	}
	c.JSON(http.StatusOK, wordCards)
}

// 创建单词卡
func CreateWordCard(c *gin.Context) {
	var wordCard models.WordCard
	if err := c.ShouldBindJSON(&wordCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// 这里要求用的govalidator， 否则直接在tag里面写validator就行了
	back := wordCard.Back
	cleanedBack := govalidator.Trim(back, "")
	if !isValidText(cleanedBack) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.InsertWordCard(&wordCard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create word card"})
		return
	}
	c.JSON(http.StatusCreated, wordCard)
}

func isValidText(text string) bool {
	// 使用正则表达式检查文本是否只包含英文和中文字符
	// 这是一个简化的示例，实际的正则表达式可能需要更复杂
	validPattern := "^[a-zA-Z\u4e00-\u9fa5]+$"
	isValid := govalidator.StringMatches(text, validPattern)
	return isValid
}

// 更新单词卡
func UpdateWordCard(c *gin.Context) {
	id := c.Param("id")
	var wordCard models.WordCard
	if err := c.ShouldBindJSON(&wordCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	err := models.UpdateWordCard(id, &wordCard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update word card"})
		return
	}
	c.JSON(http.StatusOK, wordCard)
}

// 删除单词卡
func DeleteWordCard(c *gin.Context) {
	id := c.Param("id")
	err := models.DeleteWordCard(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete word card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Word card deleted"})
}
