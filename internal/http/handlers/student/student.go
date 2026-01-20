package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ZeeshanSaleem-official/student-api/internal/config/types"
	"github.com/ZeeshanSaleem-official/student-api/internal/config/utils/response"
	"github.com/ZeeshanSaleem-official/student-api/internal/storage"
	"github.com/go-playground/validator"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a new Student")
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		// Empty Body error
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrors(fmt.Errorf("Empty Body")))
			return
		}
		// General Error
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrors(err))
		}

		//request Validation
		err = validator.New().Struct(student)
		//type casting
		if err != nil {
			val_error := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(val_error))
			return
		}

		id, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)
		slog.Info("User created successfully!!", slog.String("userId", fmt.Sprint(id)))
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err)
			return
		}
		response.WriteJson(w, http.StatusCreated, map[string]int64{"success": id})
	}
}
func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("getting a student from DB by id", slog.String("id: ", id))
		// for converting into int64
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrors(err))
			return
		}
		student, err := storage.StudentGetById(intID)
		if err != nil {
			slog.Error("user getting error", slog.String("id: ", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErrors(err))
			return
		}
		response.WriteJson(w, http.StatusOK, student)

	}
}
func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Getting all users!!")
		students, err := storage.GetAllStudents()
		if err != nil {
			slog.Error("error getting all users %w", students)
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}
		response.WriteJson(w, http.StatusOK, students)
	}
}

func Update(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Updating a user!!!")
		id := r.PathValue("id")
		// for converting into int64
		intID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Info("Error during conversion into int64 in updating user")

			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErrors(err))
			return
		}
		// DECODE
		var req types.UpdateStudent
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrors(err))
		}
		// VALIDATE
		if err := validator.New().Struct(req); err != nil {
			var valErrors validator.ValidationErrors
			if errors.As(err, &valErrors) {
				response.WriteJson(w, http.StatusBadRequest, response.ValidationError(valErrors))
				return
			}
			return
		}

		existingStudent, err := storage.StudentGetById(intID)
		if err != nil {
			slog.Info("Error in getting student for updating", slog.String("id", id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralErrors(err))
			return
		}
		// Merge Data
		if req.Name != nil {
			slog.Info("Name getting in updating", slog.String("name", *req.Name))
			existingStudent.Name = *req.Name
		}
		if req.Email != nil {
			slog.Info("Email getting in updating", slog.String("name", *req.Email))
			existingStudent.Email = *req.Email
		}
		if req.Age != nil {
			slog.Info("Age getting in updating", slog.Int("Age", *req.Age))

			existingStudent.Age = *req.Age
		}
		//save
		updatedStudent, err := storage.UpdateStudent(
			intID,
			existingStudent.Name,
			existingStudent.Email,
			existingStudent.Age,
		)
		slog.Info("updating a user!!!", slog.String("Student", updatedStudent.Name))
		slog.Info("updating a user!!!", slog.String("Student", updatedStudent.Email))
		slog.Info("Age getting in updating", slog.Int("Age", updatedStudent.Age))

		if err != nil {
			slog.Info("Error during updating a user!!!")
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralErrors(err))
			return
		}
		response.WriteJson(w, http.StatusOK, updatedStudent)
	}

}
