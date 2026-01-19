package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

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
