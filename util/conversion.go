package util

import (
	"regexp"
	"strings"

	"github.com/scaleway/scaleway-sdk-go/scw"
)

// Sometimes people state gibibytes multiplied by 1000, e.g.
// 512000000000 to represent 512GiB in bytes. However, this is wrong, as 512GiB is 512 * 1024 * 1024 * 1024 bytes
func GibiBytesMultipliedByThousandsToMebibytes(gbIn scw.Size) uint32 {
	gib := gbIn / (1000 * 1000 * 1000)

	return uint32(gib * 1024)
}

// Parse and tidy CPU names
func CleanCPUName(name string) string {
	name = strings.ToLower(name)

	// Remove extra info
	name = strings.Replace(name, "or equivalent", "", 1)

	// Remove non-alphanumeric characters
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	name = re.ReplaceAllString(name, "")

	// Trim whitespace
	name = strings.Trim(name, " ")

	return name
}
