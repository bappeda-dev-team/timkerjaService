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

type KepegawaianClient struct {
	BaseClient
}

func NewKepegawaianClient(host string, httpClient *http.Client) *KepegawaianClient {
	return &KepegawaianClient{
		BaseClient: newBaseClient(host, "api/v1/tpp", httpClient),
	}
}

func (c *KepegawaianClient) GetDetailPegawaiBatch(ctx context.Context, nipPegawais []string) ([]DetailPegawaiResponse, error) {
	// url check program unggulan
	url := fmt.Sprintf("%s/%s/jabatan/detail/by-nip-batch", c.host, c.path)

	log.Printf("URL: %s", url)

	// body id program unggulans
	payload := DetailPegawaiBatchRequest{
		NipPegawais: nipPegawais,
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
	log.Printf("Request: %v", req)

	// response status
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Detail Pegawai: tidak ditemukan. status: %d", res.StatusCode)
	}

	var result []DetailPegawaiResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("gagal decode response: %w", err)
	}

	return result, nil
}
