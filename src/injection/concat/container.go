package main

type Logger struct{}

type HttpClient struct {
	logger *Logger
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{logger}
}

func NewConcatService(logger *Logger, client *HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}

type ConcatService struct {
	logger *Logger
	client *HttpClient
}

func CreateConcatService() *ConcatService {
	logger := &Logger{}
	client := NewHttpClient(logger)
	return NewConcatService(logger, client)
}
