package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/dto"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

// GetJWT godoc
// @Summary Get a JWT
// @Description Get a JWT
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.GetJWTInput true "user credentials"
// @Success 200 {object} dto.GetJWTOutput
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var user dto.GetJWTInput

	// Recebe o usuário
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	// Procura o usuário
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	// Valida a senha
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	// Gera o token
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

// CreateUser user godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201 {string} string "created"
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
