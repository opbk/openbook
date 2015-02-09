GO = /usr/bin/go
BUILD_DIR = build
PROJECT = github.com/opbk/openbook
VERSION=$(shell cat version)

.PHONY: all build deb clear_deb test migrate_up migrate_down build_backend build_frontend

all: build

dependencies = code.google.com/p/gcfg \
github.com/lib/pq \
github.com/gorilla/schema \
github.com/gorilla/ \
github.com/gorilla/sessions \
github.com/gorilla/securecookie \
gopkg.in/check.v1 \
github.com/streadway/amqp \
github.com/cihub/seelog \
github.com/mattes/migrate \
github.com/jmoiron/sqlx \
github.com/astaxie/beego/orm \
github.com/goamz/goamz/aws \

dependencies_paths := $(addprefix $(GOPATH)/src/,$(dependencies))
$(dependencies_paths):
	for i in $(dependencies); do $(GO) get -d $$i; done

dependencies: $(dependencies_paths)

build_backend: dependencies
	rm -rf $(BUILD_DIR)/backend
	$(GO) build -o $(BUILD_DIR)/backend/usr/lib/openbook/backend/backend $(PROJECT)/backend
	mkdir -p $(BUILD_DIR)/backend/usr/lib/openbook/backend/templates
	mkdir -p $(BUILD_DIR)/backend/etc/openbook/backend/
	mkdir -p $(BUILD_DIR)/backend/etc/init.d/
	cp -r resources/backend/templates $(BUILD_DIR)/backend/usr/lib/openbook/backend/templates
	cp resources/config.gcfg $(BUILD_DIR)/backend/etc/openbook/backend/config.gcfg
	cp resources/seelog.xml $(BUILD_DIR)/backend/etc/openbook/backend/seelog.xml
	cp -r deb/backend/* $(BUILD_DIR)/backend/

deb_backend: build_backend
	fakeroot dpkg-deb --build build/backend backend_$(VERSION)_amd64.deb

build_frontend: dependencies
	rm -rf $(BUILD_DIR)/frontend
	$(GO) build -o $(BUILD_DIR)/frontend/usr/lib/openbook/frontend/frontend $(PROJECT)/frontend
	mkdir -p $(BUILD_DIR)/frontend/var/lib/openbook/frontend
	mkdir -p $(BUILD_DIR)/frontend/etc/openbook/frontend/
	mkdir -p $(BUILD_DIR)/frontend/etc/init.d/
	cp -r resources/frontend/templates $(BUILD_DIR)/frontend/usr/lib/openbook/frontend
	cp -r resources/frontend/static $(BUILD_DIR)/frontend/usr/lib/openbook/frontend
	cp resources/config.gcfg $(BUILD_DIR)/frontend/etc/openbook/frontend/config.gcfg
	cp resources/seelog.xml $(BUILD_DIR)/frontend/etc/openbook/frontend/seelog.xml
	cp -r deb/frontend/* $(BUILD_DIR)/frontend/

deb_frontend: build_frontend
	fakeroot dpkg-deb --build build/frontend frontend_$(VERSION)_amd64.deb

build: build_backend build_frontend

deb: deb_frontend deb_backend

test: 
	$(GO) test $(PROJECT)/common/model/book
	$(GO) test $(PROJECT)/common/model/book/category
	$(GO) test $(PROJECT)/common/model/book/category/utils
	$(GO) test $(PROJECT)/common/model/book/price
	$(GO) test $(PROJECT)/common/model/author
	$(GO) test $(PROJECT)/common/model/publisher
	$(GO) test $(PROJECT)/common/model/user
	$(GO) test $(PROJECT)/common/model/user/subscription
	$(GO) test $(PROJECT)/common/model/user/address
	$(GO) test $(PROJECT)/common/model/subscription
	$(GO) test $(PROJECT)/common/model/order
	$(GO) test $(PROJECT)/frontend
	$(GO) test $(PROJECT)/frontend/form
	$(GO) test $(PROJECT)/backend

migrate_up:
	-migrate -url postgres://developer:developer@localhost/openbook -path ./migrations/postgresql up

migrate_down:
	-migrate -url postgres://developer:developer@localhost/openbook -path ./migrations/postgresql down

migrate: migrate_down migrate_up

clear:
	rm -rf build

clear_deb:
	rm -f *.deb
