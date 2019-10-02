// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package eventlog

import (
	"errors"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

const (
	// Log levels.
	Info    = windows.EVENTLOG_INFORMATION_TYPE
	Warning = windows.EVENTLOG_WARNING_TYPE
	Error   = windows.EVENTLOG_ERROR_TYPE
)

const eventLogKeyName = `SYSTEM\CurrentControlSet\Services\EventLog`
const addKeyName = eventLogKeyName + `\` + `Application`

// Install modifies PC registry to allow logging with an event source src.
// It adds all required keys and values to the event log registry key.
// Install uses msgFile as the event message file. If useExpandKey is true,
// the event message file is installed as REG_EXPAND_SZ value,
// otherwise as REG_SZ. Use bitwise of log.Error, log.Warning and
// log.Info to specify events supported by the new event source.
func Install(src, msgFile string, useExpandKey bool, eventsSupported uint32) error {
	sk, err := createSubKey(registry.LOCAL_MACHINE, addKeyName, src)
	if err != nil {
		return err
	}

	err = sk.SetDWordValue("CustomSource", 1)
	if err != nil {
		return err
	}
	if useExpandKey {
		err = sk.SetExpandStringValue("EventMessageFile", msgFile)
	} else {
		err = sk.SetStringValue("EventMessageFile", msgFile)
	}
	if err != nil {
		return err
	}
	err = sk.SetDWordValue("TypesSupported", eventsSupported)
	if err != nil {
		return err
	}
	return nil
}

// InstallAsEventCreate is the same as Install, but uses
// %SystemRoot%\System32\EventCreate.exe as the event message file.
func InstallAsEventCreate(src string, eventsSupported uint32) error {
	return Install(src, "%SystemRoot%\\System32\\EventCreate.exe", true, eventsSupported)
}

// Remove deletes all registry elements installed by the correspondent Install.
func Remove(src string) error {
	appkey, err := registry.OpenKey(registry.LOCAL_MACHINE, addKeyName, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer appkey.Close()
	return registry.DeleteKey(appkey, src)
}

// InstallCustomLog creates a custom event log under Microsoft Event Viewer.
func InstallCustomLog(name string, src string, eventsSupported uint32) error {
	k, err := createSubKey(registry.LOCAL_MACHINE, eventLogKeyName, name)
	if err != nil {
		return errors.New(name + " subkey could not be created")
	}
	defer k.Close()

	err = k.SetDWordValue("TypesSupported", eventsSupported)
	if err != nil {
		return errors.New("TypesSupported could not be created")
	}

	lk, err := createSubKey(registry.LOCAL_MACHINE, eventLogKeyName + `\` + name, name)
	if err != nil {
		return errors.New(name + " " + name + " subkey could not be created")
	}
	defer lk.Close()

	err = lk.SetExpandStringValue("EventMessageFile", "C:\\Windows\\Microsoft.NET\\Framework\\v4.0.30319\\EventLogMessages.dll")
	if err != nil {
		return errors.New("EventMessageFile")
	}

	sk, err := createSubKey(registry.LOCAL_MACHINE, eventLogKeyName + `\` + name, src)
	if err != nil {
		return err
	}
	defer sk.Close()

	err = sk.SetExpandStringValue("EventMessageFile", "C:\\Windows\\Microsoft.NET\\Framework\\v4.0.30319\\EventLogMessages.dll")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

// RemoveCustomLog deletes all registry elements installed by the correspondent InstallCustomLog.
func RemoveCustomLog(name string, src string) error {
	appkey, err := registry.OpenKey(registry.LOCAL_MACHINE, eventLogKeyName + `\` + name, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer appkey.Close()
	err = registry.DeleteKey(appkey, name)
	if err != nil {
		return err
	}
	err = registry.DeleteKey(appkey, src)
	if err != nil {
		return err
	}
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, eventLogKeyName, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()
	return registry.DeleteKey(key, name)
}

func createSubKey(key registry.Key, path string, keyName string) (registry.Key, error) {
	k, err := registry.OpenKey(key, path, registry.CREATE_SUB_KEY)
	if err != nil {
		return k, errors.New(path + " path could not be opened")
	}
	defer k.Close()
	sk, alreadyExist, err := registry.CreateKey(key, path + `\` + keyName, registry.SET_VALUE)
	if err != nil {
		return sk, errors.New(keyName + " key could not be created")
	}
	if alreadyExist {
		return sk, errors.New(path + `\` + keyName + " registry key already exists")
	}
	return sk, nil
}
