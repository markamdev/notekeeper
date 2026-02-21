package server

import (
	"context"
	"fmt"

	"github.com/markamdev/notekeeper/proto/notekeeper"
)

func NewGRPCServer() notekeeper.NoteKeeperServer {
	return &NoteKeeperServer{}
}

type NoteKeeperServer struct {
	notekeeper.UnimplementedNoteKeeperServer
}

func (nks *NoteKeeperServer) AddNote(_ context.Context, note *notekeeper.NewNote) (*notekeeper.AddResult, error) {
	// TODO implement note adding to DB
	fmt.Println("adding new note", note.Title)
	return &notekeeper.AddResult{
		Status: notekeeper.Status_STATUS_OK,
		Id:     "123",
	}, nil
}

func (nks *NoteKeeperServer) GetNotes(context.Context, *notekeeper.EmptyMsg) (*notekeeper.NotesCollection, error) {
	// TODO implement note fetching from DB
	fmt.Println("fetching notes from DB")
	return &notekeeper.NotesCollection{
		Note: []*notekeeper.NoteContent{
			{Id: "123", Title: "Fake note 1", Content: "Fake content 123"},
			{Id: "456", Title: "Fake note 2", Content: "Fake content 456"},
		},
	}, nil
}
