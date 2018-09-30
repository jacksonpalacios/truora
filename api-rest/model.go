package main

import (
    "database/sql"	
	"fmt"
)

type recipe struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Description   string    `json:"description"`
}

func (r *recipe) getRecipe(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT name, description FROM recipes WHERE id=%d", r.ID)
    return db.QueryRow(statement).Scan(&r.Name, &r.Description)
}

func (r *recipe) updateRecipe(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE recipes SET name='%s', description='%s' WHERE id=%d", r.Name, r.Description, r.ID)
	_, err := db.Exec(statement)
    return err
}

func (r *recipe) deleteRecipe(db *sql.DB) error {
    statement := fmt.Sprintf("DELETE FROM recipes WHERE id=%d", r.ID)
    _, err := db.Exec(statement)
    return err
}

func (r *recipe) createRecipe(db *sql.DB) error {
    statement := fmt.Sprintf("INSERT INTO recipes(name, description, ingredients) VALUES('%s', '%s', '%s')", r.Name, r.Description, "")
    _, err := db.Exec(statement)

    if err != nil {
        return err
    }

    err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&r.ID)

    if err != nil {
        return err
    }

    return nil
}

func getRecipes(db *sql.DB) ([]recipe, error) {
    //statement := fmt.Sprintf("SELECT id, name, description FROM recipes LIMIT %d OFFSET %d", count, start)
    statement := fmt.Sprintf("SELECT id, name, description FROM recipes ORDER BY id DESC;")
    rows, err := db.Query(statement)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    recipes := []recipe{}

    for rows.Next() {
        var r recipe
        if err := rows.Scan(&r.ID, &r.Name, &r.Description); err != nil {
            return nil, err
        }
        recipes = append(recipes, r)
    }

    return recipes, nil
}