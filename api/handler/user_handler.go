package handler

import (
	"net/http"
	"strings"

	"github.com/bytebury/fun-banking/internal/domain"
	"github.com/bytebury/fun-banking/internal/infrastructure/persistence"
	"github.com/bytebury/fun-banking/internal/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler() userHandler {
	return userHandler{
		userService: service.NewUserService(),
	}
}

func (h userHandler) Count(c *gin.Context) {
	c.HTML(http.StatusOK, "users_count", h.userService.Count())
}

func (h userHandler) SignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "users/signup", nil)
}

func (h userHandler) Create(c *gin.Context) {
	c.Request.ParseForm()

	formData := newFormData()

	for key, values := range c.Request.PostForm {
		if len(values) > 0 {
			formData.Data[key] = values[0]
		}
	}

	// TODO: we need to validate the form at this point

	if formData.Data["password"] != formData.Data["password_confirmation"] {
		formData.Errors["general"] = "Passwords provided do not match"
		formData.Errors["passwords_dont_match"] = "Passwords provided do not match"
		c.HTML(http.StatusUnprocessableEntity, "users/signup_form", formData)
		return
	}

	if err := persistence.DB.Create(&domain.User{
		FirstName: strings.ToLower(formData.Data["first_name"]),
		LastName:  strings.ToLower(formData.Data["last_name"]),
		Email:     strings.ToLower(formData.Data["email"]),
		Username:  strings.ToLower(formData.Data["username"]),
		Password:  formData.Data["password"],
	}).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			formData.Errors["general"] = "An account with that username or e-mail already exists"
		} else {
			formData.Errors["general"] = "Something went wrong creating your account. If the problem persists, please contact us."
		}
		c.HTML(http.StatusUnprocessableEntity, "users/signup_form", formData)
		return
	}

	c.Header("HX-Redirect", "/")
}
