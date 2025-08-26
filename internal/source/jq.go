package source

import (
	"bytes"
	"log"
	"os/exec"
)

type jqFileConfig struct {
	Path     string
	JQFilter string
}

var allScryfallCards jqFileConfig = jqFileConfig{
	Path:     "./src/all-cards.json",
	JQFilter: "map(select((any(.games[]; . == \"paper\")) and (.lang == \"en\" or .lang == \"es\")))",
}

func RunJQCmd(jqFile jqFileConfig) ([]byte, error) {
	jqCmd := exec.Command("jq", jqFile.JQFilter, jqFile.Path)
	jqOutBuf := bytes.Buffer{}
	jqErrBuf := bytes.Buffer{}
	jqCmd.Stdout = &jqOutBuf
	jqCmd.Stderr = &jqErrBuf
	err := jqCmd.Run()

	if err != nil {
		log.Println(jqOutBuf.String())
		log.Println(jqErrBuf.String())
		log.Println(err)
		return []byte{}, err
	}

	return jqOutBuf.Bytes(), nil
}
