# About Me
 This is a potential response to Fetch Rewards' backend interview project called receipt-processor-challenge. This project was written in Golang 1.19 and utilized the Gin HTTP web framework v1.8.2. The default port is `8080`. It contains two endpoints: a POST method with a path of  `/receipts/process` and a GET method with a path of `/receipts/:id/points`. The POST method will accept a Receipt JSON, generate an ID, calculate award points, save data to makeshift databases (maps), and return a JSON containing an id for that receipt. The GET method will take in a parameter id, access the database where points are stored using that id, and return a JSON object containing the number of points awarded.

 
 # Getting Started
 ## Requirements
 - Any Integrated Development Environment (IDE). [Microsoft Visual Studio Code](https://code.visualstudio.com/download) (VS CODE) IDE was used during the completion of this project.
 
 - An installation of Go 1.19 or later. For installation instructions, see [Downloading and Installing Go](https://go.dev/doc/install).
 
 - A command terminal. VS Code allows access to a command terminal within the IDE.
 
 ## Recommended 
 - An API platform such as [Postman](https://www.postman.com/downloads/). 
 - Curl commands within a command terminal can be used in lieu of this platform. No curl commands are provided.
 
 ## Start Up
 - After the repository has been pulled and established on the local computer.
 - Open a command terminal. 
 - Use the `cd` command to change directories to `ReceiptProcessorChallenge`. 
 - At the command prompt type `go run .` or `go run main.go` to startup the HTTP server on `localhost:8080`.
 
 ## How to Use Postman
 - Open the API platform Postman Desktop and create a new workspace using the `Workspaces` dropdown menu located at the top left-hand corner of the window.
 - Open a new tab using the `+` button located near the top middle of the workspace.
 - Using the method type dropdown that is defaulted to GET in the newly opened tab, select `POST`
 - In the `Enter request URL` field type `localhost:8080/receipts/process
 - Select `Body` and then the `raw` radio button
 - Paste JSON receipt to be uploaded within textbox
 - Press `Send`
 - A response with a JSON object will be returned. Copy the id value (not including the `""`).
 - Open a new tab using the `+` button located near the top middle of the workspace.
 - In the `Enter request URL` field type `localhost:8080/receipts/` then paste the id and continue typing `/points`
 - Press `Send`
 - A response with a JSON object will be returned.

# Reflection
This project was written in a language in which I had no prior knowledge. There were some challenges with syntax along the way but they were overcome utilizing the official [Go Documentation](https://go.dev/doc/). The project code could have been implemented more effeciently. For instance, the Receipt.go file could be broken apart into multiple files containing the different structs and their accompanying functions. A different web framework or package could be used in lieu of Gin. Gin was used because there is a helpful [Go Tuturial](https://go.dev/doc/tutorial/web-service-gin) located in the Go Documentation that assisted in building a RESTful web service API. Gin is fast, crash-free, and has methods for JSON validation. In the future this project could be adapted to work with an actual database and forego the use of maps as makeshift databases. Overall this was a fun project and provided me with the opportunity to learn a new language. Thank you for your time and consideration.
