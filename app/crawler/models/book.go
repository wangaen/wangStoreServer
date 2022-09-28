package models

import "fmt"

type Book struct {
	BookId          int     `json:"bookId" gorm:"primaryKey;autoIncrement;comment:书籍Id"`
	Name            string  `json:"name" gorm:"size:128,comment:书名"`
	Author          string  `json:"author" gorm:"size:64,comment:作者"`
	Press           string  `json:"press" gorm:"size:64,comment:出版社"`
	PublicationTime string  `json:"publicationTime" gorm:"size:32,comment:出版时间"`
	PageNum         int     `json:"pageNum" gorm:"size:10,comment:页码"`
	RatingNum       float64 `json:"ratingNum" gorm:"size:10,comment:评分"`
	Price           string  `json:"price" gorm:"size:16,comment:定价"`
	Info            string  `json:"info" gorm:"size:1024,comment:内容简介"`
}

func (b Book) PrintBookDetails() {
	fmt.Printf("书名：%v\n作者:%v\n出版社:%v\n出版时间:%v\n评分:%v\n页码:%v\n定价:%v\n内容简介:%v\n",
		b.Name, b.Author, b.Press, b.PublicationTime, b.RatingNum, b.PageNum, b.Price, b.Info)
}
