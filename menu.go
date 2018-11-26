// 2018-11
// test additions by Whakapapa

package ui


import (
//	"unsafe"
)


// #include "pkgui.h"
import "C"


// Menu is attached to windows if flag has been set to true
type Menu struct {
	m *C.uiMenu
}


type MenuItem struct {
	mi *C.uiMenuItem
}


// NewMenu creates a new menu
func NewMenu(text string) *Menu {
	m := new(Menu)

	ctext := C.CString(text)
	m.m = C.uiNewMenu(ctext)
	freestr(ctext)

	return m
}


// MenuAppendSeparator adds a separator item
func (m *Menu) MenuAppendSeparator() {
	C.uiMenuAppendSeparator(m.m)
}


// creates new MenuItem object that the gui is aware of
func NewMenuItem() *MenuItem {
	return new(MenuItem)
}


// MenuAppendItem adds a custom item
func (m *Menu) MenuAppendItem(mi *MenuItem, text string) {
	ctext := C.CString(text)
	mi.mi = C.uiMenuAppendItem(m.m, ctext)
	freestr(ctext)
}


// MenuAppendAboutItem adds an about item
func (m *Menu) MenuAppendAboutItem(mi *MenuItem) {
	mi.mi = C.uiMenuAppendAboutItem(m.m)
}


// MenuAppendPreferencesItem adds preferences item
func (m *Menu) MenuAppendPreferencesItem(mi *MenuItem) {
	mi.mi = C.uiMenuAppendPreferencesItem(m.m)
}


// MenuAppendQuitItem adds a quit menu
func (m *Menu) MenuAppendQuitItem(mi *MenuItem) {
	mi.mi = C.uiMenuAppendQuitItem(m.m)
}


// MenuAppendCheckItem adds a check item
func (m *Menu) MenuAppendCheckItem(mi *MenuItem, text string) {
	ctext := C.CString(text)
	mi.mi = C.uiMenuAppendCheckItem(m.m, ctext)
	freestr(ctext)
}


// MenuItemOnClicked triggers assoc procedure
func (mi *MenuItem) MenuItemOnClicked(f func(w *Window, ret *interface{}) *interface{}) {
	ret = C.uiMenuItemOnClicked(mi.mi, w.w, unsafe.Pointer(ret))

	//TODO empty for now
	/*
	_UI_EXTERN void uiMenuItemOnClicked(uiMenuItem *m, void (*f)(uiMenuItem *sender, uiWindow *window, void *data), void *data);
	*/
}


// MenuItemEnable enables the menu
func (mi *MenuItem) MenuItemEnable() {
	C.uiMenuItemEnable(mi.mi)
}


// MenuItemDisable disables the menu
func (mi *MenuItem) MenuItemDisable() {
	C.uiMenuItemDisable(mi.mi)
}


//TODO convert check mark and result to bool
//TODO get conversion of int between C and Go working
// verify if menu item is checked or not
func (mi *MenuItem) MenuItemChecked() int {
	var checked int
	checked = C.uiMenuItemChecked(mi.mi)
	return checked
}

/*
//TODO convert check mark and result to bool
//TODO get conversion of int between C and Go working
// set the checked flag
func (mi *MenuItem) MenuItemSetChecked(checked int) {
	C.uiMenuItemSetChecked(mi.mi, checked)
}
*/

