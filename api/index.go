package handler

import (
	"bytes"
	"encoding/json"
	"image/png"
	"net/http"

	"github.com/atuy1213/textintoimage/pkg"
)

type Request struct {
	MainText string `json:"main_text"`
	SubText  string `json:"sub_text"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

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

	textImage, err := pkg.TextToImage(&pkg.TextToImageInput{
		MainText: req.MainText,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	imageInsertedImage, err := pkg.TextIntoImage(&pkg.TextIntoImageInput{
		SrcImage:      textImage.MainTextImage,
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
