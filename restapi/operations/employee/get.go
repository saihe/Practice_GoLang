package employee

import (
	"practiceGoLang/domain"
	"practiceGoLang/models"
	"practiceGoLang/restapi/operations"
	"strconv"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func EmployeeID(params operations.GetEmployeeIDParams) middleware.Responder {
	id, _ := strconv.ParseUint(params.ID, 10, 64)
	employeeEntity := domain.FindById(id)
	employee := &models.Employee{
		ID:           strconv.FormatUint(employeeEntity.ID, 10),
		Name:         employeeEntity.Name,
		DepartmentID: int64(employeeEntity.DepartmentID),
		CreatedAt:    strfmt.DateTime(employeeEntity.CreatedAt),
		UpdatedAt:    strfmt.DateTime(employeeEntity.UpdatedAt),
	}

	payload := &operations.GetEmployeeIDOKBody{
		Employee: employee,
	}

	return operations.NewGetEmployeeIDOK().WithPayload(payload)
}
