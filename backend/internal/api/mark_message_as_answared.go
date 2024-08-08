package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/williamMDsilva/str-go-back-end/internal/constants"
	typesApi "github.com/williamMDsilva/str-go-back-end/internal/types"
)

func (h apiHandler) handleMarkMessageAsAnswered(w http.ResponseWriter, r *http.Request) {
	rawRoomID := chi.URLParam(r, "room_id")
	_, err := uuid.Parse(rawRoomID)
	if err != nil {
		http.Error(w, "Invalid room id", http.StatusBadGateway)

	}

	rawID := chi.URLParam(r, "message_id")
	id, err := uuid.Parse(rawID)
	if err != nil {
		http.Error(w, "invalid message id", http.StatusBadRequest)
		return
	}

	err = h.q.MarkMessageAsAnswered(r.Context(), id)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		slog.Error("failed to react to message", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)

	go h.notifyClients(typesApi.Message{
		Kind:   constants.MessageKindMessageAnswered,
		RoomID: rawRoomID,
		Value: typesApi.MessageMessageAnswered{
			ID: rawID,
		},
	})
}
