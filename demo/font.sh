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
  -F "color=black" \
  -F "font=Bold" \
  http://localhost:8080 > dist/font/bold.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=black" \
  -F "font=Regular" \
  http://localhost:8080 > dist/font/regular.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=black" \
  -F "font=Light" \
  http://localhost:8080 > dist/font/light.jpeg
