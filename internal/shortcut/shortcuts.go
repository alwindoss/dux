package shortcut

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var AltA = &desktop.CustomShortcut{
	KeyName:  fyne.KeyA,
	Modifier: fyne.KeyModifierAlt,
}

var AltF = &desktop.CustomShortcut{
	KeyName:  fyne.KeyF,
	Modifier: fyne.KeyModifierAlt,
}

var Escape = &desktop.CustomShortcut{
	KeyName: fyne.KeyEscape,
}

var CtrlQ = &desktop.CustomShortcut{
	KeyName:  fyne.KeyQ,
	Modifier: fyne.KeyModifierControl,
}

var CtrlN = &desktop.CustomShortcut{
	KeyName:  fyne.KeyN,
	Modifier: fyne.KeyModifierControl,
}

var SuperQ = &desktop.CustomShortcut{
	KeyName:  fyne.KeyQ,
	Modifier: fyne.KeyModifierSuper,
}

var Tab = &desktop.CustomShortcut{
	KeyName: fyne.KeyTab,
}
