package models

import (
	"encoding/json"
	"fmt"
)

// User 用户基本信息
type User struct {
	ID            string `json:"id" gorm:"primaryKey;comment:id"`
	Name          string `json:"name" gorm:"size:128,comment:用户名"`
	Sex           string `json:"sex" gorm:"size:16,comment:性别"`
	Age           string `json:"age" gorm:"size:16,comment:年龄"`
	Height        string `json:"height" gorm:"size:16,comment:身高"`
	Weight        string `json:"weight" gorm:"size:16,comment:体重"`
	Salary        string `json:"salary" gorm:"size:64,comment:月收入"`
	Status        string `json:"status" gorm:"size:16,comment:婚况"`
	XingZuo       string `json:"xingZuo" gorm:"size:16,comment:星座"`
	XueLi         string `json:"xueLi" gorm:"size:64,comment:学历"`
	Work          string `json:"work" gorm:"size:64,comment:职业"`
	WorkAddress   string `json:"workAddress" gorm:"size:64,comment:工作地"`
	Signature     string `json:"signature" gorm:"size:1024,comment:内心独白"`
	GirlCondition GirlCondition
}

// GirlCondition 择偶条件
type GirlCondition struct {
	Age         string `json:"age" gorm:"size:32,comment:年龄"`
	Height      string `json:"height" gorm:"size:32,comment:身高"`
	Salary      string `json:"salary" gorm:"size:64,comment:月收入"`
	XueLi       string `json:"xueLi" gorm:"size:64,comment:学历"`
	WorkAddress string `json:"workAddress" gorm:"size:64,comment:工作地"`
}

func (u User) PrintUserDetails() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("|	  用户基本信息：")
	fmt.Printf("|	\t用户名: %v\n", u.Name)
	fmt.Printf("|	\t性别: %v\n", u.Sex)
	fmt.Printf("|	\t年龄: %v\n", u.Age)
	fmt.Printf("|	\t身高: %v\n", u.Height)
	fmt.Printf("|	\t体重: %v\n", u.Weight)
	fmt.Printf("|	\t月收入: %v\n", u.Salary)
	fmt.Printf("|	\t婚况: %v\n", u.Status)
	fmt.Printf("|	\t星座: %v\n", u.XingZuo)
	fmt.Printf("|	\t学历: %v\n", u.XueLi)
	fmt.Printf("|	\t职业: %v\n", u.Work)
	fmt.Printf("|	\t工作地: %v\n", u.WorkAddress)
	fmt.Printf("|	\t内心独白: %v\n", u.Signature)
	fmt.Println("|	  择偶条件：")
	fmt.Printf("|	\t年龄: %v\n", u.Age)
	fmt.Printf("|	\t身高: %v\n", u.Height)
	fmt.Printf("|	\t月收入: %v\n", u.Salary)
	fmt.Printf("|	\t学历: %v\n", u.XueLi)
	fmt.Println("-------------------------------------------------")
}

func (u User) WriteUserToFile() {

	//wd, _ := os.Getwd()
	//path := wd + "\\static\\book.txt"
	////err := os.MkdirAll(path, 0660)
	////if err != nil {
	////	fmt.Println("创建文件夹失败,", err.Error())
	////	return
	////}
	//
	//file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("book.txt 打开异常：", err.Error())
	//	return
	//	//if !os.IsNotExist(err) {
	//	//	file, err = os.Create(path + "\\books.txt")
	//	//	if err != nil {
	//	//		fmt.Println("创建 books.txt 文件失败,", err.Error())
	//	//		return
	//	//	}
	//	//}
	//}
	//defer file.Close()
	//
	//bookItemStr := fmt.Sprintf("书名：%v\n作者:%v\n出版社:%v\n出版时间:%v\n评分:%v\n页码:%v\n定价:%v\n内容简介:%v\n\n",
	//	b.Name, b.Author, b.Press, b.PublicationTime, b.RatingNum, b.PageNum, b.Price, b.Info)
	//
	//file.WriteString(bookItemStr)
}

func FormatUserObj(obj interface{}) (User, error) {
	var user User
	str, err := json.Marshal(obj)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(str, &user)
	return user, err
}

func FormatGirlConditionObj(obj interface{}) (GirlCondition, error) {
	var girlCondition GirlCondition
	str, err := json.Marshal(obj)
	if err != nil {
		return girlCondition, err
	}
	err = json.Unmarshal(str, &girlCondition)
	return girlCondition, err
}
