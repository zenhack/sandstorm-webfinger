package main

type Jrd struct {
	Subject    string             `json:"subject"`
	Aliases    []string           `json:"aliases,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
	Links      []Link             `json:"links,omitempty"`
}

type Link struct {
	Rel        string             `json:"rel"`
	HRef       string             `json:"href"`
	Type       string             `json:"type,omitempty"`
	Titles     map[string]string  `json:"titles,omitempty`
	Properties map[string]*string `json:"properties,omitempty"`
}
