package internal

import (
	"bytes"
	"encoding/json"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log/slog"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/atuy1213/textintoimage/api/pkg"
	"github.com/atuy1213/textintoimage/api/pkg/textimage"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
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
	Src       multipart.File
}

func (u *handler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var req Request
		req.Text = r.FormValue("text")
		req.Size, _ = strconv.ParseFloat(r.FormValue("size"), 64)
		req.Width, _ = strconv.Atoi(r.FormValue("width"))
		req.Height, _ = strconv.Atoi(r.FormValue("height"))
		req.TopMargin, _ = strconv.Atoi(r.FormValue("topMargin"))
		req.Color = r.FormValue("color")
		req.Font = r.FormValue("font")
		file, _, err := r.FormFile("upload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		req.Src = file
		slog.Info("rquest", "req", req)

		srcImage, _, err := image.Decode(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		textImage := image.NewRGBA(image.Rect(0, 0, req.Width, req.Height))
		src := textimage.NewColor(req.Color).Uniform()
		face, err := textimage.NewFont(req.Font).Face(req.Size)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dr := &font.Drawer{
			Dst:  textImage,
			Src:  src,
			Face: *face,
			Dot:  fixed.Point26_6{},
		}
		dr.Dot.X = (fixed.I(req.TopMargin) - dr.MeasureString(req.Text)) / 2
		dr.Dot.Y = fixed.I(req.TopMargin)
		dr.DrawString(req.Text)

		imageInsertedImage, err := pkg.TextIntoImage(&pkg.TextIntoImageInput{
			SrcImage:      srcImage,
			MainTextImage: textImage,
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
