package pie

import (
	"strings"
	"testing"
)

func BenchmarkStringsJoin(b *testing.B) {
	ss := Strings{"A", "utility", "library", "for", "dealing", "with", "slices", "and", "maps", "that", "focuses", "on", "type", "safety", "and", "performance"}
	for i := 0; i < b.N; i++ {
		result := ss.Join(" ")
		sinkString = result
	}
}

func BenchmarkStdlibStringsJoin(b *testing.B) {
	ss := Strings{"A", "utility", "library", "for", "dealing", "with", "slices", "and", "maps", "that", "focuses", "on", "type", "safety", "and", "performance"}
	for i := 0; i < b.N; i++ {
		result := strings.Join(ss, " ")
		sinkString = result
	}
}

// Prevent compiler from agressively optimizing away the result
var sinkString string
