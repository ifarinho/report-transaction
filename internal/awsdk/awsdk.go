package awsdk

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"net/http"
	"report-transaction/internal/env"
)

var (
	awsSession *session.Session
	s3Client   *s3.S3
	sesClient  *ses.SES
)

func Init() error {
	var err error

	awsSession, err = session.NewSession(&aws.Config{
		Region: aws.String(env.AwsRegion),
		Credentials: credentials.NewStaticCredentials(
			env.AwsAccessKeyId,
			env.AwsSecretAccessKey,
			env.AwsCredentialToken,
		),
	})
	if err != nil {
		return err
	}

	s3Client = s3.New(awsSession)

	sesClient = ses.New(awsSession)

	return nil
}

func PutObject(content []byte, key string) error {
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(env.AwsS3Bucket),
		Key:                  aws.String(fmt.Sprintf("%s/%s", env.AwsFullPath, key)),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(content),
		ContentLength:        aws.Int64(int64(len(content))),
		ContentType:          aws.String(http.DetectContentType(content)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}

func SendEmail(content []byte, destinations []string) error {
	_, err := sesClient.SendRawEmail(&ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: content,
		},
		Destinations: aws.StringSlice(destinations),
		Source:       aws.String(env.ServiceEmail),
	})

	return err
}
