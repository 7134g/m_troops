package clock

import "github.com/zeromicro/go-zero/core/rescue"

func RunSafe(fn func(), cleanups ...func()) {
	defer rescue.Recover(cleanups...)

	fn()
}
