package model

type Ccgp struct {
	Type    int    `gorm:"index;column:type;type:int(1)"`
	Title   string `gorm:"index;column:title;type:varchar(1024)"`
	Url     string `gorm:"unique_index;column:url;type:varchar(1024)"`
	Html    string `gorm:"column:html;type:longtext"`
	Content string `gorm:"column:content;type:longtext"`
}

type Data struct {
	ID     string
	Type   int
	Title  string
	OriUrl string
}
