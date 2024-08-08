package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/williamMDsilva/str-go-back-end/internal/constants"
	"github.com/williamMDsilva/str-go-back-end/internal/store/pgstore"
	typesApi "github.com/williamMDsilva/str-go-back-end/internal/types"
)

func (h apiHandler) handleCreateRoomMessage(w http.ResponseWriter, r *http.Request) {
	rawRoomID := chi.URLParam(r, "room_id")
	roomID, err := uuid.Parse(rawRoomID)
	if err != nil {
		http.Error(w, "Invalid room id", http.StatusBadGateway)

	}
	type _body struct {
		Message string `json:"message"`
	}
	var body _body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	messageID, err := h.q.InsertMessage(r.Context(), pgstore.InsertMessageParams{RoomID: roomID, Message: body.Message})
	if err != nil {
		slog.Error("failed to insert message", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}

	// TODO - handle errors in the feature
	data, _ := json.Marshal(response{ID: roomID.String()})
	w.Header().Set("content-type", "application/json")
	_, _ = w.Write(data)

	go h.notifyClients(typesApi.Message{
		Kind:   constants.MessageKindMessageCreated,
		RoomID: rawRoomID,
		Value: typesApi.MessageMessageCreated{
			ID:      messageID.String(),
			Message: body.Message,
		},
	})
}
