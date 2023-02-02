package src

import (
	"crypto/rand"
	"fmt"
	"sync"
)

func generateUuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	u := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	if !safeUsedUuidMapInstance.Get(u) {
		safeUsedUuidMapInstance.Add(u)
		return u, nil
	}
	return generateUuid()
}

type safeUsedUuidMap struct {
	sync.RWMutex
	usedUuid map[string]struct{}
}

var safeUsedUuidMapInstance safeUsedUuidMap

func init() {
	safeUsedUuidMapInstance = safeUsedUuidMap{usedUuid: map[string]struct{}{}}
}

func (s *safeUsedUuidMap) Add(uuid string) {
	s.Lock()
	defer s.Unlock()
	s.usedUuid[uuid] = struct{}{}
}

func (s *safeUsedUuidMap) Get(uuid string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.usedUuid[uuid]
	return ok
}
