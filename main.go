package main

import (
	"net/http"
	logger "github.com/inconshreveable/log15"
	"io/ioutil"
	"io"
	"os"
)

type credentials struct {
	Username string
	Password string
}

var (
	creds credentials
	requestURL = "http://localhost:8088/message/retort.wav"
	tempFilePath = "voicemails/"
)

func main() {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		logger.Error("Unable to create request", "error", err)
		return
	}
	creds.Username = "a_username"
	creds.Password = "a_password"
	req.SetBasicAuth(creds.Username, creds.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Failed to make authenticated request", "error", err)
		return
	}
	if resp.StatusCode != 200 {
		logger.Error("status code error", "statuscode", resp.StatusCode)
	}
	logger.Debug("client http response", "raw resp", resp)
	defer resp.Body.Close()

	//create temporary file to store mailbox message audio file
	wavFile, err := ioutil.TempFile(tempFilePath, "msg-*.wav")
	logger.Debug("Created temp file", "filename", wavFile.Name())
	if err != nil {
		logger.Error("Failed to create temp file", "error", err)
		return
	}
	defer os.Remove(wavFile.Name())
	var numWritten int64
	if numWritten, err = io.Copy(wavFile, resp.Body); err != nil {
		logger.Error("Failed to copy mailbox message file to tmp file", "error", err)
		return
	}
	logger.Debug("copied length from resp body", "bytes written", numWritten)

	// TODO: can delete this Stat call and just call .Name() on the file object for same result
	fileInfo, err := wavFile.Stat()
	if err != nil {
		logger.Error("unable to get stats of file: %v", err)
		return
	}
	logger.Debug("File info", "file info", fileInfo)

}
