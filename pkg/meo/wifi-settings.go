package meo

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pedrobarbosak/meo/pkg/meo/requests"

	"github.com/pedrobarbosak/go-errors"
)

func (s *Service) SetWifiSettings(ctx context.Context, settings requests.PutWifiSettings) error {
	settings.Version = "1.0"

	body, err := json.Marshal(settings)
	if err != nil {
		return errors.New("failed to marshal settings:", err)
	}

	resp, err := s.doRequest(ctx, http.MethodPut, s.hostname+"locallmngt.cmd?request=/generic-wifi", nil, bytes.NewReader(body))
	if err != nil {
		return errors.New("failed to set wifi settings:", err)
	}
	_ = resp.Body.Close()

	return nil
}
