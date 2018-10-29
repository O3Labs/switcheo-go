package client

type SwitchError struct {
	Code    int
	Message string
}

var possibleErrors map[int]SwitchError = map[int]SwitchError{
	400: SwitchError{
		Code:    400,
		Message: "Bad Request -- Your request is badly formed.",
	},
	401: SwitchError{
		Code:    401,
		Message: "Unauthorized -- You did not provide a valid signature.",
	},
	404: SwitchError{
		Code:    404,
		Message: "Not Found -- The specified endpoint or resource could not be found.",
	},
	406: SwitchError{
		Code:    406,
		Message: "Not Acceptable -- You requested a format that isn't json.",
	},
	429: SwitchError{
		Code:    429,
		Message: "Too Many Requests -- Slow down requests and use Exponential backoff timing.",
	},
	422: SwitchError{
		Code:    422,
		Message: "Unprocessible Entity -- Your request had validation errors.",
	},
	500: SwitchError{
		Code:    500,
		Message: "Internal Server Error -- We had a problem with our server. Try again later.",
	},
	503: SwitchError{
		Code:    503,
		Message: "Service Unavailable -- We're temporarily offline for maintenance. Please try again later.",
	},
}
