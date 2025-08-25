package source

type jqFileConfig struct {
	Path     string
	JQFilter string
}

var allSets jqFileConfig = jqFileConfig{
	Path:     "./src/SetList.json",
	JQFilter: "{meta, data: [.data | .[] | {code, name}]}",
}

var allAtomicCards jqFileConfig = jqFileConfig{
	Path:     "./src/AtomicCards.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | .[0] | {name, identifiers}]}",
}

var allSetCards jqFileConfig = jqFileConfig{
	Path:     "./src/AllIdentifiers.json",
	JQFilter: "{meta, data: [.data | to_entries | .[] | .value | {uuid, identifiers, name, setCode, number}]}",
}
