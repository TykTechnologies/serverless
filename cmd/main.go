package main

import (
	"log"

	"github.com/TykTechnologies/serverless/provider"
	"github.com/TykTechnologies/serverless/provider/aws"
	_ "github.com/TykTechnologies/serverless/provider/azure"
)

func main() {

	lambda, err := provider.GetProvider("aws-lambda")
	if err != nil {
		log.Fatal(err)
	}

	conf := &aws.AWSConf{
		Region: "us-east-1",
	}

	if err := lambda.Init(conf); err != nil {
		log.Fatal(err)
	}

	detail, err := lambda.List()
	if err != nil {
		log.Fatal(err)
	}

	f := provider.Function{
		Name: "helloWorld",
	}

	resp, err := lambda.Invoke(f, []byte{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", detail)
	log.Println(resp)
}
