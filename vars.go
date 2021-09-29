package main

type SonarQubeCallBackData struct {
	ServerURL  string `json:"serverUrl,omitempty"`
	TaskID     string `json:"taskId,omitempty"`
	Status     string `json:"status,omitempty"`
	AnalysedAt string `json:"analysedAt,omitempty"`
	Revision   string `json:"revision,omitempty"`
	ChangedAt  string `json:"changedAt,omitempty"`
	Project    struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"project,omitempty"`
	Branch struct {
		Name   string `json:"name,omitempty"`
		Type   string `json:"type,omitempty"`
		IsMain bool   `json:"isMain,omitempty"`
		URL    string `json:"url,omitempty"`
	} `json:"branch,omitempty"`
	QualityGate struct {
		Name       string `json:"name,omitempty"`
		Status     string `json:"status,omitempty"`
		Conditions []struct {
			Metric         string `json:"metric,omitempty"`
			Operator       string `json:"operator,omitempty"`
			Status         string `json:"status,omitempty"`
			ErrorThreshold string `json:"errorThreshold,omitempty"`
		} `json:"conditions,omitempty"`
	}
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type MeasuresData struct {
	Measures []MeasureData `json:"measures,omitempty"`
}

type MeasureData struct {
	Metric    string       `json:"metric,omitempty"`
	Value     string       `json:"value,omitempty"`
	Component string       `json:"component,omitempty"`
	Periods   []PeriodData `json:"periods,omitempty"`
}

type PeriodData struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

type DingMsg struct {
	MsgType    string         `json:"msgtype"`
	ActionCard DingActionCard `json:"actionCard,omitempty"`
}

type DingActionCard struct {
	Title          string              `json:"title,omitempty"`
	Text           string              `json:"text,omitempty"`
	BtnOrientation string              `json:"btnOrientation,omitempty"`
	Btns           []DingActionCardBtn `json:"btns,omitempty"`
}

type DingActionCardBtn struct {
	Title     string `json:"title,omitempty"`
	ActionURL string `json:"actionURL,omitempty"`
}
