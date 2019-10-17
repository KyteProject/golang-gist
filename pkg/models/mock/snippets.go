package mock

import (
	"time"

	"github.com/kyteproject/golang-gist/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

// SnippetModel is a mock Snippet
type SnippetModel struct{}

// Insert inserts a mock snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 2, nil
}

// Get retreives a mock snippet
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

// Latest retreives latest mock snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
