package httptransport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	entity "github.com/TalantedMonkey21/GoLectures/internal/entity"
	"github.com/TalantedMonkey21/GoLectures/internal/response"
)

type NoteUseCase interface {
	Create(ctx context.Context, content string) (entity.Note, error)
	GetByID(ctx context.Context, id uint, userId uint) (entity.Note, error)
	Update(ctx context.Context, note entity.Note) (entity.Note, error)
	Delete(ctx context.Context, id uint, userId uint) error
}

type Handler struct {
	notes NoteUseCase
}

type createNoteRequest struct {
	Content string `json:"content"`
}

type updateNoteRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

type noteResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewHandler(notes NoteUseCase) *Handler {
	return &Handler{notes: notes}
}

func (h *Handler) Health(w http.ResponseWriter, _ *http.Request) {
	response.WriteJSONResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	note, err := h.notes.Create(r.Context(), req.Content)
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusCreated, toNoteResponse(note))
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	// path := strings.Split(r.URL.Path, "/")
	// pathId, err := strconv.Atoi(path[len(path)-1])
	pathId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	pathUserId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	note, err := h.notes.GetByID(r.Context(), uint(pathId), uint(pathUserId))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, toNoteResponse(note))
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var req updateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	note, err := h.notes.Update(r.Context(), entity.Note{ID: req.ID, Content: req.Content})
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, toNoteResponse(note))
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	pathId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	pathUserId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "Invalid user_id")
		return
	}

	err = h.notes.Delete(r.Context(), uint(pathId), uint(pathUserId))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, "Deleted")
}


func toNoteResponse(note entity.Note) noteResponse {
	return noteResponse{
		ID:        note.ID,
		Content:   note.Content,
		CreatedAt: note.CreatedAt.Format(http.TimeFormat),
		UpdatedAt: note.UpdatedAt.Format(http.TimeFormat),
	}
}