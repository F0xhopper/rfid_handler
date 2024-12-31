package state

import (
	"sync"
	"time"
)


type Item struct {
	ID       string
	Collected bool
	CollectionDate time.Time
}


type AppState struct {
	Items []Item
	Mutex sync.Mutex
}


func NewAppState() *AppState {
	return &AppState{
		Items: []Item{	
		   	{ID: "a3b2d1", Collected: false},
			{ID: "d6e4f7", Collected: false},
			{ID: "z1x3c8", Collected: false},
			{ID: "j5k7l2", Collected: false},
			{ID: "m9n0o4", Collected: false},
			{ID: "p8q3r1", Collected: false},
			{ID: "u5v2w9", Collected: false},
		},
	}
}