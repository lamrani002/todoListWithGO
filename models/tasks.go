// models/tasks.go

package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// creation du type de la tache en utilisant les structures
// [`json:{column}`] permet la conversion en json direct

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Etat string `json:"etat"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// fonction recuperer les donnes

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	// effacer une fois le programe est fermé
	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		erroRows := rows.Scan(&task.ID, &task.Name, &task.Etat)

		if erroRows != nil {
			panic(erroRows)
		}
		//ajouter les info dans la table a partir de la DB
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

// fonction pour chercher une tache

func FindTaskById(db *sql.DB, id int) Task {

	sql := "SELECT * FROM tasks WHERE id = ?"
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	// effacer une fois le programe est fermé
	defer rows.Close()

	task := Task{}

	for rows.Next() {
		erroRows := rows.Scan(&task.ID, &task.Name, &task.Etat)

		if erroRows != nil {
			panic(erroRows)
		}
		//ajouter les info dans la table a partir de la DB

	}
	return task
}

// fonction pour l'insertion

func AddTask(db *sql.DB, name string, etat string) (int64, error) {
	sql := "INSERT INTO tasks(name,etat) VALUES(?,?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}
	// effacer une fois le programe est fermé
	defer stmt.Close()

	// executer la requete en passant la valeur de {name}
	result, err2 := stmt.Exec(name, etat)
	// femer si une erreur est trouvée
	if err2 != nil {
		panic(err2)
	}

	// recuperer la derniere ligne insérée
	return result.LastInsertId()
}

// fonction pour la suppression

func DeleteTask(db *sql.DB, id int) (int64, error) {

	sql := "DELETE FROM tasks WHERE id = ?"
	stmt, err := db.Prepare(sql)

	// femer si une erreur est trouvée
	if err != nil {
		panic(err)
	}
	result, err2 := stmt.Exec(id)

	// femer si une erreur est trouvée
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

// fonction pour modifier

func EditTask(db *sql.DB, id int, name string, etat string) Task {

	findId := "UPDATE tasks SET name= ? ,etat= ? WHERE id= ?"
	stmt, err1 := db.Prepare(findId)

	if err1 != nil {
		panic(err1)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, etat, id)

	if err2 != nil {
		panic(err2)
	}

	result.RowsAffected()
	task := FindTaskById(db, id)

	return task
}

func SearchForTask(db *sql.DB, search string) TaskCollection {

	sql := "SELECT * FROM tasks WHERE name like '%" + search + "%'"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	// effacer une fois le programe est fermé
	defer rows.Close()

	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		erroRows := rows.Scan(&task.ID, &task.Name, &task.Etat)

		if erroRows != nil {
			panic(erroRows)
		}
		//ajouter les info dans la table a partir de la DB
		result.Tasks = append(result.Tasks, task)
	}
	return result
}
