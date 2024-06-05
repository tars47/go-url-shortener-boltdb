package urlshort

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/tars47/go-url-shortener-boltdb/store"
)

type Service struct {
	Store store.Store `json:"-"`
	Path  string
	Url   string
}

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (*Service) send(w http.ResponseWriter, res response) {
	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{status:500,message:InternalServerError}"))
		return
	}

	w.WriteHeader(res.Status)
	w.Write([]byte(bytes))
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{

			url, err := s.Store.Get(r.URL.Path)

			if err != nil || url == "" {
				s.send(w, response{http.StatusNotFound, "Not a registered route"})
				return
			}
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		}
	case http.MethodPost:
		{
			if r.URL.Path != "/" {
				http.NotFound(w, r)
				return
			}

			var req Service
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil || req.Path == "" || req.Url == "" || req.Path == req.Url {
				s.send(w, response{http.StatusBadRequest, "Malformed request body, should have path (eg: '/123') and url (eg: 'http://google.com')"})
				return

			}

			if !IsUrl(req.Url) {
				s.send(w, response{http.StatusBadRequest, "Invalid url, (eg: 'http://google.com')"})
				return
			}

			if !strings.HasPrefix(req.Path, "/") {
				req.Path = "/" + req.Path
			}

			if err := s.Store.Upsert(req.Path, req.Url); err != nil {
				s.send(w, response{http.StatusInternalServerError, "InternalServerError"})
				return
			}

			s.send(w, response{http.StatusCreated, "Created"})

		}
	case http.MethodDelete:
		{
			if err := s.Store.Delete(r.URL.Path); err != nil {

				s.send(w, response{http.StatusInternalServerError, "InternalServerError"})
				return
			}
			s.send(w, response{http.StatusOK, "Deleted"})

		}
	}

}

func IsUrl(str string) bool {
	u, err := url.Parse(str)

	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return false
	}

	return true
}
