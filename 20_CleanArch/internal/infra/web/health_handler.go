package web

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/streadway/amqp"
)

type HealthHandler struct {
	DB              *sql.DB
	RabbitMQChannel *amqp.Channel
}

type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

func NewHealthHandler(db *sql.DB, rabbitMQChannel *amqp.Channel) *HealthHandler {
	return &HealthHandler{
		DB:              db,
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	services := make(map[string]string)
	overallStatus := "healthy"

	// Check Database
	if err := h.DB.Ping(); err != nil {
		services["database"] = "unhealthy"
		overallStatus = "unhealthy"
	} else {
		services["database"] = "healthy"
	}

	// Check RabbitMQ
	if h.RabbitMQChannel == nil {
		services["rabbitmq"] = "unhealthy"
		overallStatus = "unhealthy"
	} else {
		// Tenta declarar uma exchange temporária para verificar conexão
		err := h.RabbitMQChannel.ExchangeDeclarePassive(
			"amq.direct",
			"direct",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			services["rabbitmq"] = "unhealthy"
			overallStatus = "unhealthy"
		} else {
			services["rabbitmq"] = "healthy"
		}
	}

	response := HealthResponse{
		Status:   overallStatus,
		Services: services,
	}

	w.Header().Set("Content-Type", "application/json")

	if overallStatus == "unhealthy" {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
