package cpu

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Post("/validate", validateHandler)
}

type validateRequest struct {
	Messages []string `json:"messages"`
}

type validateResponse struct {
	Processed int    `json:"processed"`
	Checksum  string `json:"checksum"`
}

const emailRegexPattern = `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

func validateHandler(res http.ResponseWriter, req *http.Request) {
	var messagesReq validateRequest
	if err := json.NewDecoder(req.Body).Decode(&messagesReq); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	var report string

	for _, m := range messagesReq.Messages {
		matched, err := regexp.MatchString(emailRegexPattern, m)
		if err != nil {
			continue
		}

		cleanMsg := m
		if matched {
			re, err := regexp.Compile(emailRegexPattern)
			if err == nil {
				cleanMsg = re.ReplaceAllString(m, "***")
			}
		}

		report = report + cleanMsg + "\n"
	}

	hash := md5.Sum([]byte(report))
	checksum := fmt.Sprintf("%x", hash)

	resp := validateResponse{
		Processed: len(messagesReq.Messages),
		Checksum:  checksum,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(resp)
}
