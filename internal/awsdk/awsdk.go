package awsdk

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"io"
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

func GetObject(key string) ([]byte, error) {
	result, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(env.AwsS3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(result.Body)
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
