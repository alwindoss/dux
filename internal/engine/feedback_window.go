package engine

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	widgetx "fyne.io/x/fyne/widget"
	"github.com/alwindoss/dux/internal/feedback"
	"github.com/alwindoss/dux/internal/shortcut"
	"github.com/alwindoss/dux/internal/user"
)

func ShowNewFeedbackWindow(a fyne.App, fmgr *feedback.Manager, umgr *user.Manager) {
	w := createNewFeedbackWindow(a, fmgr, umgr)
	w.Canvas().AddShortcut(shortcut.CtrlQ, func(shortcut fyne.Shortcut) {
		w.Close()
	})
	w.Show()
}

func ShowAddActionWindow(a fyne.App) {
	w := a.NewWindow("New Action")
	w.Resize(fyne.NewSquareSize(600))
	w.Show()
}

var userNameList = []string{}

func queryUserName(wid *widgetx.CompletionEntry) func(string) {
	return func(s string) {
		wid.SetOptions([]string{time.Now().String()})
	}
}

func createNewFeedbackWindow(a fyne.App, fmgr *feedback.Manager, umgr *user.Manager) fyne.Window {
	w := a.NewWindow("New Feedback")
	w.Resize(fyne.NewSquareSize(600))

	nameSearchTextInput := widgetx.NewCompletionEntry(userNameList)
	nameSearchTextInput.OnChanged = func(s string) {
		uList, _ := umgr.Query(s)
		nameSearchTextInput.SetOptions(uList)
		nameSearchTextInput.ShowCompletion()
	}

	dateTextInput := widget.NewEntry()
	dateTextInput.Text = time.Now().Format("02 Jan 2006")

	feedbackTextArea := widget.NewMultiLineEntry()
	feedbackTextArea.SetPlaceHolder("Enter feedback...")

	f := widget.NewForm(
		widget.NewFormItem("Name", nameSearchTextInput),
		widget.NewFormItem("Date", dateTextInput),
		widget.NewFormItem("Feedback", feedbackTextArea),
	)
	f.OnSubmit = func() {}
	f.OnCancel = func() {}

	w.SetContent(f)
	w.Canvas().Focus(nameSearchTextInput)
	return w
}
