package constants

import "regexp"

// Stack overflow urls
const (
	STACK_OVERFLOW_API_SEARCH_URL  = "https://api.stackexchange.com/2.3/search/advanced"
	STACK_OVERFLOW_API_ANSWER_URL  = "https://api.stackexchange.com/2.3/questions/%d/answers"
	STACK_OVERFLOW_UI_QUESTION_URL = "https://stackoverflow.com/questions/"
	ANSWER_FOOTER                  = "[black:#00FFFF]Ctrl+R[-:-] Change Query [black:#00FFFF]Backspace[-:-] Go Back [black:#00FFFF]Ctrl+C[-:-] Quit"
	QUESTION_FOOTER                = "[black:#00FFFF]Ctrl+R[-:-] Change Query [black:#00FFFF]Ctrl+C[-:-] Quit"
	FORM_FOOTER                    = "[black:#00FFFF]Ctrl+R[-:-] Go Back [black:#00FFFF]Ctrl+C[-:-] Quit"
)

// Regex
var (
	REPLACE_MULTIPLE_NEW_LINE_REGEX = regexp.MustCompile("(\n\n)+")
	PYTHON_EXEPECTED_ERRORS_REGEX   = regexp.MustCompile(`KeyboardInterrupt|SystemExit|GeneratorExit`)
)
