package internal

import (
	"time"
)

type AuditAction string

const (
	AuditActionCreate AuditAction = "CREATE"
	AuditActionUpdate AuditAction = "UPDATE"
	AuditActionDelete AuditAction = "DELETE"
	AuditActionClone  AuditAction = "CLONE"
)

const (
	AuditEventTypeDataChanged = "DATA_CHANGED"
)

type AuditEvent struct {
	EventType string      `json:"event_type"`
	Action    AuditAction `json:"action"`

	ServiceName string `json:"service_name"`
	ProjectID   string `json:"project_id"`

	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`

	ActorID       *string `json:"actor_id,omitempty"`
	ActorUsername *string `json:"actor_username,omitempty"`

	RequestID *string `json:"request_id,omitempty"`
	SessionID *string `json:"session_id,omitempty"`

	OccurredAt time.Time `json:"occurred_at"`

	Before   any            `json:"before,omitempty"`
	After    any            `json:"after,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// HELPER
func NewCreateEvent(
	entityType string,
	entityID string,
	after any,
) AuditEvent {
	return AuditEvent{
		EventType:  AuditEventTypeDataChanged,
		Action:     AuditActionCreate,
		EntityType: entityType,
		EntityID:   entityID,
		OccurredAt: time.Now(),
		After:      after,
	}
}

func NewUpdateEvent(
	entityType string,
	entityID string,
	before any,
	after any,
) AuditEvent {
	return AuditEvent{
		EventType:  AuditEventTypeDataChanged,
		Action:     AuditActionUpdate,
		EntityType: entityType,
		EntityID:   entityID,
		OccurredAt: time.Now(),
		Before:     before,
		After:      after,
	}
}

func NewDeleteEvent(
	entityType string,
	entityID string,
	before any,
) AuditEvent {
	return AuditEvent{
		EventType:  AuditEventTypeDataChanged,
		Action:     AuditActionDelete,
		EntityType: entityType,
		EntityID:   entityID,
		OccurredAt: time.Now(),
		Before:     before,
	}
}

func NewCloneEvent(
	entityType string,
	entityID string,
	after any,
	metadata map[string]any,
) AuditEvent {
	return AuditEvent{
		EventType:  AuditEventTypeDataChanged,
		Action:     AuditActionClone,
		EntityType: entityType,
		EntityID:   entityID,
		OccurredAt: time.Now(),
		After:      after,
		Metadata:   metadata,
	}
}
