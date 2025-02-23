package services

import (
	"github.com/fujihara/go-api/models"
	"github.com/fujihara/go-api/repositories"
)

// PostCommentHandlerで使うサービス
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
