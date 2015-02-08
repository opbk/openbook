package ses

import (
	"encoding/base64"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/goamz/goamz/aws"
)

type SESWrapper struct {
	auth     aws.Auth
	endpoint string
	signer   *aws.Route53Signer
}

func NewSESWrapper(auth aws.Auth, regioin aws.Region) *SESWrapper {
	return &SESWrapper{
		auth:     auth,
		endpoint: regioin.SESEndpoint,
		signer:   aws.NewRoute53Signer(auth),
	}
}

func (ses *SESWrapper) SendRawEmail(to, rawBody string) error {
	data := make(map[string]string)
	data["Action"] = "SendRawEmail"
	data["Destinations.member.1"] = to
	data["RawMessage.Data"] = base64.StdEncoding.EncodeToString([]byte(rawBody))

	_, err := ses.query("POST", "/", data)
	return err
}

func (ses *SESWrapper) VerifyEmail(email string) error {
	data := make(map[string]string)
	data["Action"] = "VerifyEmailIdentity"
	data["EmailAddress"] = email

	_, err := ses.query("POST", "/", data)
	return err
}

func (ses *SESWrapper) DeleteEmail(email string) error {
	data := make(map[string]string)
	data["Action"] = "DeleteIdentity"
	data["Identity"] = email

	_, err := ses.query("POST", "/", data)
	return err
}

func (ses *SESWrapper) EnableSNSNotification(email string) error {
	data := make(map[string]string)
	data["Action"] = "SetIdentityFeedbackForwardingEnabled"
	data["Identity"] = email
	data["ForwardingEnabled"] = "false"

	_, err := ses.query("POST", "/", data)
	return err
}

func (ses *SESWrapper) SetNotificationTopic(email, nType, snsTopic string) error {
	data := make(map[string]string)
	data["Action"] = "SetIdentityNotificationTopic"
	data["Identity"] = email
	data["NotificationType"] = nType
	data["SnsTopic"] = snsTopic

	_, err := ses.query("POST", "/", data)
	return err
}

func (ses *SESWrapper) query(method, path string, params map[string]string) (content string, err error) {
	params["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	u, err := url.Parse(ses.endpoint)
	if err != nil {
		return "", err
	}

	u.Path = path
	if method == "GET" {
		u.RawQuery = multimap(params).Encode()
	}

	reqBody := multimap(params).Encode()
	req, _ := http.NewRequest(method, u.String(), strings.NewReader(reqBody))
	ses.signer.Sign(req)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

	_, err = httputil.DumpRequest(req, true)
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode > 204 {
		return "", buildError(res)
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	return string(resBody), err
}

func buildError(r *http.Response) error {
	errors := aws.ErrorResponse{}
	xml.NewDecoder(r.Body).Decode(&errors)
	var err aws.Error
	err = errors.Errors
	err.RequestId = errors.RequestId
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}
