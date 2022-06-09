package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Insert(Marca string, Nome string, Placa string, Preco float64, Vendido bool) (int, error) {
	db, err := sql.Open("mysql", "root:1234@/poo_avancado_web")

	if err != nil {
		return 0, err
	}

	var placa string
	db.QueryRow("select placa from carros where placa = ?", Placa).Scan(&placa)

	if placa != "" {
		return 0, nil
	}

	stmt := "insert into carros (marca, nome, placa, preco, vendido) values (?, ?, ?, ?, ?)"
	result, err := db.Exec(stmt, Marca, Nome, Placa, Preco, Vendido)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func Update(Id int, Marca string, Nome string, Placa string, Preco float64, Vendido bool) (bool, error) {
	db, err := sql.Open("mysql", "root:1234@/poo_avancado_web")

	if err != nil {
		return false, err
	}

	stmt := "update carros set marca = ?, nome = ?, placa = ?, preco = ?, vendido = ? where id = ?"
	result, err := db.Exec(stmt, Marca, Nome, Placa, Preco, Vendido, Id)

	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rows > 0, nil
}

func Delete(Id int) (bool, error) {
	db, err := sql.Open("mysql", "root:1234@/poo_avancado_web")

	if err != nil {
		return false, err
	}

	stmt := "delete from carros where id = ?"
	result, err := db.Exec(stmt, Id)

	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func List() (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root:1234@/poo_avancado_web")

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select * from carros")

	if err != nil {
		return nil, err
	}

	return rows, nil
}
