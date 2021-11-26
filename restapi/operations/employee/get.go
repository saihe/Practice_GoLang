package employee

import (
	"practiceGoLang/models"
	"practiceGoLang/restapi/operations"
	"time"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

func EmployeeID(params operations.GetEmployeeIDParams) middleware.Responder {
	employee := &models.Employee{
		ID:           params.ID,
		Name:         "Kyohei Saito",
		DepartmentID: 1,
		CreatedAt:    strfmt.DateTime(time.Now()),
		UpdatedAt:    strfmt.DateTime(time.Now()),
	}

	payload := &operations.GetEmployeeIDOKBody{
		Employee: employee,
	}

	return operations.NewGetEmployeeIDOK().WithPayload(payload)
}
