package main

import (
    "os"
    "log"
    "testing"
)

var a App

func TestMain(m *testing.M) {
    a = App{}
    a.Initialize("root", "", "go-api")

    ensureTableExists()

    code := m.Run()

    clearTable()

    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM recipes")
    a.DB.Exec("ALTER TABLE recipes AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS recipes
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    ingredients TEXT DEFAULT NULL
)`