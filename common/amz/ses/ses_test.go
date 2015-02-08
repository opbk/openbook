package ses

import (
	"testing"

	"github.com/goamz/goamz/aws"
	"gopkg.in/check.v1"

	"git.2rll.net/newsgun/common/configuration"
)

func Test(t *testing.T) { check.TestingT(t) }

type TestSuite struct{}

var _ = check.Suite(new(TestSuite))

func (s *TestSuite) SetUpSuite(c *check.C) {
	configuration.LoadConfiguration("")
}

func (s *TestSuite) TestSend(c *check.C) {
	body := "From: \"Небеса\" <noreply-mailgun-test@2rll.net>\r\n" +
		"To: <bounce@simulator.amazonses.com>\r\n" +
		"Subject: \"This is a subject\"\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"List-Unsubscribe: <http://unsubscribe.email.heavens.com/?p=1,1&direct>\r\n" +
		"Precedence: bulk\r\n" +
		"Date: Wed, 01 Jan 2014 10:00:00 +0000\r\n" +
		"Message-Id: <ajsf4jfypds0@email.heavens.com>\r\n\r\n" +
		"PGRpdiBiZ2NvbG9yPSIjZmZmZmZmIiBtYXJnaW5oZWlnaHQ9IjAiIG1hcmdpbndpZHRoPSIwIiBzdHlsZT0id2lkdGg6MT\n" +
		"AwJSFpbXBvcnRhbnQ7YmFja2dyb3VuZDojZmZmZmZmIj4KICA8dGFibGUgd2lkdGg9IjYwMCIgYm9yZGVyPSIwIiBjZWxs\n" +
		"c3BhY2luZz0iMCIgY2VsbHBhZGRpbmc9IjAiIGFsaWduPSJjZW50ZXIiIGJnY29sb3I9IiNmZmZmZmYiPgogICAgPHRib2\n" +
		"R5Pjx0cj48dGQ+CiAgICAgIFRoaXMgaXMgYSBlbWFpbCBib2R5CiAgICA8L3RkPjwvdHI+PC90Ym9keT4KICA8L3RhYmxl\n" +
		"PgogIDxpbWcgc3JjPSJodHRwOi8vdmlldy5lbWFpbC5oZWF2ZW5zLmNvbS8/cD0xLDEiIHdpZHRoPSIxIiBoZWlnaHQ9Ij\n" +
		"EiIC8+CjwvZGl2Pgo="

	auth := aws.Auth{
		AccessKey: configuration.GetConfig().Aws.AccessKey,
		SecretKey: configuration.GetConfig().Aws.SecretKey,
	}

	ses := NewSESWrapper(auth, aws.Regions[configuration.GetConfig().Aws.Region])
	err := ses.SendRawEmail("bounce@simulator.amazonses.com", body)
	c.Assert(err, check.IsNil)
}

func (s *TestSuite) TestVerifyEmail(c *check.C) {
	auth := aws.Auth{
		AccessKey: configuration.GetConfig().Aws.AccessKey,
		SecretKey: configuration.GetConfig().Aws.SecretKey,
	}

	ses := NewSESWrapper(auth, aws.Regions[configuration.GetConfig().Aws.Region])
	err := ses.VerifyEmail("noreply-mailgun-test@2rll.net")
	c.Assert(err, check.IsNil)

	err = ses.SetNotificationTopic("noreply-mailgun-test@2rll.net", "Bounce", configuration.GetConfig().Aws.SNSBounceTopic)
	c.Assert(err, check.IsNil)

	err = ses.EnableSNSNotification("noreply-mailgun-test@2rll.net")
	c.Assert(err, check.IsNil)

	err = ses.DeleteEmail("noreply-mailgun-test@2rll.net")
	c.Assert(err, check.IsNil)
}
