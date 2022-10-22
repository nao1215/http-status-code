[![Build](https://github.com/nao1215/http-status-code/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/http-status-code/actions/workflows/build.yml)
[![PlatformUnitTests](https://github.com/nao1215/http-status-code/actions/workflows/platform_test.yml/badge.svg)](https://github.com/nao1215/http-status-code/actions/workflows/platform_test.yml)
[![codecov](https://codecov.io/gh/nao1215/http-status-code/branch/main/graph/badge.svg?token=AGqQgVDcL1)](https://codecov.io/gh/nao1215/http-status-code)
[![reviewdog](https://github.com/nao1215/http-status-code/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/http-status-code/actions/workflows/reviewdog.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/http-status-code)](https://goreportcard.com/report/github.com/nao1215/http-status-code)
# http-status-code - Check HTTP status code with CLI
The hsc (http-status-code) command uses the HTTP status code as a search keyword and print the meaning of that code and the RFC.

# How to install
### Use "go install"
If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install).
```
$ go install github.com/nao1215/http-status-code@latest
```

### Install from Package or Binary
[The release page](https://github.com/nao1215/http-status-code/releases) contains packages in .deb, .rpm, and .apk formats.
  

# How to use
### Search HTTP status code
```
$ hsc search 501
501 Not Implemented (ref.=RFC9110, Section 15.6.2)

$ hsc search 401 201 303
401 Unauthorized (ref.=RFC9110, Section 15.5.2)
201 Created (ref.=RFC9110, Section 15.3.2)
303 See Other (ref.=RFC9110, Section 15.4.4)
```

### List up HTTP status code
```
$ hsc list
+------------------+---------------------------------+--------------------------+
| HTTP STATUS CODE |           DESCRIPTION           |           RFC            |
+------------------+---------------------------------+--------------------------+
|              100 | Continue                        | RFC9110, Section 15.2.1  |
|              101 | Switching Protocols             | RFC9110, Section 15.2.2  |
|              102 | Processing                      | RFC2518                  |
|              103 | Early Hints                     | RFC8297                  |
|              200 | OK                              | RFC9110, Section 15.3.1  |
|              201 | Created                         | RFC9110, Section 15.3.2  |
|              202 | Accepted                        | RFC9110, Section 15.3.3  |
|              203 | Non-Authoritative Information   | RFC9110, Section 15.3.4  |
|              204 | No Content                      | RFC9110, Section 15.3.5  |
|              205 | Reset Content                   | RFC9110, Section 15.3.6  |
|              206 | Partial Content                 | RFC9110, Section 15.3.7  |
|              207 | Multi-Status                    | RFC4918                  |
|              208 | Already Reported                | RFC5842                  |
|              226 | IM Used                         | RFC3229                  |
|              300 | Multiple Choices                | RFC9110, Section 15.4.1  |
|              301 | Moved Permanently               | RFC9110, Section 15.4.2  |
|              302 | Found                           | RFC9110, Section 15.4.3  |
|              303 | See Other                       | RFC9110, Section 15.4.4  |
|              304 | Not Modified                    | RFC9110, Section 15.4.5  |
|              305 | Use Proxy                       | RFC9110, Section 15.4.6  |
|              306 | (Unused)                        | RFC9110, Section 15.4.7  |
|              307 | Temporary Redirect              | RFC9110, Section 15.4.8  |
|              308 | Permanent Redirect              | RFC9110, Section 15.4.9  |
|              400 | Bad Request                     | RFC9110, Section 15.5.1  |
|              401 | Unauthorized                    | RFC9110, Section 15.5.2  |
|              402 | Payment Required                | RFC9110, Section 15.5.3  |
|              403 | Forbidden                       | RFC9110, Section 15.5.4  |
|              404 | Not Found                       | RFC9110, Section 15.5.5  |
|              405 | Method Not Allowed              | RFC9110, Section 15.5.6  |
|              406 | Not Acceptable                  | RFC9110, Section 15.5.7  |
|              407 | Proxy Authentication Required   | RFC9110, Section 15.5.8  |
|              408 | Request Timeout                 | RFC9110, Section 15.5.9  |
|              409 | Conflict                        | RFC9110, Section 15.5.10 |
|              410 | Gone                            | RFC9110, Section 15.5.11 |
|              411 | Length Required                 | RFC9110, Section 15.5.12 |
|              412 | Precondition Failed             | RFC9110, Section 15.5.13 |
|              413 | Content Too Large               | RFC9110, Section 15.5.14 |
|              414 | URI Too Long                    | RFC9110, Section 15.5.15 |
|              415 | Unsupported Media Type          | RFC9110, Section 15.5.16 |
|              416 | Range Not Satisfiable           | RFC9110, Section 15.5.17 |
|              417 | Expectation Failed              | RFC9110, Section 15.5.18 |
|              418 | (Unused)                        | RFC9110, Section 15.5.19 |
|              421 | Misdirected Request             | RFC9110, Section 15.5.20 |
|              422 | Unprocessable Content           | RFC9110, Section 15.5.21 |
|              423 | Locked                          | RFC4918                  |
|              424 | Failed Dependency               | RFC4918                  |
|              425 | Too Early                       | RFC8470                  |
|              426 | Upgrade Required                | RFC9110, Section 15.5.22 |
|              427 | Unassigned                      | Unassigned               |
|              428 | Precondition Required           | RFC6585                  |
|              429 | Too Many Requests               | RFC6585                  |
|              430 | Unassigned                      | Unassigned               |
|              431 | Request Header Fields Too Large | RFC6585                  |
|              451 | Unavailable For Legal Reasons   | RFC7725                  |
|              500 | Internal Server Error           | RFC9110, Section 15.6.1  |
|              501 | Not Implemented                 | RFC9110, Section 15.6.2  |
|              502 | Bad Gateway                     | RFC9110, Section 15.6.3  |
|              503 | Service Unavailable             | RFC9110, Section 15.6.4  |
|              504 | Gateway Timeout                 | RFC9110, Section 15.6.5  |
|              505 | HTTP Version Not Supported      | RFC9110, Section 15.6.6  |
|              506 | Variant Also Negotiates         | RFC2295                  |
|              507 | Insufficient Storage            | RFC4918                  |
|              508 | Loop Detected                   | RFC5842                  |
|              509 | Unassigned                      | Unassigned               |
|              510 | Not Extended (OBSOLETED)        | RFC2774                  |
|              511 | Network Authentication Required | RFC6585                  |
+------------------+---------------------------------+--------------------------+
```

### Use github.com/nao1215/hsc/http package
```
package main

import (
	"fmt"

	"github.com/nao1215/hsc/http"
)

func main() {
	h := http.New()
	h = h.Search("404")
	fmt.Printf("http status code=%s, description=%s, rfc=%s\n", h.Code, h.Description, h.RFC)

	// Output:
	// http status code=404, description=Not Found, rfc=RFC9110, Section 15.5.5
}
```

# Contributing
First off, thanks for taking the time to contribute! Contributions are not only related to development. For example, GitHub Star motivates me to develop!

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/http-status-code/issues)

# LICENSE
The hsc project is licensed under the terms of the [MIT License](./LICENSE).
