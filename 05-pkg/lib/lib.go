package lib

func init(){
	//always run thins function when imported
}

func IsDigit(c int32) bool {
	return '0' <= c && c <= '9'
}

func isSpace(c int32) bool {
	//User of this library cannot access this function.
	//only the function starts with capital letters are visible to user.
	switch c{
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
