package handler

import (
	"log/slog"
	"net/http"
	"os"
	"path"
	"strings"
)

func LsHandler(w http.ResponseWriter, r *http.Request) {
	wd, _ := os.Getwd()
	dirs := strings.Split(r.URL.Path, "/")[2:]
	dir, _ := os.ReadDir(path.Join(wd, strings.Join(dirs, "/")))
	slog.Info("Loading font file...", "pwd", wd, "ls", dir)
	w.WriteHeader(http.StatusOK)
	for _, dir := range dir {
		w.Write([]byte(dir.Name()))
		w.Write([]byte("\n"))
	}
}
