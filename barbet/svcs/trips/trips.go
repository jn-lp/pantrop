package trips

import (
  "context"

  "github.com/jn-lp/pantrop/barbet"
)

type Service interface {
  ListTrips(ctx context.Context) (*[]barbet.Trip, error)
  CreateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error)
  GetTrip(ctx context.Context, tripID string) (*barbet.Trip, error)
  UpdateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error)
  DeleteTrip(ctx context.Context, tripID string) error
}
