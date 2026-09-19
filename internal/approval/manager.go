package approval

import "sync"

type Request struct { Tool string; Summary string; Risk string }
type Manager struct { mu sync.Mutex; pending map[string]Request }
func New() *Manager { return &Manager{pending: make(map[string]Request)} }
func (m *Manager) Add(id string, r Request) { m.mu.Lock(); defer m.mu.Unlock(); m.pending[id] = r }
func (m *Manager) Get(id string) (Request, bool) { m.mu.Lock(); defer m.mu.Unlock(); r, ok := m.pending[id]; return r, ok }
func (m *Manager) Resolve(id string, approved bool) { m.mu.Lock(); defer m.mu.Unlock(); delete(m.pending, id) }
