package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func main() {
	statusInput := flag.String("status", "Active", "'Active' or 'Inactive' status for AWS Access Key")
	flag.Parse()

	mySession := session.Must(session.NewSession())

	// Create a IAM client from just a session.
	svc := iam.New(mySession)

	// Get List of IAM Users
	users, err := svc.ListUsers(nil)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	// Loop through list of users and get their AWS Access Keys
	for _, user := range users.Users {
		input := &iam.ListAccessKeysInput{
			UserName: aws.String(*user.UserName),
		}

		lak, err := svc.ListAccessKeys(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case iam.ErrCodeNoSuchEntityException:
					fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
				case iam.ErrCodeServiceFailureException:
					fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		// Loop through all Access Keys
		for _, ak := range lak.AccessKeyMetadata {
			fmt.Printf("Old: %s, %s, %s\n", *ak.UserName, *ak.AccessKeyId, *ak.Status)
			// Set Access Keys to input flag value
			input := &iam.UpdateAccessKeyInput{
				AccessKeyId: ak.AccessKeyId,
				Status:      aws.String(*statusInput),
				UserName:    ak.UserName,
			}
			_, err := svc.UpdateAccessKey(input)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case iam.ErrCodeNoSuchEntityException:
						fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
					case iam.ErrCodeServiceFailureException:
						fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
					default:
						fmt.Println(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
				}
				return
			}
		}
	}

	// Loop through list of users and get their AWS Access Keys
	for _, user := range users.Users {
		input := &iam.ListAccessKeysInput{
			UserName: aws.String(*user.UserName),
		}

		lak, err := svc.ListAccessKeys(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case iam.ErrCodeNoSuchEntityException:
					fmt.Println(iam.ErrCodeNoSuchEntityException, aerr.Error())
				case iam.ErrCodeServiceFailureException:
					fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		// Loop through all Access Keys
		for _, ak := range lak.AccessKeyMetadata {
			// Print the Access Key and Status
			fmt.Printf("New: %s, %s, %s\n", *ak.UserName, *ak.AccessKeyId, *ak.Status)
		}
	}
}
