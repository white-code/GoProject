package model

import (
	"errors"
	"fmt"
)

var (NotfindBookerr = errors.New("this book is not borrowed by you"))

type BorrowItem struct{
	book *Book
	num int
}
type Student struct{
	Name string
	Grade string
	Id string
	Sex string
	Books []*BorrowItem
}

func CreateStudent(name,grade,id,sex string) *Student{
	stu := &Student{
		Name:name,
		Grade:grade,
		Id:id,
		Sex:sex,
	}
	return stu
}

func (s *Student) AddBook(b *BorrowItem) {
	s.Books = append(s.Books,b)
}
func (s *Student) DelBook(b *BorrowItem) (err error){
	for i :=0 ; i<len(s.Books);i++{
		if s.Books[i].book.Name == b.book.Name {
			if b.num == s.Books[i].num{
				left := s.Books[0:i]
				right := s.Books[i+1:]
				s.Books = append(left,right...)
				return
			}else {
				s.Books[i].num -= b.num
				return
			}
		}
	}
	err = NotfindBookerr
	return
}

func (s *Student) GetBooks(b Student){
	for i := 0;i<len(b.Books);i++{
		fmt.Println(b.Books[i])
	}
}


