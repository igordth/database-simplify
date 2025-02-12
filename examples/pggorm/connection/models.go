package main

type BaseModel struct {
	ID   uint   `gorm:"primary_key;column:id"`
	Name string `gorm:"column:name"`
}

type Galaxy struct {
	BaseModel
}

func (Galaxy) TableName() string {
	return "galaxies"
}

type Stars struct {
	BaseModel
	GalaxyId *int `gorm:"column:galaxy_id"`
	Galaxy   *Galaxy
}

func (Stars) TableName() string {
	return "stars"
}

type Planet struct {
	BaseModel
	StarId *int `gorm:"column:star_id"`
	Star   *Stars
}

func (Planet) TableName() string {
	return "planets"
}
