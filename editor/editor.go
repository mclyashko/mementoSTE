package editor

import (
	"fmt"
	"strings"
)

type editor struct {
	editorState string
}

func (ed *editor) doAddRune(suffix rune) {
	ed.editorState += string(suffix)
}

func (ed *editor) doMultiply(factor int) {
	builder := &strings.Builder{}

	for i := 0; i < factor; i++ {
		builder.WriteString(ed.editorState)
	}

	ed.editorState = builder.String()
}

func (ed *editor) doPrint() {
	fmt.Println(ed.editorState)
}

func (ed *editor) doSaveMemento() *editorMemento {
	return &editorMemento{
		mementoState: ed.editorState,
	}
}

func (ed *editor) doRestoreMemento(memento *editorMemento) {
	ed.editorState = memento.getMementoState()
}
