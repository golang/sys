// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package svc_test

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func getState(t *testing.T, s *mgr.Service) svc.State {
	status, err := s.Query()
	if err != nil {
		t.Fatalf("Query(%s) failed: %s", s.Name, err)
	}
	return status.State
}

func testState(t *testing.T, s *mgr.Service, want svc.State) {
	have := getState(t, s)
	if have != want {
		t.Fatalf("%s state is=%d want=%d", s.Name, have, want)
	}
}

func waitState(t *testing.T, s *mgr.Service, want svc.State) {
	for i := 0; ; i++ {
		have := getState(t, s)
		if have == want {
			return
		}
		if i > 10 {
			t.Fatalf("%s state is=%d, waiting timeout", s.Name, have)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

// stopAndDeleteIfInstalled stops and deletes service name,
// if the service is running and / or installed.
func stopAndDeleteIfInstalled(t *testing.T, m *mgr.Mgr, name string) {
	s, err := m.OpenService(name)
	if err != nil {
		// Service is not installed.
		return

	}
	defer s.Close()

	// Make sure the service is not running, otherwise we won't be able to delete it.
	if getState(t, s) == svc.Running {
		_, err = s.Control(svc.Stop)
		if err != nil {
			t.Fatalf("Control(%s) failed: %s", s.Name, err)
		}
		waitState(t, s, svc.Stopped)
	}

	err = s.Delete()
	if err != nil {
		t.Fatalf("Delete failed: %s", err)
	}
}

func TestExample(t *testing.T) {
	if os.Getenv("GO_BUILDER_NAME") == "" {
		// Don't install services on arbitrary users' machines.
		t.Skip("skipping test that modifies system services: GO_BUILDER_NAME not set")
	}
	if testing.Short() {
		t.Skip("skipping test in short mode that modifies system services")
	}

	const name = "svctestservice"

	m, err := mgr.Connect()
	if err != nil {
		t.Fatalf("SCM connection failed: %s", err)
	}
	defer m.Disconnect()

	exepath := filepath.Join(t.TempDir(), "a.exe")
	o, err := exec.Command("go", "build", "-o", exepath, "golang.org/x/sys/windows/svc/example").CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build service program: %v\n%v", err, string(o))
	}

	stopAndDeleteIfInstalled(t, m, name)

	s, err := m.CreateService(name, exepath, mgr.Config{DisplayName: "x-sys svc test service"}, "-name", name)
	if err != nil {
		t.Fatalf("CreateService(%s) failed: %v", name, err)
	}
	defer s.Close()

	args := []string{"is", "manual-started", fmt.Sprintf("%d", rand.Int())}

	testState(t, s, svc.Stopped)
	err = s.Start(args...)
	if err != nil {
		t.Fatalf("Start(%s) failed: %s", s.Name, err)
	}
	waitState(t, s, svc.Running)
	time.Sleep(1 * time.Second)

	// testing deadlock from issues 4.
	_, err = s.Control(svc.Interrogate)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	_, err = s.Control(svc.Interrogate)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	time.Sleep(1 * time.Second)

	_, err = s.Control(svc.Stop)
	if err != nil {
		t.Fatalf("Control(%s) failed: %s", s.Name, err)
	}
	waitState(t, s, svc.Stopped)

	err = s.Delete()
	if err != nil {
		t.Fatalf("Delete failed: %s", err)
	}

	out, err := exec.Command("wevtutil.exe", "qe", "Application", "/q:*[System[Provider[@Name='"+name+"']]]", "/rd:true", "/c:10").CombinedOutput()
	if err != nil {
		t.Fatalf("wevtutil failed: %v\n%v", err, string(out))
	}
	want := strings.Join(append([]string{name}, args...), "-")
	// Test context passing (see servicemain in sys_386.s and sys_amd64.s).
	want += "-123456"
	if !strings.Contains(string(out), want) {
		t.Errorf("%q string does not contain %q", out, want)
	}
}

func TestIsAnInteractiveSession(t *testing.T) {
	isInteractive, err := svc.IsAnInteractiveSession()
	if err != nil {
		t.Fatal(err)
	}
	if !isInteractive {
		t.Error("IsAnInteractiveSession returns false when running interactively.")
	}
}

func TestIsWindowsService(t *testing.T) {
	isSvc, err := svc.IsWindowsService()
	if err != nil {
		t.Fatal(err)
	}
	if isSvc {
		t.Error("IsWindowsService returns true when not running in a service.")
	}
}

func TestIsWindowsServiceWhenParentExits(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "parent" {
		// in parent process

		// Start the child and exit quickly.
		child := exec.Command(os.Args[0], "-test.run=^TestIsWindowsServiceWhenParentExits$")
		child.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=child")
		err := child.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "child start failed: %v", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if os.Getenv("GO_WANT_HELPER_PROCESS") == "child" {
		// in child process
		dumpPath := os.Getenv("GO_WANT_HELPER_PROCESS_FILE")
		if dumpPath == "" {
			// We cannot report this error. But main test will notice
			// that we did not create dump file.
			os.Exit(1)
		}
		var msg string
		isSvc, err := svc.IsWindowsService()
		if err != nil {
			msg = err.Error()
		}
		if isSvc {
			msg = "IsWindowsService returns true when not running in a service."
		}
		err = os.WriteFile(dumpPath, []byte(msg), 0644)
		if err != nil {
			// We cannot report this error. But main test will notice
			// that we did not create dump file.
			os.Exit(2)
		}
		os.Exit(0)
	}

	// Run in a loop until it fails.
	for i := 0; i < 10; i++ {
		childDumpPath := filepath.Join(t.TempDir(), "issvc.txt")

		parent := exec.Command(os.Args[0], "-test.run=^TestIsWindowsServiceWhenParentExits$")
		parent.Env = append(os.Environ(),
			"GO_WANT_HELPER_PROCESS=parent",
			"GO_WANT_HELPER_PROCESS_FILE="+childDumpPath)
		parentOutput, err := parent.CombinedOutput()
		if err != nil {
			t.Errorf("parent failed: %v: %v", err, string(parentOutput))
		}
		for i := 0; ; i++ {
			if _, err := os.Stat(childDumpPath); err == nil {
				break
			}
			time.Sleep(100 * time.Millisecond)
			if i > 10 {
				t.Fatal("timed out waiting for child output file to be created.")
			}
		}
		childOutput, err := os.ReadFile(childDumpPath)
		if err != nil {
			t.Fatalf("reading child output failed: %v", err)
		}
		if got, want := string(childOutput), ""; got != want {
			t.Fatalf("child output: want %q, got %q", want, got)
		}
	}
}

func TestServiceRestart(t *testing.T) {
	if os.Getenv("GO_BUILDER_NAME") == "" {
		// Don't install services on arbitrary users' machines.
		t.Skip("Skipping test that modifies system services: GO_BUILDER_NAME not set")
	}
	if testing.Short() {
		t.Skip("Skipping test in short mode that modifies system services")
	}

	const name = "svctestservice"

	m, err := mgr.Connect()
	if err != nil {
		t.Fatalf("SCM connection failed: %v", err)
	}
	defer m.Disconnect()

	// Build the service executable
	exepath := filepath.Join(t.TempDir(), "a.exe")
	o, err := exec.Command("go", "build", "-o", exepath, "golang.org/x/sys/windows/svc/example").CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to build service program: %v\n%v", err, string(o))
	}

	// Ensure any existing service is stopped and deleted
	stopAndDeleteIfInstalled(t, m, name)

	// Create the service
	s, err := m.CreateService(name, exepath, mgr.Config{DisplayName: "x-sys svc test service"})
	if err != nil {
		t.Fatalf("CreateService(%s) failed: %v", name, err)
	}
	defer s.Close()

	// Set the service to restart on failure
	actions := []mgr.RecoveryAction{
		{Type: mgr.ServiceRestart, Delay: 1 * time.Second}, // Restart after 1 second
	}
	err = s.SetRecoveryActions(actions, 0)
	if err != nil {
		t.Fatalf("Failed to set service recovery actions: %v", err)
	}

	// Set the flag to perform recovery actions on non-crash failures
	err = s.SetRecoveryActionsOnNonCrashFailures(true)
	if err != nil {
		t.Fatalf("Failed to set RecoveryActionsOnNonCrashFailures: %v", err)
	}

	// Start the service
	testState(t, s, svc.Stopped)
	err = s.Start()
	if err != nil {
		t.Fatalf("Start(%s) failed: %v", s.Name, err)
	}

	// Wait for the service to start
	waitState(t, s, svc.Running)

	// Get the initial process ID
	status, err := s.Query()
	if err != nil {
		t.Fatalf("Query(%s) failed: %v", s.Name, err)
	}
	initialPID := status.ProcessId
	t.Logf("Initial PID: %d", initialPID)

	// Wait up to 30 seconds for the PID to change, indicating a restart
	var newPID uint32
	success := false
	for i := 0; i < 30; i++ {
		time.Sleep(1 * time.Second)

		status, err = s.Query()
		if err != nil {
			t.Fatalf("Query(%s) failed: %v", s.Name, err)
		}
		newPID = status.ProcessId

		if newPID != 0 && newPID != initialPID {
			success = true
			t.Logf("Service restarted successfully, new PID: %d", newPID)
			break
		}
	}

	if !success {
		t.Fatalf("Service did not restart within the expected time")
	}

	// Cleanup: Stop and delete the service
	_, err = s.Control(svc.Stop)
	if err != nil {
		t.Fatalf("Control(%s) failed: %v", s.Name, err)
	}
	waitState(t, s, svc.Stopped)

	err = s.Delete()
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
}
