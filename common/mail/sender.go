package mail

import (
	"encoding/base64"
	"fmt"
	"net/mail"

	logger "github.com/cihub/seelog"
	"github.com/goamz/goamz/aws"

	"github.com/opbk/openbook/common/amz/ses"
	"github.com/opbk/openbook/common/configuration"
)

var mailSender *MailSender

type MailSender struct {
	ses    *ses.SESWrapper
	helper helper
	from   string
}

type Header struct {
	Name  string
	Value string
}

func InitMailSender(config *configuration.Config) {
	auth := aws.Auth{
		AccessKey: config.Aws.AccessKey,
		SecretKey: config.Aws.SecretKey,
	}

	mailSender = new(MailSender)
	mailSender.from = config.EmailSender.From
	mailSender.helper = new(helperImpl)
	mailSender.ses = ses.NewSESWrapper(auth, aws.Regions[config.Aws.Region])
}

func (ms *MailSender) SendTo(to, subject, body string) {
	toAddr := mail.Address{"", to}
	fromAddr := mail.Address{"OpenBook", ms.from}

	headers := []Header{
		{"To", toAddr.String()},
		{"From", fromAddr.String()},
		{"Subject", ms.helper.encodeRFC2047(subject)},
		{"MIME-Version", "1.0"},
		{"Content-Type", "text/html; charset=\"utf-8\""},
		{"Content-Transfer-Encoding", "base64"},
		{"List-Unsubscribe", "http://opbook.org/unsubrcibe"},
		{"Date", ms.helper.getLocalTime().Format("Mon, 02 Jan 2006 15:04:05 -0700")},
		{"Message-Id", fmt.Sprintf("<%s@%s>", ms.helper.genLocalPartMessageId(), "opbook.org")},
	}

	var emailBody string
	for _, header := range headers {
		emailBody += fmt.Sprintf("%s: %s\r\n", header.Name, header.Value)
	}
	emailBody += "\r\n"

	body = base64.StdEncoding.EncodeToString([]byte(body))
	for i := 0; i < len(body); i += 94 {
		if i+95 > len(body) {
			emailBody += body[i:]
		} else {
			emailBody += body[i:i+94] + "\n"
		}
	}

	err := ms.ses.SendRawEmail(to, emailBody)
	if err != nil {
		logger.Errorf("Error while sending email %s", err)
	}
}

func GetMailSender() *MailSender {
	return mailSender
}

func SendTo(to, subject, body string) {
	GetMailSender().SendTo(to, subject, body)
}
