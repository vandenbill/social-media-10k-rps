package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vandenbill/social-media-10k-rps/internal/dto"
	"github.com/vandenbill/social-media-10k-rps/internal/ierr"
	"github.com/vandenbill/social-media-10k-rps/internal/service"
	response "github.com/vandenbill/social-media-10k-rps/pkg/resp"
)

type userHandler struct {
	userSvc *service.UserService
}

func newUserHandler(userSvc *service.UserService) *userHandler {
	return &userHandler{userSvc}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.ReqRegister

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	res, err := h.userSvc.Register(r.Context(), req)
	if err != nil {
		code, msg := ierr.TranslateError(err)
		http.Error(w, msg, code)
		return
	}

	successRes := response.SuccessReponse{}
	successRes.Message = "User registered successfully"
	successRes.Data = res

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(successRes)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.ReqLogin

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to parse request body", http.StatusBadRequest)
		return
	}

	res, err := h.userSvc.Login(r.Context(), req)
	if err != nil {
		code, msg := ierr.TranslateError(err)
		http.Error(w, msg, code)
		return
	}

	successRes := response.SuccessReponse{}
	successRes.Message = "User logged successfully"
	successRes.Data = res

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(successRes)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}