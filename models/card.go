package models

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type WordCard struct {
	ID    int    `json:"id,omitempty"`
	Front string `json:"front"`
	Back  string `json:"back"`
}

// 获取所有单词卡
func GetAllWordCards() ([]WordCard, error) {
	var wordCards []WordCard
	rows, err := db.Query("SELECT * FROM word_cards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var wordCard WordCard
		err := rows.Scan(&wordCard.ID, &wordCard.Front, &wordCard.Back)
		if err != nil {
			return nil, err
		}
		wordCards = append(wordCards, wordCard)
	}
	return wordCards, nil
}

// 创建单词卡
func InsertWordCard(wordCard *WordCard) error {
	r, err := db.Exec("INSERT INTO word_cards (front, back) VALUES (?, ?)", wordCard.Front, wordCard.Back)
	p, _ := r.LastInsertId()
	fmt.Println("rrrr", p)
	if err != nil {
		return err
	}
	return nil
}

// 更新单词卡
func UpdateWordCard(id string, wordCard *WordCard) error {
	_, err := db.Exec("UPDATE word_cards SET front = ?, back = ? WHERE id = ?", wordCard.Front, wordCard.Back, id)
	if err != nil {
		return err
	}
	return nil
}

// 删除单词卡
func DeleteWordCard(id string) error {
	_, err := db.Exec("DELETE FROM word_cards WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
