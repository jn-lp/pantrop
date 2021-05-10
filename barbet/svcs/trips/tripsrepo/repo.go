package tripsrepo

import (
  "context"
  "fmt"

  _ "github.com/lib/pq"
  "gorm.io/gorm"

  "github.com/jn-lp/pantrop/barbet"
)

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
  ListTrips(ctx context.Context) (*[]barbet.Trip, error)
  CreateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error)
  GetTrip(ctx context.Context, tripID string) (*barbet.Trip, error)
  UpdateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error)
  DeleteTrip(ctx context.Context, tripID string) error
}

type repository struct {
  db *gorm.DB
}

func New(db *gorm.DB) Repository {
  return &repository{
    db: db,
  }
}

func (r repository) ListTrips(ctx context.Context) (*[]barbet.Trip, error) {
  var trips []barbet.Trip
  r.db.WithContext(ctx).Find(&trips)

  return &trips, nil
}

func (r repository) CreateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error) {
  r.db.WithContext(ctx).Create(trip)

  return trip, nil
}

func (r repository) GetTrip(ctx context.Context, tripID string) (*barbet.Trip, error) {
  var trip barbet.Trip
  r.db.WithContext(ctx).First(&trip, tripID)

  if trip.ID == 0 {
    return nil, fmt.Errorf("trip not found")
  }

  return &trip, nil
}

func (r repository) UpdateTrip(ctx context.Context, trip *barbet.Trip) (*barbet.Trip, error) {
  r.db.WithContext(ctx).Save(&trip)

  return trip, nil
}

func (r repository) DeleteTrip(ctx context.Context, tripID string) error {
  r.db.WithContext(ctx).Delete(tripID)

  return nil
}
