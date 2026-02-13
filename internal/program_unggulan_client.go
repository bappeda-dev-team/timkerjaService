package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProgramUnggulanClient struct {
	BaseClient
}

func NewProgramUnggulanClient(host string, httpClient *http.Client) *ProgramUnggulanClient {
	return &ProgramUnggulanClient{
		BaseClient: newBaseClient(host, "/api/v1/laporan-tagging", httpClient),
	}
}

func (c *ProgramUnggulanClient) GetLaporanProgramUnggulanByTahun(ctx context.Context, tahun int) ([]TaggingPohonKinerjaItem, error) {
	if tahun <= 0 {
		return nil, fmt.Errorf("[ProgramUnggulanError] tahun wajib terisi")
	}
	tahunStr := strconv.Itoa(tahun)
	queries := make([]map[string]string, 0)
	queries = append(queries, map[string]string{
		// Ganti kebutuhan tagging disini
		"nama_tagging": "Program Unggulan Bupati",
		"tahun":        tahunStr,
	})
	// url get program unggulan bupati
	url, err := buildURL(c.host, c.path+"/laporan/tagging_pokin", queries)
	if err != nil {
		return nil, fmt.Errorf("[ProgramUnggulanError] Error membuat query: %w", err)
	}

	log.Printf("URL: %s", url)

	// request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("[ProgramUnggulanError] Gagal membuat request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	} else {
		log.Printf("[ProgramUnggulanError] Sesi user tidak ditemukan")
	}

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[ProgramUnggulanError] Request ke gagal: %w", err)
	}
	defer res.Body.Close()

	// response status
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("[ProgramUnggulanError] status: %d", res.StatusCode)
	}

	// safe, response pasti ada
	type wrapper struct {
		Status  int                              `json:"status"`
		Message string                           `json:"message"`
		Data    []LaporanTagPokinTahunanResponse `json:"data"`
	}

	var result wrapper
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("[ProgramUnggulanError] gagal decode response: %w", err)
	}
	if len(result.Data) == 0 {
		return []TaggingPohonKinerjaItem{}, nil
	}
	var dataTagging []TaggingPohonKinerjaItem

	for _, data := range result.Data {
		for _, tag := range data.PohonKinerjas {
			dataTagging = append(dataTagging, tag)
		}
	}

	return dataTagging, nil
}
