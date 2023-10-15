package voteragent

import (
	"testing"
)

func Test_CreatePref(t *testing.T) {
	alts := CreatePref(10)
	if len(alts) != 10 {

		t.Errorf("%v", alts)

	}
}
