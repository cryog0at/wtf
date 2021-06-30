package textfile

import (
	"github.com/cryog0at/wtf/utils"
	"github.com/gdamore/tcell"
)

func (widget *Widget) initializeKeyboardControls() {
	widget.InitializeHelpTextKeyboardControl(widget.ShowHelp)
	widget.InitializeRefreshKeyboardControl(nil)

	widget.SetKeyboardChar("l", widget.NextSource, "Select next file")
	widget.SetKeyboardChar("h", widget.PrevSource, "Select previous file")
	widget.SetKeyboardChar("o", widget.openFile, "Open file")

	widget.SetKeyboardKey(tcell.KeyRight, widget.NextSource, "Select next file")
	widget.SetKeyboardKey(tcell.KeyLeft, widget.PrevSource, "Select previous file")
	widget.SetKeyboardKey(tcell.KeyEnter, widget.openFile, "Open file")
}

func (widget *Widget) openFile() {
	src := widget.CurrentSource()
	utils.OpenFileWithEditor(src)
}
