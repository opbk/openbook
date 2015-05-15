package webtest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/check.v1"

	"github.com/opbk/openbook/common/db"
	"github.com/opbk/openbook/common/model/user"
	"github.com/opbk/openbook/common/web/auth"
)

type UserControllerTestSuit struct {
	ControllerTestSuit
	authWrapper *AuthWrapper
}

func (s *UserControllerTestSuit) SetUpSuite(c *check.C, router *mux.Router) {
	s.authWrapper = NewAuthWrapper(s.GetRouter())
	s.ControllerTestSuit.SetUpSuite(c)

	db.Connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	db.Connection().Exec("TRUNCATE users")
	(&user.User{0, "netw00rk@gmail.com", "123456", "A. Koklin", "89161234567", time.Now(), time.Now(), time.Now()}).Save()
}

func (s *UserControllerTestSuit) TearDownSuite(c *check.C) {
	db.Connection().Exec("ALTER SEQUENCE auto_id_users RESTART WITH 1")
	db.Connection().Exec("TRUNCATE users")

	s.ControllerTestSuit.TearDownSuite(c)
}

func (s *UserControllerTestSuit) WithUser(u *user.User) *UserControllerTestSuit {
	s.authWrapper.WithUser(u)
	return s
}

type AuthWrapper struct {
	router *mux.Router
	user   *user.User
}

func NewAuthWrapper(router *mux.Router) *AuthWrapper {
	return &AuthWrapper{router: router}
}

func (w *AuthWrapper) WithUser(u *user.User) {
	w.user = u
}

func (w *AuthWrapper) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if w.user != nil {
		auth.Set(w.user, rw, req)
	}

	w.router.ServeHTTP(rw, req)
	w.user = nil
}
