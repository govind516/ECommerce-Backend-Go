package user

import (
	"fmt"
	"net/http"

	"github.com/Govind516/E-Commerce-Backend/service/auth"
	"github.com/Govind516/E-Commerce-Backend/types"
	"github.com/Govind516/E-Commerce-Backend/utils"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler{
	return &Handler{store:  store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)  {
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request)  {
	// get JSON payload 
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
	}

	// if it doesn't, then create new user
	hashedPassword, err := auth.HashedPassword(payload.Password)
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})
	
	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}