package approval

import "sync"

type Status string

const (
	Pending Status = "pending"
	Approved Status = "approved"
	Denied Status = "denied"
)

type Request struct {
	Tool string
	Summary string
	Risk string
	Status Status
}

type Manager struct {
	mu sync.RWMutex
	pending map[string]Request
}

func New() *Manager { return &Manager{pending: make(map[string]Request)} }
func (m *Manager) Add(id string, r Request) { m.mu.Lock(); defer m.mu.Unlock(); if r.Status == "" { r.Status = Pending }; m.pending[id] = r }
func (m *Manager) Get(id string) (Request, bool) { m.mu.RLock(); defer m.mu.RUnlock(); r, ok := m.pending[id]; return r, ok }
func (m *Manager) Approve(id string) bool { return m.resolve(id, Approved) }
func (m *Manager) Deny(id string) bool { return m.resolve(id, Denied) }
func (m *Manager) resolve(id string, status Status) bool { m.mu.Lock(); defer m.mu.Unlock(); r, ok := m.pending[id]; if !ok { return false }; r.Status = status; m.pending[id] = r; return true }
func (m *Manager) ConsumeApproved(id string) bool { m.mu.Lock(); defer m.mu.Unlock(); r, ok := m.pending[id]; if !ok || r.Status != Approved { return false }; delete(m.pending, id); return true }
