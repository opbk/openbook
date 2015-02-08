package ses

import "time"

type Mail struct {
	Timestamp   time.Time `json:"timestamp"`
	MessageId   string    `json:"messageId"`
	Source      string    `json:"source"`
	Destination []string  `json:"destination"`
}

type BouncedRecipients struct {
	EmailAddress   string `json:"emailAddress"`
	Action         string `json:"action"`
	Status         string `json:"status"`
	DiagnosticCode string `json:"diagnosticCode"`
}

type Bounce struct {
	BounceType        string              `json:"bounceType"`
	BounceSubType     string              `json:"bounceSubType"`
	BouncedRecipients []BouncedRecipients `json:"bouncedRecipients"`
	Timestamp         time.Time           `json:"timestamp"`
	FeedbackId        string              `json:"feedbackId"`
	ReportingMTA      string              `json:"reportingMTA"`
}

type ComplaintRecipients struct {
	EmailAddress string `json:"emailAddress"`
}

type Complaint struct {
	ComplainedRecipients  []ComplaintRecipients `json:"complainedRecipients"`
	Timestamp             time.Time             `json:"timestamp"`
	FeedbackId            string                `json:"feedbackId"`
	UserAgent             string                `json:"userAgent"`
	ComplaintFeedbackType string                `json:"complaintFeedbackType"`
	ArrivalDate           time.Time             `json:"arrivalDate"`
}

type BounceNotification struct {
	NotificationType string `json:"notificationType"`
	Mail             Mail   `json:"mail"`
	Bounce           Bounce `json:"bounce"`
}

type ComplaintNotification struct {
	NotificationType string    `json:"notificationType"`
	Mail             Mail      `json:"mail"`
	Complaint        Complaint `json:"complaint"`
}
