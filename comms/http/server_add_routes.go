package http

import (
	"fmt"
	"net/http"
)

func (s *Server) AnnounceToTheJoined(sock string) string {
	return ""
}

func (s *Server) PingNode(sock string) error {
	url := "http:" + sock + "//" + urlPing

	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		return fmt.Errorf("HTTP GET request to: %s, got: %w", sock, errReq)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	var client http.Client

	resp, errCall := client.Do(req)
	if errCall != nil {
		return fmt.Errorf("ping on URL: %s returns: %w", url, errCall)
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return fmt.Errorf("status on ping to URL: %s is: %s", url, resp.Status)
	}

	return nil
}

func (s *Server) SendRangesForNode(sock string, ranges []string) string {
	return ""
}
