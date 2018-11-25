package main

//import "github.com/andlabs/ui"
import "github.com/whakapapa/ui"


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
	boxF := ui.NewHorizontalBox()
	boxF.SetPadded(true)
	boxF.Append(bbxbButton, false)
	boxF.Append(bbxbElse, false)
	boxF.Append(bbxcCheck, false)
	boxF.Append(bbxcYup, false)
	boxF.Append(bbxcNope, false)

	// separator
	boxSep		:= ui.NewHorizontalSeparator()

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
	groupInput := ui.NewGroup("Input")
	groupInput.SetMargined(true)
	groupInput.SetChild(grpForm)

	// build basic control result
	ctlBasic := ui.NewVerticalBox()
	ctlBasic.SetPadded(false)
	ctlBasic.Append(boxF, false)
	ctlBasic.Append(boxSep, false)
	ctlBasic.Append(groupInput, true)

	return ctlBasic
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
	boxNumber := ui.NewVerticalBox()
	boxNumber.SetPadded(true)
	boxNumber.Append(spinbox, false)
	boxNumber.Append(slider, false)
	boxNumber.Append(pbar, false)
	boxNumber.Append(ip, false)

	// numers group
	groupNumbers := ui.NewGroup("Numbers")
	groupNumbers.SetMargined(true)
	groupNumbers.SetChild(boxNumber)

	// objects in lists group

	// build combo box
	txtCBlinux	:= "Linux"
	txtCBapple	:= "Apple"
	txtCBms		:= "Microsoft"

	cBox			:= ui.NewCombobox()
	cBox.Append(txtCBlinux)
	cBox.Append(txtCBapple)
	cBox.Append(txtCBms)

	// build editable box
	txtEwater	:= "Water"
	txtEbeer		:= "Beer"
	txtEtea		:= "Tea"

	eBox			:= ui.NewEditableCombobox()
	eBox.Append(txtEwater)
	eBox.Append(txtEbeer)
	eBox.Append(txtEtea)

	// build radio buttons
	txtRbmw		:= "BMW"
	txtRaudi		:= "Audi"
	txtRmerc		:= "Mercedes"

	rButton	:= ui.NewRadioButtons()
	rButton.Append(txtRbmw)
	rButton.Append(txtRaudi)
	rButton.Append(txtRmerc)

	// construct the list box
	boxList := ui.NewVerticalBox()
	boxList.SetPadded(true)
	boxList.Append(cBox, false)
	boxList.Append(eBox, false)
	boxList.Append(rButton, false)

	// build the lists group
	groupLists := ui.NewGroup("Lists")
	groupLists.SetMargined(true)
	groupLists.SetChild(boxList)


	// build number control result
	ctlNumber := ui.NewHorizontalBox()
	ctlNumber.SetPadded(true)
	ctlNumber.Append(groupNumbers, true)
	ctlNumber.Append(groupLists, true)

	return ctlNumber
}

func dataTab(parentWin *ui.Window) ui.Control {

	// picker objects
	pDate		:= ui.NewDatePicker()
	pTime		:= ui.NewTimePicker()
	pDtime	:= ui.NewDateTimePicker()
	pFont		:= ui.NewFontButton()
	pColor	:= ui.NewColorButton()

	// construct box
	boxPick := ui.NewVerticalBox()
	boxPick.SetPadded(true)
	boxPick.Append(pDate, false)
	boxPick.Append(pTime, false)
	boxPick.Append(pDtime, false)
	boxPick.Append(pFont, false)
	boxPick.Append(pColor, false)

	// build choosers - open file
	fEntryOpen := ui.NewEntry()
	fEntryOpen.SetReadOnly(true)

	fOpenButton := ui.NewButton("Open File")
	fOpenButton.OnClicked(func(*ui.Button) {
		fNameOpen := ui.OpenFile(parentWin)
		if fNameOpen == "" {
			fNameOpen = "(cancelled)"
		}
		fEntryOpen.SetText(fNameOpen)
	})

	// build choosers - save file
	fEntrySave := ui.NewEntry()
	fEntrySave.SetReadOnly(true)

	fSaveButton := ui.NewButton("Save File")
	fSaveButton.OnClicked(func(*ui.Button) {
		fNameSave := ui.SaveFile(parentWin)
		if fNameSave == "" {
			fNameSave = "(cancelled)"
		}
		fEntrySave.SetText(fNameSave)
	})

	// build grid - choosers
	gridChooser := ui.NewGrid()
	gridChooser.SetPadded(true)
	gridChooser.Append(fOpenButton, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	gridChooser.Append(fEntryOpen, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)
	gridChooser.Append(fSaveButton, 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	gridChooser.Append(fEntrySave, 1, 1, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	// create message button
	messageButton := ui.NewButton("Message Box")
	messageButton.OnClicked(func(*ui.Button) {
		ui.MsgBox(parentWin, "Listen", "Good on ya mate, howyadoinalright")
	})

	// create error button
	errorButton := ui.NewButton("Error Box")
	errorButton.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(parentWin, "Panic", "Something went south.")
	})

	// build grid - message
	gridMessage := ui.NewGrid()
	gridMessage.SetPadded(true)
	gridMessage.Append(messageButton, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	gridMessage.Append(errorButton, 1, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	// Create separator
	boxSep := ui.NewVerticalSeparator()

	// build box and attach grid
	boxChoose := ui.NewVerticalBox()
	boxChoose.SetPadded(true)
	boxChoose.Append(gridChooser, false)
	boxChoose.Append(gridMessage, false)

	// build data control result
	ctlData := ui.NewHorizontalBox()
	ctlData.SetPadded(true)
	ctlData.Append(boxPick, false)
	ctlData.Append(boxSep, false)
	ctlData.Append(boxChoose, true)

	return ctlData
}

func createWindow() {

	// menu structure first
	mainMenu := ui.NewMenu("Main")
	backMenu := ui.NewMenu("Back")

	// attach items to main menu


	// attach items to main menu
//	mainMenu.MenuAppendAboutItem()
//	mainMenu.MenuAppendQuitItem()
//	mainMenu.MenuItemEnable()
//	mainMenu.MenuAppendItem("dudu")
	mainMenu.MenuAppendSeparator()

	// construct the main window
	txtWin	:= "Control Gallery"
	winX		:= 640
	winY		:= 480
	menBar	:= true
	mainWin	:= ui.NewWindow(txtWin, winX, winY, menBar)



	// construct the main window tabs
	txtBasic		:= "Basic Controls"
	txtNumbers	:= "Numbers and Lists"
	txtData		:= "Data Choosers"

	mainTab := ui.NewTab()
	mainTab.Append(txtBasic, basicTab())
	mainTab.SetMargined(0, true)
	mainTab.Append(txtNumbers, numbersTab())
	mainTab.SetMargined(1, true)
	mainTab.Append(txtData, dataTab(mainWin))
	mainTab.SetMargined(2, true)

	// quit button behavior
	/*
	backMquit.MenuItemOnClicked(func(*ui.MenuItem) {
		ui.Quit()
		return true
	})
	*/

	// main window behavior
	mainWin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWin.Destroy()
		return true
	})


	// launch the main window
	mainWin.SetChild(mainMenu)
	mainWin.SetChild(backMenu)
	mainWin.SetChild(mainTab)
	mainWin.SetMargined(true)
	mainWin.Show()
}

func main() {
	ui.Main(createWindow)
}
