package targets

import "fmt"

const defaultApiVersion = "2021-09-15-preview"

func userAgent() string {
	return fmt.Sprintf("pandora/targets/%s", defaultApiVersion)
}
