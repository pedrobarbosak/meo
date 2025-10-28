package meo

import (
	"context"
	"net/http"

	"github.com/pedrobarbosak/go-errors"
)

func (s *Service) AssignStaticIP(ctx context.Context, mac string, ip string) error {
	if _, err := s.Login(ctx); err != nil {
		return errors.New("failed to login:", err)
	}

	url := s.hostname + "/location=dhcpdstaticlease.cmd"
	resp, err := s.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.New("failed to get dhcpdstaticlease:", err)
	}
	_ = resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to get dhcpdstaticlease with status:", mac, ip, resp.Status)
	}

	return nil
}
