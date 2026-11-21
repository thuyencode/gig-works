package main

import (
	"context"
	"log"
	"log/slog"
	"os/signal"
	"syscall"

	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	inmemRepo := repository.NewInmemRepository()
	service := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{ID: primitive.NewObjectID(), UserID: "67", PackageSlug: ""}
	t, err := service.CreateTrip(ctx, fare)
	if err != nil {
		slog.Error("error creating trip", "error", err)
	}
	log.Println(t)

	<-ctx.Done()
}
