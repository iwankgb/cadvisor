package perf

import (
	"github.com/google/cadvisor/watcher"
	"k8s.io/klog"
)

// StartMonitoring sets perf events up for a container.
func StartMonitoring(event watcher.ContainerEvent) {
	klog.Infof("starting monitoring for %#v", event.Name)
	if shouldIMonitor() {
		// Do whatever needs to be done in order to monitor all
		// the necessary events.
		// Launch a goroutine responsible for reading events' values.
		// Read from configuration (I do not know how it would be 
		// encoded) and set only requested events up.
	}
}

func shouldIMonitor() bool {
	// I assume that we are not going to monitor all the containers. 
	// We need to be able to choose them somehow and configuration 
	// with have to be available here.
	return true
}

// StopMonitoring tears perf evets down for container.
func StopMonitoring(event watcher.ContainerEvent) {
	klog.Infof("stopping monitoring for %#v", event.Name)
	if isMonitored() {
		// Close file descriptors, gracefully stop all gouroutines
		// make sure that final values for events are collected.
	}
}

func isMonitored() bool {
	// We do not want to close non-existing file descriptors etc.
	return true
}
