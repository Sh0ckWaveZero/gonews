package app

import (
	"net/http"

	"github.com/Sh0ckWaveZero/gonews/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
