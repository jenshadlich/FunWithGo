package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
)

var (
    Bucket = "jns-test"
)

func main() {
    
    client := s3.New(&aws.Config{
        Region: "eu-central-1",
        Endpoint: "s3.amazonaws.com",
    })
    params := &s3.ListObjectsInput{Bucket: &Bucket}

    objects, err := client.ListObjects(params)
    if err != nil {
        panic(err)
    }

    for _, object := range objects.Contents {
        fmt.Printf("%s\n", *object.Key)
    }

}
