package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rodolfole/go-users-api/config"
	"github.com/rodolfole/go-users-api/services/auth"
	"github.com/rodolfole/go-users-api/types"
	"github.com/rodolfole/go-users-api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("falta el campo: %v", errors.Error()))
		return
	}

	u, err := h.store.GetUserByCorreo(user.Correo)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("usuario / contraseña incorrectos"))
		return
	}

	if !auth.ComparePasswords(u.Contrasena, []byte(user.Contrasena)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("usuario / contraseña incorrectos"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err := h.store.CheckIfUserExist(user.Correo, user.Telefono)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("el correo/telefono ya se encuentra registrado"))
		return
	}

	isValid := auth.IsValidPassword(user.Contrasena)
	if !isValid {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("contrasena invalida debe tener al menos 1 minuscula, 1 mayuscula, 1 numero y un caracter como: @, $ o &"))
		return
	}

	if !utils.IsValidEmail(user.Correo) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("correo invalido"))
		return
	}

	if !utils.IsValidPhone(user.Telefono) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("telefono invalido"))
		return
	}

	hashedPassword, err := auth.HashPassword(user.Contrasena)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		Usuario:    user.Usuario,
		Correo:     user.Correo,
		Telefono:   user.Telefono,
		Contrasena: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
