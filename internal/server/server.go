package server

import (
	"log"
	"net/http"

	"github.com/Fryuni/world-anvil-cli/pkg/utils"
)

type Handler struct {
	// Add dependencies here (database, services, etc.)
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", h.HealthCheck)
	mux.HandleFunc("/api/users", h.GetUsers)
	mux.HandleFunc("/api/worlds", h.GetWorlds)

	return mux
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// TODO: Implement user retrieval logic
	utils.RespondJSON(w, http.StatusOK, []map[string]string{
		{"id": "1", "username": "example", "email": "example@example.com"},
	})
}

func (h *Handler) GetWorlds(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// TODO: Implement worlds retrieval logic
	utils.RespondJSON(w, http.StatusOK, []map[string]string{
		{"id": "1", "name": "Example World", "description": "An example world"},
	})
}

func (h *Handler) Start(port string) error {
	mux := h.SetupRoutes()
	log.Printf("Server starting on port %s", port)
	return http.ListenAndServe(":"+port, mux)
}
