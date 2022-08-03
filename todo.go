// to do

package main

import (
	"database/sql"

	"github.com/labstack/echo"
	"github.com/lamrani002/todoList/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// appel de la function pour initialiser
	db := initDB("storage.db")
	migrate(db)

	// instantiation du module Echo pour creer les routes

	e := echo.New()

	//la page du front end
	e.File("/", "public/index.html")

	e.GET("/tasks", handlers.GetTasks(db))
	e.GET("/tasks/:id", handlers.FindTask(db))
	e.GET("/tasks/search", handlers.SearchTask(db))
	e.PUT("/tasks", handlers.AddTask(db))
	e.PATCH("/tasks/edit/:id", handlers.UpdateTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// lancer le serveur

	e.Logger.Fatal(e.Start(":8080"))
}

//function pour connection a la Base de donné

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	//pour tester s'il y a des erreurs
	if err != nil {
		panic(err)
	}

	//pour verifier si la connection DB a reussi

	if db == nil {
		panic("db nil")
	}

	return db
}

// fonction pour creation de table

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,name VARCHAR NOT NULL,etat VARCHAR NOT NULL);`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
