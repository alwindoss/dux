package engine

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/alwindoss/dux/internal/user"
	"github.com/google/uuid"
)

func ShowAddUserWindow(a fyne.App, mgr *user.Manager) {
	w := createNewUserWindow(a, mgr)
	w.Resize(fyne.NewSquareSize(600))
	w.Show()
}

func createNewUserWindow(a fyne.App, mgr *user.Manager) fyne.Window {
	w := a.NewWindow("New User")
	w.Resize(fyne.NewSquareSize(600))

	nameTextInput := widget.NewEntry()

	f := widget.NewForm(
		widget.NewFormItem("Name", nameTextInput),
	)
	f.OnSubmit = func() {
		id := uuid.New()
		user := user.User{
			ID:   id.String(),
			Name: nameTextInput.Text,
		}
		mgr.Save(&user)
		nameTextInput.SetText("")
	}
	f.OnCancel = func() {
		w.Close()
	}

	w.SetContent(f)
	w.Canvas().Focus(nameTextInput)

	return w
}
