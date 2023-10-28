package awsdk

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"report-transaction/internal/env"
)

var sess *session.Session

func Init() error {
	var err error

	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(env.AwsRegion),
		Credentials: credentials.NewStaticCredentials(
			env.AwsAccessKeyId,
			env.AwsSecretAccessKey,
			env.AwsCredentialToken,
		),
	})

	return err
}

func s3Client() *s3.S3 {
	return s3.New(sess)
}

func PutObject(key string, content []byte) error {
	_, err := s3Client().PutObject(&s3.PutObjectInput{
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
