// +build !linux

package resctrl

import "github.com/google/cadvisor/stats"

type manager struct {
	stats.NoopManager
}