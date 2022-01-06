package ui

func Run(query string, tags string) error {
	// Initialize the manager
	m := NewManager()
	// Set the query
	m.SetQuestionQuery(query)
	// Set the tags
	m.SetQuestionTags(tags)
	// Before calling run query and tags need to be set
	return m.Run()
}
