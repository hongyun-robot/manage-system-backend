package models

import "gorm.io/gorm"

type Article struct {
	ID        uint64         `gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"size:64"`
	Content   string         `json:"content" gorm:"type:MEDIUMTEXT"`
	CreateAt  int64          `json:"createAt" gorm:"autoCreateTime:milli;<-:create"`
	UpdateAt  int64          `json:"updateAt" gorm:"autoUpdateTime:milli"`
	DeleteAt  gorm.DeletedAt `json:"deleteAt" gorm:"index"`
	Status    uint32         `json:"status"`
	Draft     bool           `json:"draft" gorm:"default=true"`
	Classifys []*Classify    `gorm:"many2many:article_classifys;"`
}

func (a *Article) TableName() string {
	return "article"
}

type Classify struct {
	ID       uint64     `gorm:"primaryKey"`
	Name     string     `json:"name" gorm:"size:64"`
	CreateAt int64      `json:"createAt" gorm:"autoCreateTime:milli;<-:create"`
	UpdateAt int64      `json:"updateAt" gorm:"autoUpdateTime:milli"`
	Icon     string     `json:"icon" gorm:"size:64"`
	Articles []*Article `gorm:"many2many:article_classifys;"`
}

func (receiver *Classify) TableName() string {
	return "classify"
}

type Menu struct {
	Id       uint64  `gorm:"primaryKey"`
	ParentId *uint64 `json:"parentId" gorm:"index"`
	Order    uint32  `json:"order"`
	Title    string  `json:"title" gorm:"size:128"`
	Icon     string  `json:"icon" gorm:"size:64"`
	Url      string  `json:"url"`
	Disabled uint32  `json:"disabled"`
	Hide     bool    `json:"hide"`
	Children []*Menu `gorm:"foreignKey:ParentId;association_foreignkey:Id"`
}

func (receiver *Menu) TableName() string {
	return "menu"
}