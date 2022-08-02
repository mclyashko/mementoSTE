package editor

import (
	"bufio"
	"fmt"
	"mementoSTE/stack"
	"mementoSTE/stringF"
	"os"
	"strconv"
	"strings"
)

const (
	startTag        = "BEGIN"
	finishTag       = "END"
	cmdWritePrefix  = ">>> "
	wrongCommand    = "no such command expression"
	exitCaseCmd     = "!exit"
	addRuneCaseCmd  = "!add "
	multiplyCaseCmd = "!multiply "
	printCaseCmd    = "!print"
	restoreCaseCmd  = "!restore"
)

type Controller struct {
	editor  *editor
	history *stack.Stack
}

func (ctrl *Controller) historySaver(cmd string) {
	if !(strings.HasPrefix(cmd, exitCaseCmd) || strings.HasPrefix(cmd, printCaseCmd) ||
		strings.HasPrefix(cmd, restoreCaseCmd)) {
		ctrl.history.Push(ctrl.editor.doSaveMemento())
	}
}

func (ctrl *Controller) commandHandler(cmd string) (exitState bool) {
	ctrl.historySaver(cmd)

	switch { // bad practices, but alright for small program like this
	case strings.HasPrefix(cmd, exitCaseCmd):
		exitState = true
	case strings.HasPrefix(cmd, addRuneCaseCmd):
		cmd = stringF.DeleteCompositePrefix(cmd, addRuneCaseCmd)
		if len(cmd) == 0 {
			exitState = true
			break
		}
		ctrl.editor.doAddRune([]rune(cmd)[0])
	case strings.HasPrefix(cmd, multiplyCaseCmd):
		cmd = stringF.DeleteCompositePrefix(cmd, multiplyCaseCmd)
		factorArg, err := strconv.Atoi(cmd)
		if err != nil {
			exitState = true
			break
		}
		ctrl.editor.doMultiply(factorArg)
	case strings.HasPrefix(cmd, printCaseCmd):
		ctrl.editor.doPrint()
	case strings.HasPrefix(cmd, restoreCaseCmd):
		if val := ctrl.history.Pop(); val != nil {
			ctrl.editor.doRestoreMemento(val.(*editorMemento))
		} else {
			exitState = true
			break
		}
	default:
		fmt.Println(wrongCommand)
	}

	return
}

func (ctrl *Controller) commandListener() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(cmdWritePrefix)

		sc.Scan()
		cmd := sc.Text()
		err := sc.Err()

		if err != nil {
			break
		}

		if ctrl.commandHandler(cmd) {
			break
		}
	}
}

func Starter() {
	fmt.Println(startTag)

	ctrl := &Controller{
		editor:  &editor{},
		history: &stack.Stack{},
	}

	ctrl.commandListener()

	fmt.Println(finishTag)
}
