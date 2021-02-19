package windows

// InputRecord is the data structure that ReadConsoleInput writes into.
// See: https://docs.microsoft.com/en-us/windows/console/input-record-str
type InputRecord struct {
	// Type can be one of the following:
	// 0x1: The event is of type KEY_EVENT: https://docs.microsoft.com/en-us/windows/console/key-event-record-str
	// 0x2: The event is of type MOUSE_EVENT and will never be read when using ReadConsoleInput
	// 0x4: The event is of type WINDOW_BUFFER_SIZE_EVENT: https://docs.microsoft.com/en-us/windows/console/window-buffer-size-record-str
	// 0x8: The event is of type MENU_EVENT and should be ignored.
	// 0x10: The event is of type FOCUS_EVENT: https://docs.microsoft.com/en-us/windows/console/focus-event-record-str
	Type  uint16
	_ [2]byte // discard the next two bytes

	// Event contents can be one of:
	// KEY_EVENT (Type == 1)
	//  - Event[0] is 0x1 if the key is pressed, 0x0 if the key is released
	//  - Event[1] and Event[2] is random garbage
	//  - Event[3] is the keycode of the pressed key. The numbers are completely different
	//    to normal ASCII keycode, see
	//    https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
	//  - Event[4] is the keyboard scan code.
	//  - Event[5] contains the typed ascii or Unicode (not UTF8!!) keycode.
	//    from 0x20 through 0x7f they are equal
	//  - Event[6] stores the state of the modifier keys. See https://docs.microsoft.com/en-us/windows/console/key-event-record-str
	// WINDOW_BUFFER_SIZE_EVENT (TYPE == 4)
	//  - Event[0] is the new amount of character rows
	//  - Event[1] is the new amount of character columns
	// FOCUS_EVENT (Type == 16)
	//  - Event[0] is 0x1 if the window is now focused and 0x0 if the window is now unfocused
	Event [6]uint16
}

