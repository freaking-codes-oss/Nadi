package session

import "sync"

type Manager struct { mu sync.RWMutex; values map[string]string }
func New() *Manager { return &Manager{values: map[string]string{}} }
func (m *Manager) Set(key, value string) { m.mu.Lock(); defer m.mu.Unlock(); m.values[key] = value }
func (m *Manager) Get(key string) (string, bool) { m.mu.RLock(); defer m.mu.RUnlock(); v, ok := m.values[key]; return v, ok }
