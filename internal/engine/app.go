package engine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/alwindoss/dux/internal/feedback"
	"github.com/alwindoss/dux/internal/shortcut"
	"github.com/alwindoss/dux/internal/user"
)

func runApp(fmgr *feedback.Manager, umgr *user.Manager) {

	a := app.New()
	w := a.NewWindow("Dux")
	w.Resize(fyne.NewSize(800, 600))
	w.SetMainMenu(createMenu(a, fmgr, umgr))

	w.SetContent(widget.NewLabel("Hello Dux! Welcome to the world of productivity"))
	w.Canvas().AddShortcut(shortcut.AltF, func(shortcut fyne.Shortcut) {
		ShowNewFeedbackWindow(a, fmgr, umgr)
	})
	w.Canvas().AddShortcut(shortcut.AltA, func(shortcut fyne.Shortcut) {
		ShowAddActionWindow(a)
	})
	w.Canvas().AddShortcut(shortcut.CtrlQ, func(shortcut fyne.Shortcut) {
		w.Close()
	})
	w.ShowAndRun()
}

func createMenu(a fyne.App, fmgr *feedback.Manager, umgr *user.Manager) *fyne.MainMenu {
	newFeedbackMenuItem := fyne.NewMenuItem("Feedback", func() {
		ShowNewFeedbackWindow(a, fmgr, umgr)
	})
	newFeedbackMenuItem.Shortcut = shortcut.AltF

	newActionMenuItem := fyne.NewMenuItem("Action", func() {
		ShowAddActionWindow(a)
	})
	newActionMenuItem.Shortcut = shortcut.AltA

	newUserMenuItem := fyne.NewMenuItem("User", func() {
		ShowAddUserWindow(a, umgr)
	})

	newFeedbackMenu := fyne.NewMenu("New",
		newFeedbackMenuItem,
		newActionMenuItem,
		newUserMenuItem,
	)

	newMenuItem := fyne.NewMenuItem("New", nil)
	newMenuItem.ChildMenu = newFeedbackMenu

	fileSaveMenuItem := fyne.NewMenuItem("Save", func() {})
	quitMenuItem := fyne.NewMenuItem("Quit", func() {
		a.Quit()
	})
	quitMenuItem.Shortcut = shortcut.CtrlQ
	menu := fyne.NewMainMenu(
		fyne.NewMenu("File",
			newMenuItem,
			fileSaveMenuItem,
			fyne.NewMenuItemSeparator(),
			quitMenuItem,
		),
	)
	return menu
}
