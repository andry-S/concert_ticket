package session

import "sync"

var sessions = make(map[uint]bool)
var mu sync.Mutex

func SetLoggedIn(userID uint) {
	mu.Lock()
	defer mu.Unlock()
	sessions[userID] = true
}

func IsLoggedIn(userID uint) bool {
	mu.Lock()
	defer mu.Unlock()
	return sessions[userID]
}
