package api

import (
	"log/slog"

	typesApi "github.com/williamMDsilva/str-go-back-end/internal/types"
)

func (h apiHandler) notifyClients(msg typesApi.Message) {
	h.mu.Lock()
	defer h.mu.Unlock()

	subscribers, ok := h.subscribers[msg.RoomID]
	if !ok || len(subscribers) == 0 {
		return
	}

	for conn, cancel := range subscribers {
		if err := conn.WriteJSON(msg); err != nil {
			slog.Error("failed to send message to client", "error", err)
			cancel()
		}
	}
}
