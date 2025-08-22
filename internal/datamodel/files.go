package datamodel

// file: AtomicCards.json
type AllAtomicCards struct {
	Meta Meta                    `json:"meta"`
	Data map[string][]AtomicCard `json:"data"`
}

// file: SetList.json
type AllSets struct {
	Meta Meta  `json:"meta"`
	Data []Set `json:"data"`
}

// file: AllIdentifiers.json
type AllSetCards struct {
	Meta Meta               `json:"meta"`
	Data map[string]SetCard `json:"data"`
}
