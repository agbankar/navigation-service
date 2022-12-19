package v1

import (
	"encoding/json"
	"fmt"
	"github.com/agbankar/navigation-service/internal/model"
	"github.com/agbankar/navigation-service/internal/respond"
	"github.com/agbankar/navigation-service/internal/service"
	"io"
	"net/http"
)

type VisitorController struct {
	VisitorService service.VisitorService
}

func NewVisitorController(VisitorService service.VisitorService) *VisitorController {
	return &VisitorController{VisitorService: VisitorService}
}

func (v *VisitorController) Visit(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error occurred while reading the body %v\n", err.Error())
		respond.With(w, r, http.StatusInternalServerError, model.ApiResponse{
			ErrorMessage: "Something went wrong please try again",
		})
		return
	}
	u := model.User{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Printf("Error occurred while unmarshaling %v\n", err.Error())
		respond.With(w, r, http.StatusBadRequest, model.ApiResponse{
			ErrorMessage: "Not a valid request, Please check the input",
		})
		return
	}
	if len(u.UserId) == 0 || len(u.Url) == 0 {
		respond.With(w, r, http.StatusBadRequest, model.ApiResponse{
			ErrorMessage: "UserId or url can not be blank",
		})
		return
	}
	err = v.VisitorService.Visit(&u)
	if err != nil {
		respond.With(w, r, http.StatusInternalServerError, model.ApiResponse{
			ErrorMessage: "Something went wrong ",
		})
		return
	}
	respond.With(w, r, http.StatusCreated, model.ApiResponse{
		SuccessMessage: "Ok",
	})
}
func (v *VisitorController) GetUniqueVisits(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	counter := v.VisitorService.GetUniqueVisits(url)
	respond.With(w, r, http.StatusOK, model.ApiResponse{
		Page:         url,
		UniqueVisits: &counter,
	})
}
