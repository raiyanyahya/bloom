package cloud

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Object struct {
}

func (s S3Object) Client() interface{} {
	session, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})
	if err != nil {
		log.Fatalf("could not initialize new aws session: %v", err)
	}
	return s3.New(session)
}
func GetClient(objectstore ObjectStore) interface{} {
	return objectstore.Client()
}
