#!/bin/bash

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=Red" \
  http://localhost:8080 > dist/color/red.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=Blue" \
  http://localhost:8080 > dist/color/blue.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=Yellow" \
  http://localhost:8080 > dist/color/yellow.jpeg
