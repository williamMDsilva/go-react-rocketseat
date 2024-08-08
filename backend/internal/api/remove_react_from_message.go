package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/williamMDsilva/str-go-back-end/internal/constants"
	typesApi "github.com/williamMDsilva/str-go-back-end/internal/types"
)

func (h apiHandler) handleRemoveReactFromMessage(w http.ResponseWriter, r *http.Request) {
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

	count, err := h.q.RemoveReactToMessage(r.Context(), id)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		slog.Error("failed to react to message", "error", err)
		return
	}

	type response struct {
		Count int64 `json:"count"`
	}

	// TODO - handle errors in the feature
	data, _ := json.Marshal(response{Count: count})
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(data)

	go h.notifyClients(typesApi.Message{
		Kind:   constants.MessageKindMessageRactionIncreased,
		RoomID: rawRoomID,
		Value: typesApi.MessageMessageReactionDecreased{
			ID:    rawID,
			Count: count,
		},
	})
}
