package models

import (
	"errors"
	"strings"
	"time"
)

// Publication represents a post made by a user
type PublicationModel struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	UpdateAt   time.Time `json:"createdAt,omitempty"`
}

// Prepare will invoke methods to validate and format the received publication
func (publication *PublicationModel) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

func (publication *PublicationModel) validate() error {
	if publication.Title == "" {
		return errors.New("title is mandatory and cannot be blank")
	}

	if publication.Content == "" {
		return errors.New("content is mandatory and cannot be blank")
	}

	return nil
}

func (publication *PublicationModel) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
