package form

import (
	"camarinb2096/form_example/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Post  Controller
		Patch Controller
	}

	Form struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		City      string `json:"city"`
		CompType  string `json:"type"`
		Complaint string `json:"description"`
	}
)

func NewEndpoints(s Services) Endpoints {
	return Endpoints{
		Post: createForm(s),
		Patch: func(c *gin.Context) {
			utils.HandleError(c, http.StatusNotImplemented, "Not implemented")
		},
	}
}

func createForm(s Services) Controller {
	return func(c *gin.Context) {
		var form Form
		json.NewDecoder(c.Request.Body).Decode(&form)

		if form.Name == "" || form.Email == "" || form.City == "" || form.Complaint == "" {
			utils.HandleError(c, http.StatusInternalServerError, "Invalid form")
			return
		}
		err := s.CreatePqr(form.Name, form.Email, form.City, form.Complaint)
		if err != nil {
			utils.HandleError(c, http.StatusInternalServerError, err.Error())
			return

		}
		utils.HandleSuccess(c, "Form submitted successfully", form)
	}
}
