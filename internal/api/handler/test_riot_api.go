package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mgra04/backend-lol-v1/internal/constants"
	"github.com/mgra04/backend-lol-v1/internal/service"
	"github.com/rs/zerolog/log"
)

type TestRiotApiHandler struct {
	service *service.TestRiotApiService
}

func NewTestRiotApiHandler(s *service.TestRiotApiService) *TestRiotApiHandler {
	return &TestRiotApiHandler {
		service: s,
	}
}

func (h *TestRiotApiHandler) Test(w http.ResponseWriter, r *http.Request) {
	gameName := r.URL.Query().Get("gameName")
	tagLine := r.URL.Query().Get("tagLine")
	regionStr := r.URL.Query().Get("region")

	if gameName == "" || tagLine == "" || regionStr == "" {
		http.Error(w, "missing required query parameters: gameName, tagLine, region", http.StatusBadRequest)
		return
	}

	region := constants.Region(regionStr)

	result, err := h.service.LookupProfile(gameName, tagLine, region)
	if err != nil {
		log.Error().Err(err).Msg("failed to lookup profile")
		http.Error(w, "failed to fetch profile from Riot API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}