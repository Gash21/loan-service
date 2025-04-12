package mailer

type MailerAdapter struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Sender   string
	TLS      bool
}

type IMailerAdapter interface {
	SendEmail(opts EmailOpts) error
	PrepareHtmlBody(templateStr string, params map[string]interface{}) (string, error)
}

type EmailOpts struct {
	To      []string
	Subject string
	Body    []map[string]interface{}
	Headers map[string]string
}
