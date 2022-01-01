package templates

var AnswerTemplate = `[#90ee90]Question[-]

[#00FFFF]{{ .Question.Title }}[-]

{{ BeautifyHtmlText .Question.Body }}

[#90ee90]UpVotes {{ .Question.UpVoteCount }} | Asked {{ GetDateDiffInDays .Question.CreationDate }} days ago[-]

[#90ee90]Link[-] [blue]https://stackoverflow.com/questions/{{ .Question.QuestionID }}[-]

[#90ee90]{{ .SeperatorString }}[-]

{{ $save := .SeperatorString }}

{{ range $i, $e := .Answers  }}

[yellow]({{ Add $i 1 }})[-] [#90ee90]Answer[-]

{{ BeautifyHtmlText .Body }}

[#90ee90]UpVotes {{ .UpVoteCount }} | DownVotes {{ .DownVoteCount }} | IsAccepted {{ if .IsAccepted }}yes{{ else }}no{{ end }} | LastActivity {{ GetDateDiffInDays .LastActivityDate }} days ago[-]

[#90ee90]Link[-] [blue]{{ .ShareLink }}[-]

[#90ee90]{{ $save }}[-]

{{ end }}

`
