// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package windows

const (
	EVENTLOG_SUCCESS          = 0
	EVENTLOG_ERROR_TYPE       = 1
	EVENTLOG_WARNING_TYPE     = 2
	EVENTLOG_INFORMATION_TYPE = 4
	EVENTLOG_AUDIT_SUCCESS    = 8
	EVENTLOG_AUDIT_FAILURE    = 16
)

// EVT_SUBSCRIBE_FLAGS enumeration
const (
	EvtSubscribeToFutureEvents      = 1
	EvtSubscribeStartAtOldestRecord = 2
	EvtSubscribeStartAfterBookmark  = 3
)

// EVT_RENDER_FLAGS enumeration
const (
	EvtRenderEventValues = iota
	EvtRenderEventXML
	EvtRenderBookmark
)

//sys	RegisterEventSource(uncServerName *uint16, sourceName *uint16) (handle Handle, err error) [failretval==0] = advapi32.RegisterEventSourceW
//sys	DeregisterEventSource(handle Handle) (err error) = advapi32.DeregisterEventSource
//sys	ReportEvent(log Handle, etype uint16, category uint16, eventId uint32, usrSId uintptr, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) (err error) = advapi32.ReportEventW
//sys	EvtSubscribe(session Handle, signalEvent Handle, channelPath *uint16, query *uint16, bookmark Handle, context uintptr, callback syscall.Handle, flags uint32) (handle Handle, err error) [failretval==0] = wevtapi.EvtSubscribe
//sys	EvtClose(object Handle) (err error) = wevtapi.EvtClose
//sys	EvtNext(resultSet Handle, eventArraySize uint32, eventArray *Handle, timeout uint32, flags uint32, numReturned *uint32) (err error) = wevtapi.EvtNext
//sys	EvtRender(context Handle, fragment Handle, flags uint32, bufferSize uint32, buffer *byte, bufferUsed *uint32, propertyCount *uint32) (err error) = wevtapi.EvtRender
//sys	EvtCreateBookmark(bookmarkXML *uint16) (handle Handle, err error) [failretval==0] = wevtapi.EvtCreateBookmark
//sys	EvtUpdateBookmark(bookmark Handle, event Handle) (err error) = wevtapi.EvtUpdateBookmark
