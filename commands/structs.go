package commands

import "time"

type Defects struct {
	Bugs []Defect `json:"bugs"`
}

type Defect struct {
	ID           int       `json:"id"`
	Product      string    `json:"product"`
	Summary      string    `json:"summary"`
	Severity     string    `json:"severity"`
	Priority     string    `json:"priority"`
	CreationTime time.Time `json:"creation_time"`
	Component    []string  `json:"component"`
	Resolution   string    `json:"resolution"`
	QAContact    string    `json:"qa_contact"`
	Creator      string    `json:"creator"`
	URL          string    `json:"url"`
}

type BugHistoryList struct {
	Bugs []Bug `json:"bugs"`
}

type Bug struct {
	ID          int       `json:"id"`
	HistoryList []History `json:"history"`
}

type History struct {
	When    time.Time `json:"when"`
	Who     string    `json:"who"`
	Changes []Change  `json:"changes"`
}

type Change struct {
	Added     string `json:"added"`
	Removed   string `json:"removed"`
	FieldName string `json:"field_name"`
}
