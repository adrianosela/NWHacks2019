package prescriptions

// Medicine represents a particular medicine
type Medicine struct {
	Name        string   `json:"name"`
	ID          string   `json:"id"`
	Type        string   `json:"type"`
	Appereance  []string `json:"appereance"`
	SideEffects []string `json:"side_effects"`
}
