package models

import "time"

// Note represents a note created by a User
type Note struct {
	// ID populated while persisted in the db
	ID string `json:"_id,omitempty" bson:"_id"`

	// Kind determines the type of Note created. It can be checklist, bullet points etc.
	// This is enforced in on the Contents field
	Kind string `json:"kind" bson:"kind"`

	// CreatedAt when request is received by the server
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`

	// LastModified when request is received by the server
	LastModified time.Time `json:"last_modified,omitempty" bson:"last_modified"`

	// AuthorId is autofilled from jwt token
	AuthorId string `json:"author_id,omitempty" bson:"author_id"`

	// Title is provided in request body (required)
	Title string `json:"title,omitempty" bson:"title"`

	// Contents is provided in request body (required)
	Contents []NoteContent `json:"content,omitempty" bson:"content"`
}

// NoteContent is a struct that holds the content of a note
// The contents can be of different types (e.g. bullet points, checklist, etc.)
type NoteContent struct {
	Kind       string     `json:"kind" bson:"kind"`
	Text       string     `json:"text" bson:"text"`
	Decoration Decoration `json:"decoration" bson:"decoration"`
}

// Decoration is a struct that holds the decoration of a note (e.g. checkbox,
// cross line over text, etc.), this is used by the frontend to render the note
type Decoration struct {
	IsChecked    bool `json:"is_checked" bson:"is_checked"`         // for checkbox
	IsCrossLined bool `json:"is_cross_lined" bson:"is_cross_lined"` // for cross line over text
}

func IsValidCreateNote(note *Note) bool { return note.Title != "" }