/////////////////////////////
/// input from C to digest
/*

static uiMenu **menus = NULL;
static size_t len = 0;
static size_t cap = 0;
static BOOL menusFinalized = FALSE;
static WORD curID = 100;			// start somewhere safe
static BOOL hasQuit = FALSE;
static BOOL hasPreferences = FALSE;
static BOOL hasAbout = FALSE;

struct uiMenu {
	WCHAR *name;
	uiMenuItem **items;
	size_t len;
	size_t cap;
	};

	struct uiMenuItem {
		WCHAR *name;
		int type;
		WORD id;
		void (*onClicked)(uiMenuItem *, uiWindow *, void *);
		void *onClickedData;
		BOOL disabled;				// template for new instances; kept in sync with everything else
		BOOL checked;
		HMENU *hmenus;
		size_t len;
		size_t cap;
		};

		enum {
			typeRegular,
			typeCheckbox,
			typeQuit,
			typePreferences,
			typeAbout,
			typeSeparator,
			};

			#define grow 32

			static void sync(uiMenuItem *item)
			{
				size_t i;
				MENUITEMINFOW mi;

				ZeroMemory(&mi, sizeof (MENUITEMINFOW));
				mi.cbSize = sizeof (MENUITEMINFOW);
				mi.fMask = MIIM_STATE;
				if (item->disabled)
				mi.fState |= MFS_DISABLED;
				if (item->checked)
				mi.fState |= MFS_CHECKED;

				for (i = 0; i < item->len; i++)
				if (SetMenuItemInfo(item->hmenus[i], item->id, FALSE, &mi) == 0)
				logLastError(L"error synchronizing menu items");
			}

			static void onQuitClicked(uiMenuItem *item, uiWindow *w, void *data)
			{
				if (uiprivShouldQuit())
				uiQuit();
			}

			void uiMenuItemEnable(uiMenuItem *i)
			{
				i->disabled = FALSE;
				sync(i);
			}

			void uiMenuItemDisable(uiMenuItem *i)
			{
				i->disabled = TRUE;
				sync(i);
			}

			void uiMenuItemOnClicked(uiMenuItem *i, void (*f)(uiMenuItem *, uiWindow *, void *), void *data)
			{
				if (i->type == typeQuit)
				uiprivUserBug("You can not call uiMenuItemOnClicked() on a Quit item; use uiOnShouldQuit() instead.");
				i->onClicked = f;
				i->onClickedData = data;
			}

			int uiMenuItemChecked(uiMenuItem *i)
			{
				return i->checked != FALSE;
			}

			void uiMenuItemSetChecked(uiMenuItem *i, int checked)
			{
				// use explicit values
				i->checked = FALSE;
				if (checked)
				i->checked = TRUE;
				sync(i);
			}

			static uiMenuItem *newItem(uiMenu *m, int type, const char *name)
			{
				uiMenuItem *item;

				if (menusFinalized)
				uiprivUserBug("You can not create a new menu item after menus have been finalized.");

				if (m->len >= m->cap) {
					m->cap += grow;
					m->items = (uiMenuItem **) uiprivRealloc(m->items, m->cap * sizeof (uiMenuItem *), "uiMenuitem *[]");
				}

				item = uiprivNew(uiMenuItem);

				m->items[m->len] = item;
				m->len++;

				item->type = type;
				switch (item->type) {
				case typeQuit:
					item->name = toUTF16("Quit");
					break;
				case typePreferences:
					item->name = toUTF16("Preferences...");
					break;
				case typeAbout:
					item->name = toUTF16("About");
					break;
				case typeSeparator:
					break;
				default:
					item->name = toUTF16(name);
					break;
				}

				if (item->type != typeSeparator) {
					item->id = curID;
					curID++;
				}

				if (item->type == typeQuit) {
					// can't call uiMenuItemOnClicked() here
					item->onClicked = onQuitClicked;
					item->onClickedData = NULL;
					} else
					uiMenuItemOnClicked(item, defaultOnClicked, NULL);

					return item;
				}

				uiMenuItem *uiMenuAppendItem(uiMenu *m, const char *name)
				{
					return newItem(m, typeRegular, name);
				}

				uiMenuItem *uiMenuAppendCheckItem(uiMenu *m, const char *name)
				{
					return newItem(m, typeCheckbox, name);
				}

				uiMenuItem *uiMenuAppendQuitItem(uiMenu *m)
				{
					if (hasQuit)
					uiprivUserBug("You can not have multiple Quit menu items in a program.");
					hasQuit = TRUE;
					newItem(m, typeSeparator, NULL);
					return newItem(m, typeQuit, NULL);
				}

				uiMenuItem *uiMenuAppendPreferencesItem(uiMenu *m)
				{
					if (hasPreferences)
					uiprivUserBug("You can not have multiple Preferences menu items in a program.");
					hasPreferences = TRUE;
					newItem(m, typeSeparator, NULL);
					return newItem(m, typePreferences, NULL);
				}

				uiMenuItem *uiMenuAppendAboutItem(uiMenu *m)
				{
					if (hasAbout)
					// TODO place these uiprivImplBug() and uiprivUserBug() strings in a header
					uiprivUserBug("You can not have multiple About menu items in a program.");
					hasAbout = TRUE;
					newItem(m, typeSeparator, NULL);
					return newItem(m, typeAbout, NULL);
				}

				void uiMenuAppendSeparator(uiMenu *m)
				{
					newItem(m, typeSeparator, NULL);
				}

				uiMenu *uiNewMenu(const char *name)
				{
					uiMenu *m;

					if (menusFinalized)
					uiprivUserBug("You can not create a new menu after menus have been finalized.");
					if (len >= cap) {
						cap += grow;
						menus = (uiMenu **) uiprivRealloc(menus, cap * sizeof (uiMenu *), "uiMenu *[]");
					}

					m = uiprivNew(uiMenu);

					menus[len] = m;
					len++;

					m->name = toUTF16(name);

					return m;
				}

				static void appendMenuItem(HMENU menu, uiMenuItem *item)
				{
					UINT uFlags;

					uFlags = MF_SEPARATOR;
					if (item->type != typeSeparator) {
						uFlags = MF_STRING;
						if (item->disabled)
						uFlags |= MF_DISABLED | MF_GRAYED;
						if (item->checked)
						uFlags |= MF_CHECKED;
					}
					if (AppendMenuW(menu, uFlags, item->id, item->name) == 0)
					logLastError(L"error appending menu item");

					if (item->len >= item->cap) {
						item->cap += grow;
						item->hmenus = (HMENU *) uiprivRealloc(item->hmenus, item->cap * sizeof (HMENU), "HMENU[]");
					}
					item->hmenus[item->len] = menu;
					item->len++;
				}

				static HMENU makeMenu(uiMenu *m)
				{
					HMENU menu;
					size_t i;

					menu = CreatePopupMenu();
					if (menu == NULL)
					logLastError(L"error creating menu");
					for (i = 0; i < m->len; i++)
					appendMenuItem(menu, m->items[i]);
					return menu;
				}

				HMENU makeMenubar(void)
				{
					HMENU menubar;
					HMENU menu;
					size_t i;

					menusFinalized = TRUE;

					menubar = CreateMenu();
					if (menubar == NULL)
					logLastError(L"error creating menubar");

					for (i = 0; i < len; i++) {
						menu = makeMenu(menus[i]);
						if (AppendMenuW(menubar, MF_POPUP | MF_STRING, (UINT_PTR) menu, menus[i]->name) == 0)
						logLastError(L"error appending menu to menubar");
					}

					return menubar;
				}


				static void freeMenu(uiMenu *m, HMENU submenu)
				{
					size_t i;
					uiMenuItem *item;
					size_t j;

					for (i = 0; i < m->len; i++) {
						item = m->items[i];
						for (j = 0; j < item->len; j++)
						if (item->hmenus[j] == submenu)
						break;
						if (j >= item->len)
						uiprivImplBug("submenu handle %p not found in freeMenu()", submenu);
						for (; j < item->len - 1; j++)
						item->hmenus[j] = item->hmenus[j + 1];
						item->hmenus[j] = NULL;
						item->len--;
					}
				}

				void freeMenubar(HMENU menubar)
				{
					size_t i;
					MENUITEMINFOW mi;

					for (i = 0; i < len; i++) {
						ZeroMemory(&mi, sizeof (MENUITEMINFOW));
						mi.cbSize = sizeof (MENUITEMINFOW);
						mi.fMask = MIIM_SUBMENU;
						if (GetMenuItemInfoW(menubar, i, TRUE, &mi) == 0)
						logLastError(L"error getting menu to delete item references from");
						freeMenu(menus[i], mi.hSubMenu);
					}
					// no need to worry about destroying any menus; destruction of the window they're in will do it for us
				}

				void uninitMenus(void)
				{
					uiMenu *m;
					uiMenuItem *item;
					size_t i, j;

					for (i = 0; i < len; i++) {
						m = menus[i];
						uiprivFree(m->name);
						for (j = 0; j < m->len; j++) {
							item = m->items[j];
							if (item->len != 0)
							// LONGTERM uiprivUserBug()?
							uiprivImplBug("menu item %p (%ws) still has uiWindows attached; did you forget to destroy some windows?", item, item->name);
							if (item->name != NULL)
							uiprivFree(item->name);
							if (item->hmenus != NULL)
							uiprivFree(item->hmenus);
							uiprivFree(item);
						}
						if (m->items != NULL)
						uiprivFree(m->items);
						uiprivFree(m);
					}
					if (menus != NULL)
					uiprivFree(menus);
				}

				*/
