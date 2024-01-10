package clock

import (
	"log"
)

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		log.Println(p)
	}
}

func RunSafe(fn func(), cleanups ...func()) {
	defer Recover(cleanups...)

	fn()
}
