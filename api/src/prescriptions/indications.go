package prescriptions

// Indications represents the frequency and way of ingesting the meds
type Indications struct {
	DaysPerWeek int   `json:"dpw,omitempty"`
	TimesPerDay int   `json:"tpd,omitempty"`
	Time        []int `json:"tod,omitempty"` // HOUR OF DAY FMT E.G. 1100, 2359, etc
}
