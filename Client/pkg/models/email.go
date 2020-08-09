package models

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const RECIPIENT = "p.polcik1@gmail.com"
const SENDER = "polcik.piotr@gmail.com"
const CHARSET = "UTF-8"

type Email struct {
	Sender    string
	Recipient string
	Subject   string
	HtmlBody  string
	TextBody  string
	CharSet   string
}

func (e *Email) CreateEmail(m *[]Basic) {
	e.Sender = SENDER
	e.Recipient = RECIPIENT
	e.Subject = "New Beds available on dba.dk"
	e.HtmlBody = createHtmlBody(m)
	e.TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."
	e.CharSet = CHARSET
}

func createHtmlBody(m *[]Basic) string {
	htmlBody := ""
	for _, val := range *m {
		htmlBody = htmlBody + `<div style='border: 1px solid #fff; padding: 10px; margin-bottom: 10px;'>
									<h4>` + val.Title + `</h4>
									<p>Date: ` + translate(val.Date) + `</p>
									<p>Price: ` + strconv.Itoa(val.Price) + `</p>
									<a href='` + val.URL + `'>link</a>
								</div>`
	}

	return htmlBody
}

func translate(str string) string {
	if str == "I dag" {
		return "Today"
	} else if str == "I går" {
		return "Yesterday"
	}
	return str
}

func (e *Email) Send() {
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(e.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(e.CharSet),
					Data:    aws.String(e.HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(e.CharSet),
					Data:    aws.String(e.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(e.CharSet),
				Data:    aws.String(e.Subject),
			},
		},
		Source: aws.String(e.Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
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

	fmt.Println("Email Sent.")
	fmt.Println(result)
}
