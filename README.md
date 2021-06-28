# CSV File mini interpreter
CSV interpreter in Golang

This project consists of developping a CSV file interpter using Golang programming language.

Project's structure is as follow : 

-MAIN DIR
  |
  |-- APP DIR (File analyser logic)
  |   |
  |   |-- API DIR
  |   |      |-- server.go (api startpoint)
  |   |      |
  |   |-- DATA DIR
  |   |      |-- files.csv (temp files storage)
  |   |      |
  |   |-- Interpreter
  |   |      |-- interpreter (file as inputs, array of results as outputs)
  |   |      |-- parser (file as inputs, matrix as outputs)
  |   |      |-- evaluator (matrix as inputs, array of results as outputs)
  |   |      |
  |-- PUBLIC DIR (web interface)
  |   |
  |   |-- CSS DIR
  |   |      |-- style.css (styles)
  |   |      |
  |   |-- JS DIR
  |   |      |-- script.js (data validation and file uploading)
  |   |      |
  |   |-- index (interface startpoint)
  |   |
__\___\

- REST API developped with Golang (Routing and file interpreting)
- Very Simple Web interface with native HTML CSS ans JS
