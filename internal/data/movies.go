package data

import (
	"encoding/json"
	"github.com/basotyev/greenlight/internal/validator"
	"time"
)

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   Runtime
	Genres    []string
	Version   int32
}

// Implement a MarshalJSON() method on the Movie struct, so that it satisfies the
// json.Marshaler interface.
func (m Movie) MarshalJSON() ([]byte, error) {
	// Declare a variable to hold the custom runtime string (this will be the empty
	// string "" by default).
	var runtime string
	// If the value of the Runtime field is not zero, set the runtime variable to be a
	// string in the format "<runtime> mins".
	// Create an anonymous struct to hold the data for JSON encoding. This has exactly
	// the same fields, types and tags as our Movie struct, except that the Runtime
	// field here is a string, instead of an int32. Also notice that we don't include
	// a CreatedAt field at all (there's no point including one, because we don't want
	// it to appear in the JSON output).
	aux := struct {
		ID      int64    `json:"id"`
		Title   string   `json:"title"`
		Year    int32    `json:"year,omitempty"`
		Runtime string   `json:"runtime,omitempty"` // This is a string.
		Genres  []string `json:"genres,omitempty"`
		Version int32    `json:"version"`
	}{
		// Set the values for the anonymous struct.
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime, // Note that we assign the value from the runtime variable here.
		Genres:  m.Genres,
		Version: m.Version,
	}
	// Encode the anonymous struct to JSON, and return it.
	return json.Marshal(aux)
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
