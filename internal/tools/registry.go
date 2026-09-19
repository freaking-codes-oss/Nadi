package tools

type Definition struct {
	Name string
	Description string
	InputSchema map[string]any
}

type Registry struct {
	items map[string]Definition
	implementations map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{items: make(map[string]Definition), implementations: make(map[string]Tool)}
}
func (r *Registry) Register(d Definition) { r.items[d.Name] = d }
func (r *Registry) RegisterTool(t Tool) { if t == nil { return }; d := t.Definition(); r.Register(d); r.implementations[d.Name] = t }
func (r *Registry) Get(name string) (Definition, bool) { d, ok := r.items[name]; return d, ok }
func (r *Registry) Implementation(name string) (Tool, bool) { t, ok := r.implementations[name]; return t, ok }
func (r *Registry) List() []Definition { out := make([]Definition, 0, len(r.items)); for _, d := range r.items { out = append(out, d) }; return out }
