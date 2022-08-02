package editor

type editorMemento struct {
	mementoState string
}

func (memento *editorMemento) getMementoState() string {
	return memento.mementoState
}
