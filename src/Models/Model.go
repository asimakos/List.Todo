
package Models

type Todo struct {
	ID uint             `gorm:"primary_key;AUTO_INCREMENT" `
	Title string        `gorm:"not null" json:"title" `
	Description string  `gorm:"not null" json:"description" `
}

func (b Todo) SliceDescription() string {

	return string(b.Description[:45]) 
}

func (b *Todo) TableName() string {
	return "todo"
}