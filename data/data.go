package data

type student_info struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	Age        string `gorm:"size:100"`
	university string `gorm:"size:100"`
}
