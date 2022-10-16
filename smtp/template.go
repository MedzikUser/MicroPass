package smtp

import (
	"bytes"
	"html/template"
)

func ParseActivationTemplate(url string) (*string, error) {
	templateData := struct {
		URL string
	}{
		URL: url,
	}

	body, err := parseTemplate("assets/email/activation.html", templateData)

	return body, err
}

func parseTemplate(file string, data interface{}) (*string, error) {
	// open the file to parse the template
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}

	// parse the template
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return nil, err
	}

	// encode the template buffer into a string
	body := buf.String()

	return &body, nil
}
