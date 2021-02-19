package windows

// InputRecord is the data structure that ReadConsoleInput writes into.
// All Documentation originally provided by Michael Niksa, et al. at Microsoft Corporation
// under CC Attribution 4.0 International
// via https://docs.microsoft.com/en-us/windows/console/input-record-str
type InputRecord struct {
	// 0x1: Key event
	// 0x2: Will never be read when using ReadConsoleInput
	// 0x4: Window buffer size event
	// 0x8: Deprecated
	// 0x10: Deprecated
	// Original source: https://docs.microsoft.com/en-us/windows/console/input-record-str#members
	Type  uint16
	_     [2]byte // discard the next two bytes

	// Event contents can be one of:
	// KEY_EVENT (Type == 1)
	//  - Event[0] is 0x1 if the key is pressed, 0x0 if the key is released
	//  - Event[3] is the keycode of the pressed key, see
	//    https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
	//  - Event[5] is the ascii or Unicode keycode.
	//  - Event[6] stores the state of the modifier keys.
	//  Original source: https://docs.microsoft.com/en-us/windows/console/key-event-record-str
	// 
	// WINDOW_BUFFER_SIZE_EVENT (TYPE == 4)
	//  - Event[0] is the new amount of character rows
	//  - Event[1] is the new amount of character columns
	// Original source: https://docs.microsoft.com/en-us/windows/console/window-buffer-size-record-str
	Event [6]uint16
}

