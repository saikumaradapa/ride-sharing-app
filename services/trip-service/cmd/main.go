package main

import (
	"context"
	"log"
	"saikumaradapa/ride-sharing/services/trip-service/internal/domain"
	"saikumaradapa/ride-sharing/services/trip-service/internal/infrastructure/repository"
	"saikumaradapa/ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()

	svc := service.NewService(inmemRepo)

	fare := &domain.RideFareModel{
		UserID:          "42",
		PackageSlug:     "van",
		TotalPriceCents: 200,
	}

	t, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	for {
		time.Sleep(time.Second)
	}
}
