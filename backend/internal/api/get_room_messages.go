package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/williamMDsilva/str-go-back-end/internal/store/pgstore"
)

func (h apiHandler) handleGetRoomMessages(w http.ResponseWriter, r *http.Request) {
	rawRoomID := chi.URLParam(r, "room_id")
	roomID, err := uuid.Parse(rawRoomID)
	if err != nil {
		http.Error(w, "Invalid room id", http.StatusBadGateway)

	}

	messages, err := h.q.GetRoomMessage(r.Context(), roomID)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		slog.Error("failed to get room messages", "error", err)
		return
	}

	if messages == nil {
		messages = []pgstore.Message{}
	}

	// TODO - handle errors in the feature
	data, _ := json.Marshal(messages)
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(data)
}
