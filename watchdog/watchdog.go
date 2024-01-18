package watchdog

import (
	"fmt"
	"strconv"
	"time"

	"github.com/coreos/go-systemd/daemon"
)

// return only WatchdogSec valuse from service file.
func GetInterval() (time.Duration, error) { return daemon.SdWatchdogEnabled(false) }

// set watchdog timer
func SetInterval(t time.Duration) error {
	ts := strconv.FormatInt(t.Microseconds(), 10)
	ok, err := daemon.SdNotify(false, "WATCHDOG_USEC="+ts)
	if !ok {
		return fmt.Errorf("not set watchdog interval:%s error:%v", t.String(), err)
	}
	return nil
}

// disable watchdog timer
func Disable() error { return SetInterval(0) }

// reset watchdog timer
func Refresh() error { return SendNotify(daemon.SdNotifyWatchdog) }

// send a ready signal to systemd
func ServiceStarted() error { return SendNotify(daemon.SdNotifyReady) }

func SendNotify(signal string) error {
	ok, err := daemon.SdNotify(false, daemon.SdNotifyWatchdog)
	if !ok {
		return fmt.Errorf("watchdog not updated error:%v", err)
	}
	return nil
}
