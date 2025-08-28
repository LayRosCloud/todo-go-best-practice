package handlers

import (
	"encoding/json"
	"leafall/todo-service/internal/dto"
	handlersUtils "leafall/todo-service/internal/handlers/utils"
	"leafall/todo-service/internal/services"
	"leafall/todo-service/utils/exceptions"
	"leafall/todo-service/utils/query"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (u UserHandler) FindAllPagination(w http.ResponseWriter, r *http.Request) {
	limit := query.GetQueryInt(r, "limit", 10)
	page := query.GetQueryInt(r, "page", 1)
	dto := dto.UserFindAllRequest{Limit: limit, Page: page}
	users, totalCount, err := u.Service.FindAllPagination(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	w.Header().Set("X-Total-Count", strconv.FormatInt(totalCount, 10))
	json.NewEncoder(w).Encode(users)
}

func (u UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	limit := query.GetQueryInt(r, "limit", 10)
	page := query.GetQueryInt(r, "page", 1)
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        handlersUtils.HandleError(
			exceptions.NewBadRequestSimple("Invalid user id", err.Error()), w)
        return
    }
	dto := dto.UserFindByIdRequest{Id: id, Limit: limit, Page: page}
	user, totalCount, err := u.Service.FindById(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	w.Header().Set("X-Total-Count", string(totalCount))
	json.NewEncoder(w).Encode(user)
}

func (u UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handlersUtils.HandleError(
			exceptions.NewBadRequestSimple("Invalid JSON", err.Error()), w)
		return
	}

	user, err := u.Service.Create(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handlersUtils.HandleError(
			exceptions.NewBadRequestSimple("Invalid JSON", err.Error()), w)
		return
	}

	user, err := u.Service.Update(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserUpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		handlersUtils.HandleError(
			exceptions.NewBadRequestSimple("Invalid JSON", err.Error()), w)
		return
	}

	user, err := u.Service.UpdatePassword(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (u UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
		handlersUtils.HandleError(
			exceptions.NewBadRequestSimple("Invalid JSON", err.Error()), w)
        return
    }
	dto := dto.UserDeleteByIdRequest{Id: id}
	result, err := u.Service.Delete(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}