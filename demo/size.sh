#!/bin/bash

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=50" \
  -F "width=150" \
  -F "height=500" \
  -F "topMargin=50" \
  -F "color=Black" \
  http://localhost:8080 > dist/size/small.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=90" \
  -F "width=300" \
  -F "height=100" \
  -F "topMargin=90" \
  -F "color=Black" \
  http://localhost:8080 > dist/size/medium.jpeg

curl \
  -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "upload=@/Users/s22286/src/github.com/atuy1213/textintoimage/src/sample.jpeg" \
  -F "text=冷蔵庫" \
  -F "size=150" \
  -F "width=450" \
  -F "height=400" \
  -F "topMargin=150" \
  -F "color=Black" \
  http://localhost:8080 > dist/size/large.jpeg
