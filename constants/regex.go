package constants

import "regexp"

var SwaggerUIPathRegexChecker = regexp.MustCompile("/swagger-ui/*").MatchString
