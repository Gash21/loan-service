package mailer

func NewMailerAdapter(opts *MailerAdapter) IMailerAdapter {
	if opts != nil && opts.Type == "smtp" {
		return &MailerAdapter{
			Host:     opts.Host,
			Port:     opts.Port,
			Username: opts.Username,
			Password: opts.Password,
			Sender:   opts.Sender,
		}
	}
	return &MailerAdapter{}
}
