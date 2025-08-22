package datamodel

import "github.com/google/uuid"

type AtomicCard struct {
	Name        string      `json:"name"`
	Identifiers Identifiers `json:"identifiers"`
}

type SetCard struct {
	Id          uuid.UUID   `json:"uuid"`
	Identifiers Identifiers `json:"identifiers"`
	Name        string      `json:"name"`
	Set         string      `json:"setCode"`
	NumberInSet string      `json:"number"`
}
