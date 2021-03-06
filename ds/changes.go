package ds

import (
	"context"
	"time"
)

// Change is used to track changes at a path.
// If this format changes, you should also change in firestore and other
// backends that don't directly use this struct on set.
type Change struct {
	Path      string    `json:"path" firestore:"path"`
	Timestamp time.Time `json:"ts" firestore:"ts"`
}

// Direction is ascending or descending.
type Direction string

const (
	// Ascending direction.
	Ascending Direction = "asc"
	// Descending direction.
	Descending Direction = "desc"
)

// Changes describes changes to a path.
type Changes interface {
	ChangeAdd(ctx context.Context, name string, id string, ref string) error
	Changes(ctx context.Context, name string, from time.Time, limit int, direction Direction) ([]*Change, time.Time, error)
}
