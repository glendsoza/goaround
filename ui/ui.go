package ui

func Run(query string, tags string) error {
	m := NewManager()
	m.SetQuestionQuery(query)
	m.SetQuestionTags(tags)
	return m.Run()
}
