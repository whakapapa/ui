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
