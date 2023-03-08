package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	awsv2 "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	credentialsv2 "github.com/aws/aws-sdk-go-v2/credentials"
	stscredsv2 "github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	stsv2 "github.com/aws/aws-sdk-go-v2/service/sts"
	color "github.com/logrusorgru/aurora"
)

func printUsage() {
	fmt.Println(color.Bold(color.Cyan("usage:")))
	fmt.Println(color.Bold(color.BrightWhite("--client-id={clientID}")))
	fmt.Println(color.Bold(color.BrightWhite("--client-secret={clientSecret}")))
	fmt.Println(color.Bold(color.BrightWhite("--role-arn={ arn:aws:iam::728806012345:role/some-role-to-assume}")))
	fmt.Println(color.Bold(color.BrightWhite("--region={some aws valid region..defaults to us-east-1 if not used}")))
	fmt.Println(color.Bold(color.BrightWhite("--credentials-file-path=/")))
	fmt.Println(color.Bold(color.Yellow("help: { list this message :-) }")))

}

/*
func createFolderFile() {

	e := ioutil.WriteFile(user.HomeDir+"/.escli/config.json", credsFile, 0644)
	if e != nil {
		panic(e)
	}
	fmt.Println("config file written..")

}
*/

func main() {
	clientID := flag.String("client-id", "ASXXXTMCDDWMDYYY6ZZZ", "ClienID of the IAM User")
	clientSecret := flag.String("client-secret", "f345j13hrl3jkff3kjfq;l3j34t", "Client Secret")
	roleArn := flag.String("role-arn", "null", "Role ARN of the role you want to assume")
	awsRegion := flag.String("region", "us-east-1", "AWS Region")
	credentialsFilePath := flag.String("credentials-file-path", "./", "Credentials file path")
	flag.Parse()

	if len(os.Args) == 1 {
		printUsage()
		os.Exit(1)
	}

	_, err := os.Stat(string(*credentialsFilePath))
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(*credentialsFilePath, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	ctx := context.Background()
	assumecnf, _ := config.LoadDefaultConfig(
		ctx, config.WithRegion(string(*awsRegion)),
		config.WithCredentialsProvider(awsv2.NewCredentialsCache(
			credentialsv2.NewStaticCredentialsProvider(
				string(*clientID),
				string(*clientSecret), "",
			)),
		),
	)

	stsclient := stsv2.NewFromConfig(assumecnf)

	cnf, _ := config.LoadDefaultConfig(
		ctx, config.WithRegion(string(*awsRegion)),
		config.WithCredentialsProvider(awsv2.NewCredentialsCache(
			stscredsv2.NewAssumeRoleProvider(
				stsclient,
				string(*roleArn),
				func(o *stscredsv2.AssumeRoleOptions) {
					o.RoleARN = *roleArn
				},
			)),
		),
	)
	creds, err := cnf.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AWS Assume Role Credentials")
	fmt.Println("Expires :", creds.Expires)
	fmt.Println()
	fmt.Println("[default]")
	fmt.Println("aws_access_key_id =", creds.AccessKeyID)
	fmt.Println("aws_secret_access_key =", creds.SecretAccessKey)
	fmt.Println("aws_session_token =", creds.SessionToken)
	fil, _ := os.Create(*credentialsFilePath + "/credentials")
	wr := bufio.NewWriter(fil)
	wr.WriteString("[default]\n")
	wr.WriteString("aws_access_key_id = " + creds.AccessKeyID + "\n")
	wr.WriteString("aws_secret_access_key = " + creds.SecretAccessKey + "\n")
	wr.WriteString("aws_session_token = " + creds.SessionToken + "\n")
	wr.Flush()

}
