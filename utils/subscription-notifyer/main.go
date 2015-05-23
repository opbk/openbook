package main

import (
	"bytes"
	"flag"
	"html/template"
	"path"
	"runtime"
	"sync"
	"time"

	logger "github.com/cihub/seelog"

	"github.com/opbk/openbook/common/configuration"
	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/mail"
	"github.com/opbk/openbook/common/model/book"
	"github.com/opbk/openbook/common/model/order"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/model/user/subscription"
)

func initLogging(config *configuration.Config) {
	newLogger, _ := logger.LoggerFromConfigAsFile(config.Main.LogFile)
	logger.ReplaceLogger(newLogger)
}

func initDataBases(config *configuration.Config) {
	db.InitDbConnection(config.Db.Driver, config.Db.Connection)
}

func run() {
	var wg sync.WaitGroup
	tPath := configuration.GetConfig().Frontend.TemplatePath
	t := template.Must(template.New("index").Delims("{%", "%}").ParseFiles(
		path.Join(tPath, "email", "footer.html"),
		path.Join(tPath, "email", "header.html"),
		path.Join(tPath, "email", "notify.html"),
	))

	for _, usub := range subscription.List() {
		sy, sm, sd := time.Now().AddDate(0, 0, 7).Date()
		uy, um, ud := usub.Expiration.Date()
		if sy == uy && sm == um && sd == ud {
			u := user.Find(usub.UserId)
			var o *order.Order
			var b *book.Book
			var returnDate time.Time
			if order.CountByUserAndStatus(u.Id, order.ONHAND) > 0 {
				o = order.ListByUserAndStatusWithLimit(u.Id, order.ONHAND, 1, 0)[0]
				b = book.Find(o.BookId)
				returnDate = usub.Expiration.AddDate(0, 0, 14)
			}

			wg.Add(1)
			go func() {
				defer wg.Done()

				body := bytes.NewBuffer([]byte{})
				t.ExecuteTemplate(body, "email_notify", map[string]interface{}{
					"user":         u,
					"order":        o,
					"book":         b,
					"returnDate":   returnDate,
					"subscription": usub,
					"domain":       configuration.GetConfig().Main.Domain,
				})

				logger.Infof("Sending notifying to user %d", u.Id)
				mail.SendTo(u.Email, "Ваша подписка скоро закончится!", body.String())
			}()
		}
	}

	wg.Wait()
}

func main() {
	var configFile *string = flag.String("config", "/etc/openbook/frontend/config.gcfg", "configuration file")
	flag.Parse()

	config := configuration.LoadConfiguration(*configFile)
	runtime.GOMAXPROCS(config.Main.MaxProc)
	initLogging(config)
	initDataBases(config)
	mail.InitMailSender(config)

	run()
}
