package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type EventClient struct {
	httpClient  *http.Client
	host        string
	projectID   string
	serviceName string
}

func NewEventClient(httpClient *http.Client) *EventClient {
	host := os.Getenv("EVENT_AUDIT_HOST")
	projectID := os.Getenv("PROJECT_ID")
	serviceName := os.Getenv("SERVICE_NAME")

	return &EventClient{
		httpClient:  httpClient,
		host:        host,
		projectID:   projectID,
		serviceName: serviceName,
	}
}

func (c *EventClient) CreateEvent(
	ctx context.Context,
	event AuditEvent,
) error {
	if event.EventType == "" {
		event.EventType = "DATA_CHANGED"
	}

	if event.OccurredAt.IsZero() {
		event.OccurredAt = time.Now()
	}

	event.ProjectID = c.projectID
	event.ServiceName = c.serviceName

	url := fmt.Sprintf("%s/events", strings.TrimRight(c.host, "/"))

	jsonBody, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("gagal encode audit event: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return fmt.Errorf("gagal membuat request audit event: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim audit event: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)

		return fmt.Errorf(
			"audit event gagal dibuat, status: %d, response: %s",
			resp.StatusCode,
			string(body),
		)
	}

	return nil
}
