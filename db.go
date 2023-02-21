package main

import (
    "database/sql"  // database interface
    "fmt"
    "log"
    _"github.com/lib/pq"  // postgresql driver
)
const (
    host     = "localhost"
    port     = 6500
    user     = "postgres"
    password = "adminpswd"
    dbname   = "postgres"
)

var dbconn *sql.DB
// album model
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

func ConnectDB() (error) {
    connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)
    db, err := sql.Open("postgres", connString)
    if err != nil {
        log.Printf("failed to connect to database: %v", err)
        return err
    }
	dbconn = db
    return nil
}

func allAlbums() ([]album, error) {
    var albums []album
    rows, err := dbconn.Query("SELECT * FROM albums")
    if err != nil {
        return nil, fmt.Errorf("getAllAlbums: %v", err)
    }
    defer rows.Close()
    for rows.Next() {
        var alb album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("getAllAlbums: %v", err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getAllAlbums: %v", err)
    }
    return albums, nil
}

func addAlbum(alb album) (string, error) {
	q := "INSERT INTO albums (title, artist, price) VALUES ($1,$2,$3) RETURNING id"
    var id int
	err := dbconn.QueryRow(q, alb.Title, alb.Artist, alb.Price).Scan(&id)
    if err != nil {
        return "", fmt.Errorf("addAlbum: %v", err)
    }
    return fmt.Sprint(id), nil
}

func albumById(id int) (album, error) {
    var alb album
	q := "SELECT * FROM albums WHERE id=$1;"
    row := dbconn.QueryRow(q, id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumById %d: %v", id, err)
    }
    return alb, nil
}