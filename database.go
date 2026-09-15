package main

import (
	"database/sql"
	"time"
	_"modernc.org/sqlite"
)

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite", "tasks.db")
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			priority TEXT NOT NULL,
			deadline TEXT NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT 0
		)
	`)
	return err
}

func addTask(task todo) error {
	_, err := db.Exec(`
		INSERT INTO tasks (name, priority, deadline, completed)
		VALUES (?, ?, ?, ?)
	`,
		task.name,
		task.priority,
		task.deadline.Format(time.RFC3339),
		task.completed,
)
	return err
}

func getTasks() ([]todo, error) {
	rows, err := db.Query(`
		SELECT name, priority, deadline, completed
		FROM tasks
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []todo

	for rows.Next() {
		var task todo
		var deadline string
		err := rows.Scan(
			&task.name,
			&task.priority,
			&deadline,
			&task.completed,
		)
		if err != nil {
			return nil, err
		}
		task.deadline, err = time.Parse(time.RFC3339, deadline)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, rows.Err()
}

func deleteTask(name string) error {
	_, err := db.Exec(`
		DELETE FROM tasks
		WHERE name = ?
	`, name)
	return err
}

func clearTasks() error {
	_, err := db.Exec(`
		DELETE FROM tasks
	`)
	return err
}

func completeTask(name string) error {
	_, err := db.Exec(`
		UPDATE tasks
		SET completed = 1
		WHERE name = ?
	`, name)
	return err
}
