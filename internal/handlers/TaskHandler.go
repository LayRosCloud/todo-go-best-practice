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

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	Service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) FindAllByUserIdPagination(w http.ResponseWriter, r *http.Request) {
	limit := query.GetQueryInt(r, "limit", 10)
	page := query.GetQueryInt(r, "page", 1)
	userIdStr := chi.URLParam(r, "id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		exceptions.WriteError(w, exceptions.GetBadRequestError("UserId has bad request", err.Error()))
	}
	dto := dto.TaskFindAllRequest{UserId: userId, Limit: limit, Page: page}
	tasks, totalCount, err := h.Service.FindAllByUserIdPagination(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	w.Header().Set(handlersUtils.TotalCountHeader, string(totalCount))
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) FindById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		exceptions.WriteError(w, exceptions.GetBadRequestError("Id has bad format", err.Error()))
	}
	dto := dto.TaskByIdRequest{Id: id}
	task, err := h.Service.FindById(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.TaskCreateRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		exceptions.WriteError(w, exceptions.GetBadRequestError("JSON Invalid", err.Error()))
	}
	task, err := h.Service.Create(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	var dto dto.TaskUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		exceptions.WriteError(w, exceptions.GetBadRequestError("JSON Invalid", err.Error()))
	}
	task, err := h.Service.Update(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "id")
	taskId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		exceptions.WriteError(w, exceptions.GetBadRequestError("Id has bad format", err.Error()))
	}
	dto := dto.TaskByIdRequest{Id: taskId}
	result, err := h.Service.Delete(r.Context(), &dto)
	if err != nil {
		handlersUtils.HandleError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)
}