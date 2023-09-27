package comment

import "github.com/jinzhu/gorm"

//Service -- the struct for our comment service

type Service struct{
	DB *gorm.DB
}


//CommentService --- the interface for our comment service
type CommentService interface{
	       
}

//NewService --- returns a new comment service 

func NewService(db *gorm.DB) *Service{
	return &Service{
		DB: db,
	}
}