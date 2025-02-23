package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "inies"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	article_id := 1
	const sqlGetNice = `
		select nice 
		from articles 
		where article_id = ?
	`

	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()

	// r := mux.NewRouter()

	// r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	// r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	// r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// // http.HandleFunc("/hello", handlers.HelloHandler)
	// // http.HandleFunc("/article", handlers.PostArticleHandler)
	// // http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// // http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// // http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// // http.HandleFunc("/comment", handlers.PostCommentHandler)

	// log.Println("server start port 8080")
	// log.Fatal(http.ListenAndServe(":8080", r))
}
