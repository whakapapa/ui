package main

import (
	"github.com/andlabs/ui"
	//	_ "github.com/andlabs/ui/winmanifest"
)

var mainWin *ui.Window

func basicTab() ui.Control {

	// buttons for first box
	bbxbButton	:= ui.NewButton("Button")
	bbxbElse		:= ui.NewButton("Else")

	// checkboxes for first box
	bbxcCheck	:= ui.NewCheckbox("Check")
	bbxcYup		:= ui.NewCheckbox("Yup")
	bbxcYup.SetChecked(true)
	bbxcNope		:= ui.NewCheckbox("Nope")

	// build the first box
	tabBox := ui.NewHorizontalBox()
	tabBox.SetPadded(true)
	tabBox.Append(bbxbButton, false)
	tabBox.Append(bbxbElse, false)
	tabBox.Append(bbxcCheck, false)
	tabBox.Append(bbxcYup, false)
	tabBox.Append(bbxcNope, false)

	// separator
	tabSep		:= ui.NewHorizontalSeparator()

	// different group form entry fields
	grpEntry		:= ui.NewEntry()
	grpPasswd	:= ui.NewPasswordEntry()
	grpSearch	:= ui.NewSearchEntry()
	grpMulti		:= ui.NewMultilineEntry()
	grpMnwrap	:= ui.NewNonWrappingMultilineEntry()

	// build the group form
	grpForm := ui.NewForm()
	grpForm.SetPadded(true)
	grpForm.Append("Entry", grpEntry, false)
	grpForm.Append("Password", grpPasswd, false)
	grpForm.Append("Search", grpSearch, false)
	grpForm.Append("Multi", grpMulti, true)
	grpForm.Append("NoWrap", grpMnwrap, true)

	// build the group
	tabGroup := ui.NewGroup("Input")
	tabGroup.SetMargined(true)
	tabGroup.SetChild(grpForm)

	// now construct the tab
	btab := ui.NewVerticalBox()
	btab.SetPadded(false)
	btab.Append(tabBox, false)
	btab.Append(tabSep, false)
	btab.Append(tabGroup, true)

	return btab
}

func numbersTab() ui.Control {

	// objects in number group
	spinbox	:= ui.NewSpinbox(0, 100)
	slider	:= ui.NewSlider(0, 100)
	pbar		:= ui.NewProgressBar()

	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})

	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})

	ip := ui.NewProgressBar()
	ip.SetValue(-1)

	// construct number box in group
	nbox := ui.NewVerticalBox()
	nbox.SetPadded(true)
	nbox.Append(spinbox, false)
	nbox.Append(slider, false)
	nbox.Append(pbar, false)
	nbox.Append(ip, false)

	// numers group
	tabNumbers := ui.NewGroup("Numbers")
	tabNumbers.SetMargined(true)
	tabNumbers.SetChild(nbox)

	// objects in lists group

	// build combo box
	cbLinux	:= "Linux"
	cbApple	:= "Apple"
	cbMS		:= "Microsoft"

	cBox		:= ui.NewCombobox()
	cBox.Append(cbLinux)
	cBox.Append(cbApple)
	cBox.Append(cbMS)

	// build editable box
	ebWater	:= "Water"
	ebBeer	:= "Beer"
	ebTea		:= "Tea"

	eBox		:= ui.NewEditableCombobox()
	eBox.Append(ebWater)
	eBox.Append(ebBeer)
	eBox.Append(ebTea)

	// build radio buttons
	rbBMW		:= "BMW"
	rbAudi	:= "Audi"
	rbMerc	:= "Mercedes"

	rButton	:= ui.NewRadioButtons()
	rButton.Append(rbBMW)
	rButton.Append(rbAudi)
	rButton.Append(rbMerc)

	// construct the list box
	lbox := ui.NewVerticalBox()
	lbox.SetPadded(true)
	lbox.Append(cBox, false)
	lbox.Append(eBox, false)
	lbox.Append(rButton, false)

	// build the lists group
	tabLists := ui.NewGroup("Lists")
	tabLists.SetMargined(true)
	tabLists.SetChild(lbox)


	// build the numbers tab
	ntab := ui.NewHorizontalBox()
	ntab.SetPadded(true)
	ntab.Append(tabNumbers, true)
	ntab.Append(tabLists, true)

	return ntab
}

func dataTab() ui.Control {

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewHorizontalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainWin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainWin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainWin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func setupUI() {

	// construct the main window tabs
	tabBasic		:= "Basic Controls"
	tabNumbers	:= "Numbers and Lists"
	tabData		:= "Data Choosers"

	mainTab := ui.NewTab()

	mainTab.Append(tabBasic, basicTab())
	mainTab.SetMargined(0, true)

	mainTab.Append(tabNumbers, numbersTab())
	mainTab.SetMargined(1, true)

	mainTab.Append(tabData, dataTab())
	mainTab.SetMargined(2, true)


	// construct the main window
	winHead := "Control Gallery"
	winX		:= 640
	winY		:= 480

	mainWin	= ui.NewWindow(winHead, winX, winY, true)

	// main window behavior
	mainWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
		return true
	})

	mainWin.SetChild(mainTab)
	mainWin.SetMargined(true)

	mainWin.Show()
}

func main() {
	ui.Main(setupUI)
}
