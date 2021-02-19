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

	// Data contents are:
	// If the event is a key event (Type == 1):
	//  - Data[0] is 0x1 if the key is pressed, 0x0 if the key is released
	//  - Data[3] is the keycode of the pressed key, see
	//    https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
	//  - Data[5] is the ascii or Unicode keycode.
	//  - Data[6] stores the state of the modifier keys.
	//  Original source: https://docs.microsoft.com/en-us/windows/console/key-event-record-str
	// 
	// If the event is a window buffer size event (Type == 4):
	//  - Data[0] is the new amount of character rows
	//  - Data[1] is the new amount of character columns
	// Original source: https://docs.microsoft.com/en-us/windows/console/window-buffer-size-record-str
	Data [6]uint16
}

