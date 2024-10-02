package data

type Student_info struct {
	ID         uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string  `gorm:"size:100;not null"`
	Age        string  `gorm:"size:100"`
	University string  `gorm:"size:100"`
	CGPA       float64 `gorm:"size:100"`
	Gender     string  `gorm:"size:100"`
}
