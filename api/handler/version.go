package handler

import (
	"net/http"

	"github.com/karasunokami/otus-calendar/internal/version"
)

// VersionHandler Return current application version
// GET /
func (h *Handler) VersionHandler(_ *http.Request) APIResponse {
	return h.JSON(http.StatusOK, map[string]string{"version": version.Version})
}
