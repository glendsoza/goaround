package constants

import "regexp"

// Stack overflow urls

var STACK_OVERFLOW_API_SEARCH_URL = "https://api.stackexchange.com/2.3/search/advanced"
var STACK_OVERFLOW_API_ANSWER_URL = "https://api.stackexchange.com/2.3/questions/%d/answers"
var STACK_OVERFLOW_UI_QUESTION_URL = "https://stackoverflow.com/questions/"

// Regex

var REPLACE_MULTIPLE_NEW_LINE_REGEX = regexp.MustCompile("(\n\n)+")
