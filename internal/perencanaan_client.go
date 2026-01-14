// TODO: ambil data pokin dari laporan tagging
package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (c *PerencanaanClient) GetDataRincianKerja(
	ctx context.Context,
	idRekin string,
	idPegawai string,
	bulan int,
	tahun int,
) (*DataRincianKerja, error) {
	// TODO: optimize with batch id pegawai id rekin
	url := fmt.Sprintf("%s/api/v1/perencanaan/rencana_kinerja/%s/pegawai/%s/rincian_kak_by_bulan_tahun?bulan=%d&tahun=%d", c.host, idRekin, idPegawai, bulan, tahun)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	sessionID := getSessionID(ctx)
	if sessionID != "" {
		req.Header.Set("X-Session-Id", sessionID)
	} else {
		log.Printf("Session Id ditemukan, mungkin akan 401")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var wrapper DataRincianKerjaWrapper
	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	if len(wrapper.RencanaKinerja) == 0 {
		return nil, nil
	}

	return &wrapper.RencanaKinerja[0], nil
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
		log.Printf("X-Session-Id diterapkan: %s", sessionID)
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
