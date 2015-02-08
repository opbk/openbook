package mail

import (
	"net/mail"
	"strconv"
	"strings"
	"time"
)

type helper interface {
	encodeRFC2047(String string) string
	getLocalTime() time.Time
	genLocalPartMessageId() string
}

type helperImpl int

func (h *helperImpl) encodeRFC2047(String string) string {
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func (h *helperImpl) genLocalPartMessageId() string {
	return strconv.FormatInt(h.getLocalTime().UnixNano(), 36)
}

func (h *helperImpl) getLocalTime() time.Time {
	return time.Now()
}
