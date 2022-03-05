package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Album struct {
	AlbumId  int
	Title    string
	ArtistId int
}

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(192.168.1.18)/Chinook")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	print("albums :")
	rows, _ := db.Query("SELECT * FROM Album")
	var album = Album{}
	for rows.Next() {
		rows.Scan(&album.AlbumId, &album.Title, &album.ArtistId)
		print(album.AlbumId)
		print(album.Title)
	}

	print("artist 75 :")
	stmtOut, _ := db.Prepare("SELECT Name FROM Artist WHERE ArtistId = ?")
	defer stmtOut.Close()
	var vinicius string
	stmtOut.QueryRow(75).Scan(&vinicius)
	print(vinicius)
}
