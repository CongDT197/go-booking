package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "Invalid param",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Check token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token has expired",
	ERROR_AUTH_TOKEN:               "Error when generate token",
	ERROR_AUTH:                     "Auth error",
	ERROR_BOOK_NOTFOUND:            "Book not found",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
