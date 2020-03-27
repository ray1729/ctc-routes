# CTC Routes

Quick and dirty script to generate Java code for routes database from a spreadsheet.

## Usage

Export the simplified route sheet from Google Sheets (File > Download > Comma-separated values).

    go run main.go "Routes - Sheet2.csv" routes.java

If you prefer, you can install the binary:

    go get github.com/ray1729/ctc-routes
    go install github.com/ray1729/ctc-routes
    
Then

    ${GOPATH}/bin/ctc-routes "Routes - Sheet2.csv" routes.java
