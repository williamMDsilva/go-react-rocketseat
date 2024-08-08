package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/williamMDsilva/str-go-back-end/internal/store/pgstore"
)

func (h apiHandler) handleGetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.q.GetRooms(r.Context())
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		slog.Error("failed to get rooms", "error", err)
		return
	}

	if rooms == nil {
		rooms = []pgstore.Room{}
	}

	data, _ := json.Marshal(rooms)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(data)
}
