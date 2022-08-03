//handlers/tasks.go

package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lamrani002/todoList/models"
)

type H map[string]interface{}

//fonction por recuperer tous les taches en utilisant l'interface Echo.HandlerFunc

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// Creation d'une tache

func AddTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var task models.Task
		ctx.Bind(&task)

		id, err := models.AddTask(db, task.Name, task.Etat)

		if err == nil {
			return ctx.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

// fonction pour trouver une tache a travers ID
func FindTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err == nil {
			return ctx.JSON(http.StatusOK, models.FindTaskById(db, id))

		} else {
			return err
		}

	}
}

//fonction pour modifier
func UpdateTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		id, _ := strconv.Atoi(ctx.Param("id"))

		task := models.FindTaskById(db, id)
		var err error
		if err != nil {
			return err
		}
		ctx.Bind(&task)
		taskUpt := models.EditTask(db, id, task.Name, task.Etat)
		var err2 error
		if err2 == nil {
			return ctx.JSON(http.StatusOK, H{
				"task": taskUpt,
			})
		} else {
			return ctx.JSON(http.StatusNotFound, H{
				"message": "probleme DB",
			})
		}

	}
}

// supprimer une tache

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// permet de caster id en type integer
		id, _ := strconv.Atoi(ctx.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			return ctx.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}

// chercher une tache

func SearchTask(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		search := ctx.QueryParam("name")

		return ctx.JSON(http.StatusOK, models.SearchForTask(db, search))

	}
}
