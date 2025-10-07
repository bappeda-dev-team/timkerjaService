package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type PerencanaanClient struct {
	httpClient      *http.Client
	host            string
	perencanaanPath string
}

// DATA HELPER
type ProgramUnggulanResponse struct {
	Id                        int    `json:"id"`
	KodeProgramUnggulan       string `json:"kode_program_unggulan"`
	NamaTagging               string `json:"nama_program_unggulan"`
	KeteranganProgramUnggulan string `json:"rencana_implementasi"`
	Keterangan                string `json:"keterangan"`
	TahunAwal                 string `json:"tahun_awal"`
	TahunAkhir                string `json:"tahun_akhir"`
	IsActive                  bool   `json:"is_active"`
}

func NewPerencanaanClient(host string, httpClient *http.Client) *PerencanaanClient {
	return &PerencanaanClient{
		httpClient:      httpClient,
		host:            host,
		perencanaanPath: "/api/v1/perencanaan/%s",
	}
}

// key untuk context session id
type ctxKey string

const SessionIDKey ctxKey = "X-Session-Id"

// Inject session ID ke context (opsional)
func WithSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionIDKey, sessionID)
}

// Ambil session ID dari context
func getSessionID(ctx context.Context) string {
	if v := ctx.Value(SessionIDKey); v != nil {
		if s, ok := v.(string); ok && s != "" {
			return s
		}
	}
	return os.Getenv("DEV_SESSION_ID") // fallback
}

func (c *PerencanaanClient) GetProgramUnggulan(ctx context.Context, idProgramUnggulan int) (*ProgramUnggulanResponse, error) {
	// url check program unggulan
	url := fmt.Sprintf("%s/api/v1/perencanaan/program_unggulan/detail/%d", c.host, idProgramUnggulan)
	// request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Gagal membuat request: %w", err)
	}

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
		log.Printf("🪪 X-Session-Id diterapkan: %s", sessionID)
	} else {
		log.Printf("⚠️ Tidak ada X-Session-Id ditemukan, mungkin akan 401")
	}

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request ke perencanaan gagal: %w", err)
	}
	defer res.Body.Close()

	// response status
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Program unggulan: %d tidak ditemukan. status: %d", idProgramUnggulan, res.StatusCode)
	}

	// safe, response pasti ada
	type wrapper struct {
		Code   int                      `json:"code"`
		Status string                   `json:"status"`
		Data   *ProgramUnggulanResponse `json:"data"`
	}

	var result wrapper
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return result.Data, nil
}
