package handlers

import (
	"net/http"
	"word-card-app/models"

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
	err := models.InsertWordCard(&wordCard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create word card"})
		return
	}
	c.JSON(http.StatusCreated, wordCard)
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
