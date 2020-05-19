package ornot

// GoOrNot runs a function async, or not (synchronously)
//
func GoOrNot(async bool, fn func()) {
	if !async {
		fn()
	} else {
		go fn()
	}
}

// AwaitOrNot runs a function sync, or not (asynchronously)
//
func AwaitOrNot(sync bool, fn func()) {
	GoOrNot(!sync, fn)
}
