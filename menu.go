// 2018-11

package ui

import (
	"unsafe"
)

// #include "pkgui.h"
import "C"

// Menu is attached to windows if flag has been set to true


type Menu struct {
	ControlBase
	m *C.uiMenu
}


// NewMenu creates a new menu
func NewMenu(text string) *Menu {
	m := new(Menu)

	ctext := C.CString(text)
	freestr(ctext)

	m.ControlBase = NewControlBase(m, uintptr(unsafe.Pointer(m.m)))
	return m
}




//TODO old label code for reference


// Text returns the Label's text.
func (l *Label) Text() string {
	ctext := C.uiLabelText(l.l)
	text := C.GoString(ctext)
	C.uiFreeText(ctext)
	return text
}

// SetText sets the Label's text to text.
func (l *Label) SetText(text string) {
	ctext := C.CString(text)
	C.uiLabelSetText(l.l, ctext)
	freestr(ctext)
}
