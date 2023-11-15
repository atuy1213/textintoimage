package internal

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/atuy1213/textintoimage/api/pkg"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

type Request struct {
	Text      string
	Size      float64
	Width     int
	Height    int
	TopMargin int
	Color     string
	Font      string
}

func (u *handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()

		var req Request
		req.Text = r.FormValue("text")
		req.Size, _ = strconv.ParseFloat(r.FormValue("size"), 64)
		req.Width, _ = strconv.Atoi(r.FormValue("width"))
		req.Height, _ = strconv.Atoi(r.FormValue("height"))
		req.TopMargin, _ = strconv.Atoi(r.FormValue("topMargin"))
		req.Color = r.FormValue("color")
		req.Font = r.FormValue("font")

		slog.Info("rquest", "req", req)
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileSrc, _, err := r.FormFile("upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fileSrc.Close()

		srcImage, _, err := image.Decode(fileSrc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		textImage, err := pkg.TextToImage(&pkg.TextToImageInput{
			Text:          req.Text,
			Size:          req.Size,
			ImageWidth:    req.Width,
			ImageHeight:   req.Height,
			TextTopMargin: req.TopMargin,
			Color:         req.Color,
			Font:          req.Font,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		imageInsertedImage, err := pkg.TextIntoImage(&pkg.TextIntoImageInput{
			SrcImage:      srcImage,
			MainTextImage: textImage.MainTextImage,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf := &bytes.Buffer{}
		if err := png.Encode(buf, imageInsertedImage.Image); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(buf.Bytes())
	}
}

func (u *handler) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		healthz := struct {
			Now string `json:"now"`
		}{
			Now: time.Now().Format(time.RFC3339),
		}
		b, err := json.Marshal(healthz)
		if err != nil {
			slog.Error("failed to marshal healthz response", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(b)))
		w.Header().Set("Cache-control", "no-cache, no-store")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "-1")
		if _, err := w.Write(b); err != nil {
			slog.Error("failed to write healthz response", err)
			return
		}
	}
}
