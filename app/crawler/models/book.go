package models

import (
	"fmt"
	"os"
)

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

func (b Book) WriteBookToFile() {

	wd, _ := os.Getwd()
	path := wd + "\\static"
	err := os.MkdirAll(path, 0660)
	if err != nil {
		fmt.Println("创建文件夹失败,", err.Error())
		return
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		if !os.IsNotExist(err) {
			file, err = os.Create(path + "\\books.txt")
			if err != nil {
				fmt.Println("创建 books.txt 文件失败,", err.Error())
				return
			}
		}
	}
	if file != nil {
		defer file.Close()
	}

	bookItemStr := fmt.Sprintf("书名：%v\n作者:%v\n出版社:%v\n出版时间:%v\n评分:%v\n页码:%v\n定价:%v\n内容简介:%v\n\n",
		b.Name, b.Author, b.Press, b.PublicationTime, b.RatingNum, b.PageNum, b.Price, b.Info)

	file.WriteString(bookItemStr)
}
