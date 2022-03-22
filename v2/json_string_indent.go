package pie

import (
	"encoding/json"
	"golang.org/x/exp/constraints"
)

// JSONStringIndent returns the JSON encoded array as a string with indent applied.
//
// One important thing to note is that it will treat a nil slice as an empty
// slice to ensure that the JSON value return is always an array. See
// json.MarshalIndent for details.
func JSONStringIndent[T constraints.Ordered](ss []T, prefix, indent string) string {
	if ss == nil {
		return "[]"
	}

	// An error should not be possible.
	data, _ := json.MarshalIndent(ss, prefix, indent)

	return string(data)
}
