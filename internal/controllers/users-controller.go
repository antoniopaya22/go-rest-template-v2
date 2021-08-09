package controllers

import (
	"ai4am.com/go-restapi/internal/repository"
	"ai4am.com/go-restapi/models"
	"ai4am.com/go-restapi/pkg/crypto"
	"ai4am.com/go-restapi/pkg/helpers"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInput struct {
	Username  string `json:"username" binding:"required"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}


// GetUserById godoc
// @Tags Users
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /api/users/{id} [get]
func GetUserById(c *gin.Context) {
	s := repository.GetUserRepository()
	id := c.Param("id")
	if user, err := s.GetUserById(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers godoc
// @Tags Users
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []models.User
// @Router /api/users [get]
func GetAllUsers(c *gin.Context) {
	s := repository.GetUserRepository()
	var q models.User
	_ = c.Bind(&q)
	if users, err := s.QueryUsers(&q); err != nil {
		c.JSON(http.StatusOK, []models.User{})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// GetUsers godoc
// @Tags Users
// @Summary Creates a new User
// @Description Add User
// @Accept  json
// @Produce json
// @Param user body models.User true "Add user"
// @Success 201 {object} UserInput
// @Router /api/users [post]
// @Security Authorization Token
func CreateUser(c *gin.Context) {
	s := repository.GetUserRepository()
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		Role:      userInput.Role,
	}
	if err := s.AddUser(&user); err != nil {
		helpers.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

// GetUsers godoc
// @Tags Users
// @Summary Updates an User by ID
// @Description Update User
// @Accept  json
// @Produce json
// @Param id path integer true "User ID"
// @Param user body models.User true "Add user"
// @Success 200 {object} UserInput
// @Router /api/users/{id} [put]
// @Security Authorization Token
func UpdateUser(c *gin.Context) {
	s := repository.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.GetUserById(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		user.Username = userInput.Username
		user.Lastname = userInput.Lastname
		user.Firstname = userInput.Firstname
		user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
		user.Role = userInput.Role
		if err := s.UpdateUser(user); err != nil {
			helpers.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

// GetUserById godoc
// @Tags Users
// @Summary Deletes user based on given ID
// @Description Delete user
// @Produce json
// @Param id path integer true "User ID"
// @Success 204
// @Router /api/users/{id} [get]
// @Security Authorization Token
func DeleteUser(c *gin.Context) {
	s := repository.GetUserRepository()
	id := c.Params.ByName("id")
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	if user, err := s.GetUserById(id); err != nil {
		helpers.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if err := s.DeleteUser(user); err != nil {
			helpers.NewError(c, http.StatusNotFound, err)
			log.Println(err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
