package controllers

import (
	"sync"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	userscontroller "truenorth/services/users_service/controllers/users"
)

type Server struct {
	mu sync.RWMutex
	// If shutdown is true, it's expected all serving status is NOT_SERVING, and
	// will stay in NOT_SERVING.
	shutdown bool
	// statusMap stores the serving status of the services this Server monitors.
	statusMap map[string]healthgrpc.HealthCheckResponse_ServingStatus
	updates   map[string]map[healthgrpc.Health_WatchServer]chan healthgrpc.HealthCheckResponse_ServingStatus

	healthgrpc.UnimplementedHealthServer
	*userscontroller.UsersControllerImpl
}

func NewServer() *Server {
	return &Server{

		statusMap:           map[string]healthgrpc.HealthCheckResponse_ServingStatus{"": healthgrpc.HealthCheckResponse_SERVING},
		updates:             make(map[string]map[healthgrpc.Health_WatchServer]chan healthgrpc.HealthCheckResponse_ServingStatus),
		UsersControllerImpl: userscontroller.NewUsersController(),
	}
}
