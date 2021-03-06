package awsses

// Config for sending a Message to an Email Address in Amazon SES
type Config struct {
	Region    string // AWS Region for Amazon SES
	Sender    string // This address must be verified with Amazon SES.
	Recipient string // If your account is still in the sandbox, this address must be verified.
	Subject   string // The subject line for the email.
	HTMLbody  string // The HTML body for the email.
	Textbody  string // The email body for recipients with non-HTML email clients.
	Charset   string // The character encoding for the email.
}
