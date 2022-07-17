package repository

import (
	"back-usm/internals/comments/core/domain"
	"log"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CommentRepository struct {
	dsn string
	db  *gorm.DB
}

func NewCommentRepository(dsn string) *CommentRepository {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(domain.Comment{})

	log.Print(color.GreenString("Comments repository connected to database"))
	return &CommentRepository{
		dsn: dsn,
		db:  db,
	}
}

func (r *CommentRepository) Create(comment *domain.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetAll(productId string) ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Where("product_id = ?", productId).Order("created_at desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) GetLast(productId string) ([]domain.Comment, error) {
	var comments []domain.Comment

	err := r.db.Where("product_id = ?", productId).Order("created_at desc").Limit(3).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&domain.Comment{}).Error
}
