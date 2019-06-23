package s3

import (
	"github.com/copybird/copybird/modules/backup/output"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const MODULE_NAME = "s3"

type S3 struct {
	output.Output
	reader  io.Reader
	writer  io.Writer
	session *session.Session
	config  *Config
}

func (s *S3) GetName() string {
	return MODULE_NAME
}

func (s *S3) GetConfig() interface{} {
	return &Config{}
}

func (s *S3) InitPipe(w io.Writer, r io.Reader) error {
	s.reader = r
	s.writer = w
	return nil
}

func (s *S3) InitModule(_config interface{}) error {
	s.config = _config.(*Config)
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(s.config.Region),
		Credentials: credentials.NewStaticCredentials(s.config.AccessKeyID, s.config.SecretAccessKey, ""),
	})
	if err != nil {
		return err
	}

	s.session = session
	return nil
}

func (s *S3) Run() error {

	svc := s3manager.NewUploader(s.session)

	input := &s3manager.UploadInput{
		Bucket: aws.String(s.config.Bucket),
		Key:    aws.String(s.config.FileName),
		Body:   s.reader,
	}

	_, err := svc.Upload(input)
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) Close() error {
	return nil
}