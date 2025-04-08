package envs

import "os"

func Set(key, val string) {
	os.Setenv(key, val)
}
