package model

type Stock struct {
	ID       uint64 `json:"id" gorm:"primarykey"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Ticker   string `json:"ticker" gorm:"type:varchar(255)"`
	Cik      string `json:"cik" gorm:"type:varchar(255)"`
	Stype    string `json:"stype" gorm:"type:varchar(255)"`
	Sic      uint64 `json:"sic" gorm:"index"`
	Sector   string `json:"sector" gorm:"type:varchar(255)"`
	Delisted bool   `json:"delisted,omitempty" gorm:"type:boolean;default:false"`

	BaseModel
}
