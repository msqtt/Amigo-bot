package logcat_test

import (
	"testing"

	"github.com/mosqu1t0/Amigo-bot/utils/logcat"
)

func TestLog(t *testing.T) {
	logcat.Good("深爱你%s", "怒")
	logcat.Error("深爱你%s", "怒")
	logcat.Warn("深爱你%s", "怒")
	logcat.Debug("深爱你%s", "怒")
}
