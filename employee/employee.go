package employee

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id           int
	Name         string
	DepartmentId int
	CreatedAt    string
	UpdatedAt    string
}

func Get(id int) (Employee, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sandbox")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var employee Employee
	err = db.QueryRow("SELECT * FROM employee where id = ?", id).Scan(
		&employee.Id,
		&employee.Name,
		&employee.DepartmentId,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	return employee, err
}
