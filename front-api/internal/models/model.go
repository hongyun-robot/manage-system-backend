package models

type Article struct {
	ID           uint        `gorm:"primaryKey"`
	Title        string      `json:"title"`
	Content      string      `json:"content"`
	CreateAt     int64       `json:"createAt" gorm:"autoCreateTime:milli"`
	UpdateAt     int64       `json:"updateAt" gorm:"autoUpdateTime:milli"`
	DeleteAt     int64       `json:"deleteAt" gorm:"index"`
	Status       uint        `json:"status"`
	ClassifyList []Classify  `json:"classifyList" gorm:"-"`
	Classifys    []*Classify `json:"classifys" gorm:"many2many:article_classifys;"`
}

type Classify struct {
	ID          uint       `gorm:"primaryKey"`
	Name        string     `json:"name"`
	CreateAt    int64      `json:"createAt" gorm:"autoCreateTime:milli"`
	UpdateAt    int64      `json:"updateAt" gorm:"autoUpdateTime:milli"`
	Icon        string     `json:"icon"`
	ArticleList []Article  `json:"articleList" gorm:"-"`
	Articles    []*Article `json:"articles" gorm:"many2many:article_classifys;"`
}
