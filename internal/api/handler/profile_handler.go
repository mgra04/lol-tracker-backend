package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mgra04/backend-lol-v1/internal/service"
	"github.com/mgra04/backend-lol-v1/pkg/errors"
)

type ProfileHandler struct {
    service *service.ProfileService
}

func NewProfileHandler(s *service.ProfileService) *ProfileHandler {
    return &ProfileHandler{service: s}
}

func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
    gameName := r.URL.Query().Get("gameName")
    tagLine := r.URL.Query().Get("tagLine")

    if gameName == "" || tagLine == "" {
        http.Error(w, "Missing query parameters: gameName, tagLine", http.StatusBadRequest)
        return
    }

    profile, err := h.service.GetProfile(gameName, tagLine)
    if err != nil {
        if appErr, ok := err.(*errors.AppError); ok {
            http.Error(w, appErr.Message, appErr.Code)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(profile); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}
