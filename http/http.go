// Package http returns information on HTTP status codes.
package http

// HTTP is HTTP status code information. HTTP hava status code,
// status code description, and referenced RFC.
type HTTP struct {
	// Code is http status Code
	Code string
	// Description is Description for http status code
	Description string
	// RFC is referenced RFC
	RFC string
}

var hs = []*HTTP{
	/*
		{
			Code:        "1xx",
			Description: "Informational - Request received, continuing process",
			RFC:         "",
		},
		{
			Code:        "2xx",
			Description: "Success - The action was successfully received, understood, and accepted",
			RFC:         "",
		},
		{
			Code:        "3xx",
			Description: "Redirection - Further action must be taken in order to complete the request",
			RFC:         "",
		},
		{
			Code:        "4xx",
			Description: "Client Error - The request contains bad syntax or cannot be fulfilled",
			RFC:         "",
		},
		{
			Code:        "5xx",
			Description: "Server Error - The server failed to fulfill an apparently valid request",
			RFC:         "",
		},
	*/
	{
		Code:        "100",
		Description: "Continue",
		RFC:         "RFC9110, Section 15.2.1",
	},
	{
		Code:        "101",
		Description: "Switching Protocols",
		RFC:         "RFC9110, Section 15.2.2",
	},
	{
		Code:        "102",
		Description: "Processing",
		RFC:         "RFC2518",
	},
	{
		Code:        "103",
		Description: "Early Hints",
		RFC:         "RFC8297",
	},
	{
		Code:        "200",
		Description: "OK",
		RFC:         "RFC9110, Section 15.3.1",
	},
	{
		Code:        "201",
		Description: "Created",
		RFC:         "RFC9110, Section 15.3.2",
	},
	{
		Code:        "202",
		Description: "Accepted",
		RFC:         "RFC9110, Section 15.3.3",
	},
	{
		Code:        "203",
		Description: "Non-Authoritative Information",
		RFC:         "RFC9110, Section 15.3.4",
	},
	{
		Code:        "204",
		Description: "No Content",
		RFC:         "RFC9110, Section 15.3.5",
	},
	{
		Code:        "205",
		Description: "Reset Content",
		RFC:         "RFC9110, Section 15.3.6",
	},
	{
		Code:        "206",
		Description: "Partial Content",
		RFC:         "RFC9110, Section 15.3.7",
	},
	{
		Code:        "207",
		Description: "Multi-Status",
		RFC:         "RFC4918",
	},
	{
		Code:        "208",
		Description: "Already Reported",
		RFC:         "RFC5842",
	},
	{
		Code:        "226",
		Description: "IM Used",
		RFC:         "RFC3229",
	},
	{
		Code:        "300",
		Description: "Multiple Choices",
		RFC:         "RFC9110, Section 15.4.1",
	},
	{
		Code:        "301",
		Description: "Moved Permanently",
		RFC:         "RFC9110, Section 15.4.2",
	},
	{
		Code:        "302",
		Description: "Found",
		RFC:         "RFC9110, Section 15.4.3",
	},
	{
		Code:        "303",
		Description: "See Other",
		RFC:         "RFC9110, Section 15.4.4",
	},
	{
		Code:        "304",
		Description: "Not Modified",
		RFC:         "RFC9110, Section 15.4.5",
	},
	{
		Code:        "305",
		Description: "Use Proxy",
		RFC:         "RFC9110, Section 15.4.6",
	},
	{
		Code:        "306",
		Description: "(Unused)",
		RFC:         "RFC9110, Section 15.4.7",
	},
	{
		Code:        "307",
		Description: "Temporary Redirect",
		RFC:         "RFC9110, Section 15.4.8",
	},
	{
		Code:        "308",
		Description: "Permanent Redirect",
		RFC:         "RFC9110, Section 15.4.9",
	},
	{
		Code:        "400",
		Description: "Bad Request",
		RFC:         "RFC9110, Section 15.5.1",
	},
	{
		Code:        "401",
		Description: "Unauthorized",
		RFC:         "RFC9110, Section 15.5.2",
	},
	{
		Code:        "402",
		Description: "Payment Required",
		RFC:         "RFC9110, Section 15.5.3",
	},
	{
		Code:        "403",
		Description: "Forbidden",
		RFC:         "RFC9110, Section 15.5.4",
	},
	{
		Code:        "404",
		Description: "Not Found",
		RFC:         "RFC9110, Section 15.5.5",
	},
	{
		Code:        "405",
		Description: "Method Not Allowed",
		RFC:         "RFC9110, Section 15.5.6",
	},
	{
		Code:        "406",
		Description: "Not Acceptable",
		RFC:         "RFC9110, Section 15.5.7",
	},
	{
		Code:        "407",
		Description: "Proxy Authentication Required",
		RFC:         "RFC9110, Section 15.5.8",
	},
	{
		Code:        "408",
		Description: "Request Timeout",
		RFC:         "RFC9110, Section 15.5.9",
	},
	{
		Code:        "409",
		Description: "Conflict",
		RFC:         "RFC9110, Section 15.5.10",
	},
	{
		Code:        "410",
		Description: "Gone",
		RFC:         "RFC9110, Section 15.5.11",
	},
	{
		Code:        "411",
		Description: "Length Required",
		RFC:         "RFC9110, Section 15.5.12",
	},
	{
		Code:        "412",
		Description: "Precondition Failed",
		RFC:         "RFC9110, Section 15.5.13",
	},
	{
		Code:        "413",
		Description: "Content Too Large",
		RFC:         "RFC9110, Section 15.5.14",
	},
	{
		Code:        "414",
		Description: "URI Too Long",
		RFC:         "RFC9110, Section 15.5.15",
	},
	{
		Code:        "415",
		Description: "Unsupported Media Type",
		RFC:         "RFC9110, Section 15.5.16",
	},
	{
		Code:        "416",
		Description: "Range Not Satisfiable",
		RFC:         "RFC9110, Section 15.5.17",
	},
	{
		Code:        "417",
		Description: "Expectation Failed",
		RFC:         "RFC9110, Section 15.5.18",
	},
	{
		Code:        "418",
		Description: "(Unused)",
		RFC:         "RFC9110, Section 15.5.19",
	},
	{
		Code:        "421",
		Description: "Misdirected Request",
		RFC:         "RFC9110, Section 15.5.20",
	},
	{
		Code:        "422",
		Description: "Unprocessable Content",
		RFC:         "RFC9110, Section 15.5.21",
	},
	{
		Code:        "423",
		Description: "Locked",
		RFC:         "RFC4918",
	},
	{
		Code:        "424",
		Description: "Failed Dependency",
		RFC:         "RFC4918",
	},
	{
		Code:        "425",
		Description: "Too Early",
		RFC:         "RFC8470",
	},
	{
		Code:        "426",
		Description: "Upgrade Required",
		RFC:         "RFC9110, Section 15.5.22",
	},
	{
		Code:        "427",
		Description: "Unassigned",
		RFC:         "Unassigned",
	},
	{
		Code:        "428",
		Description: "Precondition Required",
		RFC:         "RFC6585",
	},
	{
		Code:        "429",
		Description: "Too Many Requests",
		RFC:         "RFC6585",
	},
	{
		Code:        "430",
		Description: "Unassigned",
		RFC:         "Unassigned",
	},
	{
		Code:        "431",
		Description: "Request Header Fields Too Large",
		RFC:         "RFC6585",
	},
	{
		Code:        "451",
		Description: "Unavailable For Legal Reasons",
		RFC:         "RFC7725",
	},
	{
		Code:        "500",
		Description: "Internal Server Error",
		RFC:         "RFC9110, Section 15.6.1",
	},
	{
		Code:        "501",
		Description: "Not Implemented",
		RFC:         "RFC9110, Section 15.6.2",
	},
	{
		Code:        "502",
		Description: "Bad Gateway",
		RFC:         "RFC9110, Section 15.6.3",
	},
	{
		Code:        "503",
		Description: "Service Unavailable",
		RFC:         "RFC9110, Section 15.6.4",
	},
	{
		Code:        "504",
		Description: "Gateway Timeout",
		RFC:         "RFC9110, Section 15.6.5",
	},
	{
		Code:        "505",
		Description: "HTTP Version Not Supported",
		RFC:         "RFC9110, Section 15.6.6",
	},
	{
		Code:        "506",
		Description: "Variant Also Negotiates",
		RFC:         "RFC2295",
	},
	{
		Code:        "507",
		Description: "Insufficient Storage",
		RFC:         "RFC4918",
	},
	{
		Code:        "508",
		Description: "Loop Detected",
		RFC:         "RFC5842",
	},
	{
		Code:        "509",
		Description: "Unassigned",
		RFC:         "Unassigned",
	},
	{
		Code:        "510",
		Description: "Not Extended (OBSOLETED)",
		RFC:         "RFC2774",
	},
	{
		Code:        "511",
		Description: "Network Authentication Required",
		RFC:         "RFC6585",
	},
}

// New return empty HTTP sturct variable
func New() *HTTP {
	return &HTTP{}
}

// Search returns an HTTP structure for the specified HTTP status code.
// Search returns an empty HTTP structure if there is no code matching
// the specified HTTP status code
func (h *HTTP) Search(statusCode string) *HTTP {
	for _, v := range hs {
		if v.Code == statusCode {
			return v
		}
	}
	return &HTTP{}
}

// All returns a slice of the HTTP structure and
// stores all HTTP status code information in them([]*HTTP).
func (h *HTTP) All() []*HTTP {
	return hs
}
