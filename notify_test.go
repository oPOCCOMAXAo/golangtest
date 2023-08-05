package golangtest

import (
	"os/exec"
	"testing"
)

func TestNotify(t *testing.T) {
	_ = exec.Command("notify-send", "test", "test\ntest\ntest").Run()
	_ = exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga").Run()
}
