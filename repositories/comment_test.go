package repositories_test

import (
	"testing"

	"github.com/fujihara/go-api/models"
	"github.com/fujihara/go-api/repositories"
	_ "github.com/go-sql-driver/mysql"
)

// SelectCommentListのテスト
func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got %d\n", articleID, comment.ArticleID)
		}
	}
}

// InsertCommentのテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "CommentInsertTest",
	}

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments 
			where message = ?
		`

		testDB.Exec(sqlStr, comment.Message)
	})
}
