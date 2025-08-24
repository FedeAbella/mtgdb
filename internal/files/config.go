package files

type fileConfig struct {
	Path     string
	JQFilter string
}

var allSets fileConfig = fileConfig{
	Path:     "./src/SetList.json",
	JQFilter: "{meta, data: [.data | .[] | {code, name}]}",
}

var allAtomicCards fileConfig = fileConfig{
	Path:     "./src/AtomicCards.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | .[0] | {name, identifiers}]}",
}

var allSetCards fileConfig = fileConfig{
	Path:     "./src/AllIdentifiers.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | {uuid, identifiers, name, setCode, number}]}",
}
