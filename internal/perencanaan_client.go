package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PerencanaanClient struct {
	BaseClient
}

func NewPerencanaanClient(host string, httpClient *http.Client) *PerencanaanClient {
	return &PerencanaanClient{
		BaseClient: newBaseClient(host, "/api/v1/perencanaan/%s", httpClient),
	}
}

func (c *PerencanaanClient) GetRincianProgramUnggulans(ctx context.Context, kodeProgramUnggulans []string, tahun int) ([]TaggingPohonKinerjaItem, error) {
	// url check program unggulan
	url := fmt.Sprintf("%s/api/v1/laporan-tagging/tagging/getDetailBatch", c.host)
	// body kode program unggulans
	payload := FindByKodeProgramUnggulansRequest{
		KodeProgramUnggulan: kodeProgramUnggulans,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal encode body: %w", err)
	}

	// request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("Gagal membuat request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	} else {
		log.Printf("Tidak ada Session Id ditemukan, mungkin akan 401")
	}

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request ke program unggulan gagal: %w", err)
	}
	defer res.Body.Close()

	// response status
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Program unggulan: %v tidak ditemukan. status: %d", kodeProgramUnggulans, res.StatusCode)
	}

	var result LaporanTaggingPohonKinerjaResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}
	return result.Data, nil

}

func (c *PerencanaanClient) GetDataRincianKerjaBatch(
	ctx context.Context,
	idRekins []string,
	bulan int,
	tahun int,
) ([]DataRincianKerja, error) {

	if len(idRekins) == 0 {
		return []DataRincianKerja{}, nil
	}

	url := fmt.Sprintf(
		"%s/api/v1/perencanaan/rencana_kinerja/detail/findbatch",
		c.host,
	)

	payload := FindByIdRekinsRequest{
		Ids:   idRekins,
		Bulan: bulan,
		Tahun: tahun,
	}

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal encode body: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	} else {
		log.Printf("⚠️ session id tidak ditemukan, kemungkinan unauthorized")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// === Handle non-200 ===
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"unexpected status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var wrapper DataRincianKerjaWrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return wrapper.RencanaKinerja, nil
}

func (c *PerencanaanClient) GetNamaProgramUnggulanBatch(ctx context.Context, idProgramUnggulans []int) ([]ProgramUnggulanResponse, error) {
	// url check program unggulan
	url := fmt.Sprintf("%s/api/v1/perencanaan/program_unggulan/findbyidterkait", c.host)

	// body id program unggulans
	payload := FindByIdTerkaitRequest{
		Ids: idProgramUnggulans,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("gagal encode body: %w", err)
	}

	// request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("Gagal membuat request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	} else {
		log.Printf("Tidak ada X-Session-Id ditemukan, mungkin akan 401")
	}

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request ke perencanaan gagal: %w", err)
	}
	defer res.Body.Close()

	// response status
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Program unggulan: tidak ditemukan. status: %d", res.StatusCode)
	}

	// safe, response pasti ada
	type wrapper struct {
		Code   int                       `json:"code"`
		Status string                    `json:"status"`
		Data   []ProgramUnggulanResponse `json:"data"`
	}

	var result wrapper
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return result.Data, nil
}
