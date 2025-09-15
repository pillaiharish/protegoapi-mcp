package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type GradeReq struct {
	Code string `json:"code"`
}
type Finding struct {
	Rule     string `json:"rule"`
	Detail   string `json:"detail"`
	Severity string `json:"severity"`
	Passed   bool   `json:"passed"`
}
type GradeRes struct {
	Pass     bool      `json:"pass"`
	Findings []Finding `json:"findings"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// health check
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// tiny stub grader: fails if code contains "SELECT * FROM"
	r.Post("/api/grade", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req GradeReq
		_ = json.NewDecoder(r.Body).Decode(&req)

		res := GradeRes{
			Pass: true,
			Findings: []Finding{},
		}
		if req.Code != "" && containsSelectStar(req.Code) {
			res.Pass = false
			res.Findings = append(res.Findings, Finding{
				Rule:     "SQLI-1",
				Detail:   "Found raw 'SELECT * FROM' — avoid unsafe patterns. Use explicit columns and parameters.",
				Severity: "high",
				Passed:   false,
			})
		}
		json.NewEncoder(w).Encode(res)
	})

	log.Println("server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func containsSelectStar(s string) bool {
	return strings.Contains(strings.ToLower(s), "select * from")
}
