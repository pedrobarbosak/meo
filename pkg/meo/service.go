package meo

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/pedrobarbosak/go-errors"
)

type Service struct {
	username   string
	password   string
	hostname   string
	httpClient *http.Client
}

func (s *Service) doRequest(ctx context.Context, method string, url string, headers map[string]string, body ...io.Reader) (*http.Response, error) {
	var bodyReader io.Reader
	if len(body) > 0 {
		bodyReader = body[0]
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, errors.New("failed to create request:", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:143.0) Gecko/20100101 Firefox/143.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("DNT", "1")
	req.Header.Set("Referer", s.hostname+"/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, errors.New("failed to execute request:", err)
	}

	return resp, nil
}

func (s *Service) Login(ctx context.Context) (string, error) {
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(s.username+":"+s.password)),
	}

	url := s.hostname + "/index.html"
	resp, err := s.doRequest(ctx, http.MethodGet, url, headers)
	if err != nil {
		return "", errors.New("failed to login:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("login failed with status:", resp.Status)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "SESSIONID" {
			return cookie.Value, nil
		}
	}

	return "", errors.New("session cookie not found")
}

func New(username string, password string, hostname ...string) (*Service, error) {
	host := "https://192.168.1.254"
	if len(hostname) > 0 {
		if !strings.HasPrefix(hostname[0], "http") {
			hostname[0] = "https://" + hostname[0]
		}

		host = strings.TrimSuffix(hostname[0], "/")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.New("failed to create cookie jar:", err)
	}

	httpClient := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	s := &Service{
		username:   username,
		password:   password,
		hostname:   host,
		httpClient: httpClient,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if _, err = s.Login(ctx); err != nil {
		return nil, errors.New("failed to login:", err)
	}

	return s, nil
}
