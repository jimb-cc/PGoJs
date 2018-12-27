package processing

////////////////////////////////////////////////////////////
//     					TYPOGRAPHY

// Text(str, x, y, [x2], [y2])
func Text(str, x, y interface{}, extraParameters ...interface{}) {
	switch len(extraParameters) {
	case 0:
		pG.Call("text", str, x, y)
		break
	case 2:
		pG.Call("text", str, x, y, extraParameters[0], extraParameters[1])
		break

	default:
		println("Error in Text function: incorrect number of arguments, got:%v", len(extraParameters))
		return
	}
}

// TextFont(str)
func TextFont(str interface{}) {
	pG.Call("textFont", str)
}

// TextSize(int32)
func TextSize(theSize interface{}) {
	pG.Call("textSize", theSize)
}
