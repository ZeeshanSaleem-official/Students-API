package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/ZeeshanSaleem-official/student-api/internal/config"
	"github.com/ZeeshanSaleem-official/student-api/internal/config/types"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Postgresql, error) {
	// Open Connection using the postgres driver

	db, err := sql.Open("postgres", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	//Ping to verify connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	//Create Table in Database
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER
	)`)
	if err != nil {
		return nil, err
	}
	return &Postgresql{
		Db: db,
	}, nil

}
func (p *Postgresql) CreateStudent(name string, email string, age int) (int64, error) {
	//Use $1, $2, $3 for dymanic values
	query := "INSERT INTO students (name, email, age) VALUES ($1, $2, $3) RETURNING id"
	var id int64
	err := p.Db.QueryRow(query, name, email, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (p *Postgresql) StudentGetById(id int64) (types.Student, error) {
	stmt, err := p.Db.Prepare(`SELECT id, name, email, age FROM students WHERE id = $1 LIMIT 1`)
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		//User not Found
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}
	return student, nil
}

func (p *Postgresql) GetAllStudents() ([]types.Student, error) {
	stmt, err := p.Db.Prepare(`SELECT id, name, email, age FROM students`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var students []types.Student
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
