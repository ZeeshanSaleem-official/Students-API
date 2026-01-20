package storage

import "github.com/ZeeshanSaleem-official/student-api/internal/config/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	StudentGetById(id int64) (types.Student, error)
	GetAllStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) (types.Student, error)
	DeleteStudent(id int64) error
}
