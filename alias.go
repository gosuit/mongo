package mongo

import "go.mongodb.org/mongo-driver/v2/mongo"

var (
	// Common errors returned by the mongo package.

	ErrClientDisconnected     = mongo.ErrClientDisconnected
	ErrEmptySlice             = mongo.ErrEmptySlice
	ErrFileNotFound           = mongo.ErrFileNotFound
	ErrInvalidIndexValue      = mongo.ErrInvalidIndexValue
	ErrMissingChunk           = mongo.ErrMissingChunk
	ErrMissingGridFSChunkSize = mongo.ErrMissingGridFSChunkSize
	ErrMissingResumeToken     = mongo.ErrMissingResumeToken
	ErrMultipleIndexDrop      = mongo.ErrMultipleIndexDrop
	ErrNilCursor              = mongo.ErrNilCursor
	ErrNilDocument            = mongo.ErrNilDocument
	ErrNilValue               = mongo.ErrNilValue
	ErrNoDocuments            = mongo.ErrNoDocuments
	ErrNonStringIndexName     = mongo.ErrNonStringIndexName
	ErrNotSlice               = mongo.ErrNotSlice
	ErrStreamClosed           = mongo.ErrStreamClosed
	ErrWrongClient            = mongo.ErrWrongClient
	ErrWrongSize              = mongo.ErrWrongSize
)
