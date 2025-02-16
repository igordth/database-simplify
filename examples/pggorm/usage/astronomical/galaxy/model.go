package galaxy

type Model struct {
	ID   uint   `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}

func (Model) TableName() string {
	return "galaxies"
}
