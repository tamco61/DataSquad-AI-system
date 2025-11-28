package model

type TaskDTO struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	RawText        string `json:"raw_text"`
	Classification string `json:"classification,omitempty"`
	Summary        string `json:"summary,omitempty"`
	RiskLevel      string `json:"risk_level,omitempty"`

	AssigneeID  *string `json:"assignee_id,omitempty"`
	AddressedID *string `json:"addressed_id,omitempty"`
	Status      string  `json:"status"`

	SLADeadline *string `json:"sla_deadline,omitempty"`
	ResolvedAt  *string `json:"resolved_at,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TaskFilter struct {
	Status     string
	AssigneeID string
}
