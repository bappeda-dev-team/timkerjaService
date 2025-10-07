package internal

import (
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
	Id                        int     `json:"id"`
	KodeProgramUnggulan       string  `json:"kode_program_unggulan"`
	NamaTagging               string  `json:"nama_program_unggulan"`
	KeteranganProgramUnggulan *string `json:"rencana_implementasi"`
	Keterangan                *string `json:"keterangan"`
	TahunAwal                 string  `json:"tahun_awal"`
	TahunAkhir                string  `json:"tahun_akhir"`
	IsActive                  bool    `json:"is_active"`
}

func NewPerencanaanClient(host string, httpClient *http.Client) *PerencanaanClient {
	return &PerencanaanClient{
		httpClient:      httpClient,
		host:            host,
		perencanaanPath: "/api/v1/perencanaan/%s",
	}
}

func (c *PerencanaanClient) GetProgramUnggulan(idProgramUnggulan int) (*ProgramUnggulanResponse, error) {
	// url check program unggulan
	url := fmt.Sprintf("%s/api/v1/perencanaan/program_unggulan/detail/%d", c.host, idProgramUnggulan)
	log.Printf("url: %v", url)
	// request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Gagal membuat request: %w", err)
	}

	if os.Getenv("APP_ENV") == "development" {
		sessionID := os.Getenv("DEV_SESSION_ID")
		req.Header.Set("X-Session-Id", sessionID)
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
	log.Printf("resp: %v", result)

	return result.Data, nil
}
