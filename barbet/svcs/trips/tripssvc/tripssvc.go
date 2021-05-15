package tripssvc

import (
  "context"

  "github.com/jn-lp/pantrop/barbet"
  "github.com/jn-lp/pantrop/barbet/svcs/trips"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripsrepo"
)

type service struct {
  repository tripsrepo.Repository
}

func New(r tripsrepo.Repository) trips.Service {
  return &service{
    repository: r,
  }
}

func (s *service) ListTrips(ctx context.Context) (*[]barbet.Trip, error) {
  return s.repository.ListTrips(ctx)
}

func (s *service) CreateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error) {
  return s.repository.CreateTrip(ctx, trip)
}

func (s *service) GetTrip(ctx context.Context, tripID string) (*barbet.Trip, error) {
  return s.repository.GetTrip(ctx, tripID)
}

func (s *service) UpdateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error) {
  return s.repository.UpdateTrip(ctx, trip)
}

func (s *service) DeleteTrip(ctx context.Context, tripID string) error {
  return s.repository.DeleteTrip(ctx, tripID)
}
