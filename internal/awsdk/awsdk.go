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

const charSet = "UTF-8"

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

func SendEmail(body string, subject string, recipients []string) error {
	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toAddresses(recipients),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(body),
				},
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(env.ServiceEmail),
	})

	return err
}

func toAddresses(recipients []string) []*string {
	var addresses []*string
	for _, recipient := range recipients {
		addresses = append(addresses, aws.String(recipient))
	}
	return addresses
}
