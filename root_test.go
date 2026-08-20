package canary

import "testing"

func TestUntrustedForkExecutesTaggedConsumer(t *testing.T) {
	t.Parallel()
	if RootMarker != "root-tags-applied" {
		t.Fatalf("RootMarker = %q, want %q", RootMarker, "root-tags-applied")
	}
}
