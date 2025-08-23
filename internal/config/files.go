package config

type fileConfig struct {
	Path     string
	JQFilter string
}

var AllSets fileConfig = fileConfig{
	Path:     "./src/SetList.json",
	JQFilter: "{meta, data: [.data | .[] | {code, name}]}",
}

var AllAtomicCards fileConfig = fileConfig{
	Path:     "./src/AtomicCards.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | .[0] | {name, identifiers}]}",
}

var AllSetCards fileConfig = fileConfig{
	Path:     "./src/AllIdentifiers.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | {uuid, identifiers, name, setCode, number}]}",
}
