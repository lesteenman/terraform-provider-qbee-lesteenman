package qbee

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
)

type FilesService struct {
	Client *HttpClient
}

type FileUploadResponse struct {
	File string `json:"file"`
	Path string `json:"path"`
}

func (s *FilesService) Upload(sourceFile string, targetPath string, filename string) (*FileUploadResponse, error) {
	file, errFile1 := os.Open(sourceFile)
	if errFile1 != nil {
		return nil, fmt.Errorf("could not open source file '%v': %w", sourceFile, errFile1)
	}
	defer file.Close()

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	err := writer.WriteField("path", targetPath)
	if err != nil {
		return nil, fmt.Errorf("could not write multipart field 'path': %w", err)
	}

	if filename == "" {
		filename = filepath.Base(file.Name())
	}

	part, _ := writer.CreateFormFile("file", filename)
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("could not copy bytes to form part: %w", err)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.UploadFile("/file", payload, writer.FormDataContentType())
	if err != nil {
		return nil, fmt.Errorf("form Post to /file failed: %w", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("error response code %v after file Upload (resp. body = %v)", resp.StatusCode, string(b))
	}

	var uploadResponse FileUploadResponse
	err = json.Unmarshal(b, &uploadResponse)
	if err != nil {
		return nil, fmt.Errorf("could not parse reponse from POST /files, was '%v': %w", string(b), err)
	}

	return &uploadResponse, nil
}

type DownloadOptions struct {
	Path string `url:"path,omitempty"`
}

type FileDownloadResponse struct {
	Contents string
}

func (s *FilesService) Download(path string) (*FileDownloadResponse, error) {
	query := DownloadOptions{Path: path}
	resp, err := s.Client.Get("/file", query)
	if err != nil {
		return nil, fmt.Errorf("files.Download(%v): %w", path, err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("files.Download io.ReadAll: %w", err)
	}

	return &FileDownloadResponse{Contents: string(b)}, nil
}

type ListFilesResponse struct {
	Items []FileMetadata `json:"items"`
}

type FileMetadata struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Path      string `json:"path"`
	IsDir     bool   `json:"is_dir"`
	Created   uint64 `json:"created"`
	Mime      string `json:"mime"`
	Size      uint64 `json:"size"`
	Digest    string `json:"digest"`
}

func (s *FilesService) List() (*ListFilesResponse, error) {
	resp, err := s.Client.Get("/files", nil)
	if err != nil {
		log.Printf("Err in Client.Get: %v", err)
		return nil, fmt.Errorf("files.ListFiles: %w", err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Err in io.ReadAll: %v", err)
		return nil, fmt.Errorf("files.Download io.ReadAll: %w", err)
	}

	var l ListFilesResponse
	err = json.Unmarshal(b, &l)
	if err != nil {
		log.Printf("could not parse json: %v\n", string(b))
		return nil, fmt.Errorf("files.Download Unmarshal: %w", err)
	}

	return &l, nil
}

type getFileMetadata struct {
	Path string `url:"path"`
}

var ErrFileNotFound = errors.New("file not found")

func (s *FilesService) GetMetadata(path string) (*FileMetadata, error) {
	r, err := s.Client.Get("/file/metadata", getFileMetadata{
		Path: path,
	})

	if err != nil {
		log.Printf("Err in Client.Get: %v", err)

		var re QbeeHttpClientError
		if errors.As(err, &re) {
			var apiErr = re.ApiError
			log.Printf("Error is a valid QbeeApiError: code=%v, message=%v", apiErr.Code, apiErr.Message)
			if apiErr.Code == 404 && apiErr.Message == "File not found" {
				return nil, ErrFileNotFound
			}
		} else {
			log.Printf("Could not cast err %+v (typeof=%v, kind=%v) to a QbeeHttpClientError", err, reflect.TypeOf(err), reflect.ValueOf(err).Kind())
		}

		return nil, fmt.Errorf("files.GetMetadata: %w", err)
	}

	var response FileMetadata

	err = s.Client.ParseJsonBody(r, &response)
	if err != nil {
		return nil, fmt.Errorf("files.GetMetadata(%v): %w", path, err)
	}

	return &response, nil
}

type DeleteOptions struct {
	Path string `json:"path"`
}

func (s *FilesService) Delete(path string) error {
	opt := DeleteOptions{Path: path}
	_, err := s.Client.Delete("/file", opt)
	if err != nil {
		return fmt.Errorf("files.Delete(%v): %w", path, err)
	}

	return nil
}

type CreateDirOptions struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type CreateDirResponse struct {
	Dir  string `json:"dir"`
	Path string `json:"path"`
}

func (s *FilesService) CreateDir(path string, dirName string) (response *CreateDirResponse, err error) {
	opt := CreateDirOptions{
		Path: path,
		Name: dirName,
	}

	if path == "" {
		return nil, fmt.Errorf("files.CreateDir(%v, %v): empty path given", path, dirName)
	}

	if dirName == "" {
		return nil, fmt.Errorf("files.CreateDir(%v, %v): empty dirName given", path, dirName)
	}

	r, err := s.Client.Post("/file/createdir", opt)
	if err != nil {
		return nil, fmt.Errorf("files.CreateDir(%v, %v): %w", path, dirName, err)
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("files.CreateDir(%v, %v) io.ReadAll: %w", path, dirName, err)
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, fmt.Errorf("files.CreateDir Unmarshal: %w", err)
	}

	return response, nil
}
