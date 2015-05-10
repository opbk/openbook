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

######## сборка frontend ########

build_frontend: dependencies
	rm -rf $(BUILD_DIR)/frontend
	$(GO) build -o $(BUILD_DIR)/frontend/usr/lib/openbook/frontend/frontend $(PROJECT)/frontend
	mkdir -p $(BUILD_DIR)/frontend/var/lib/openbook/frontend
	mkdir -p $(BUILD_DIR)/frontend/etc/openbook/frontend/
	mkdir -p $(BUILD_DIR)/frontend/etc/init.d/
	cp -r resources/frontend/templates $(BUILD_DIR)/frontend/usr/lib/openbook/frontend
	cp -r resources/frontend/static $(BUILD_DIR)/frontend/usr/lib/openbook/frontend
	cp resources/configuration/config.example.gcfg $(BUILD_DIR)/frontend/etc/openbook/frontend/config.gcfg
	cp resources/configuration/seelog.example.xml $(BUILD_DIR)/frontend/etc/openbook/frontend/seelog.xml
	cp -r deb/frontend/* $(BUILD_DIR)/frontend/

package_frontend: build_frontend
	sed -i s/Version:.*/Version:\ $(VERSION)/g $(BUILD_DIR)/frontend/DEBIAN/control
	fakeroot dpkg-deb --build build/frontend frontend_$(VERSION)_amd64.deb

######## сборка migrations ########

build_migrations: dependencies
	rm -rf $(BUILD_DIR)/migrations
	mkdir -p $(BUILD_DIR)/migrations/usr/lib/newsgun/
	cp -r ./migrations $(BUILD_DIR)/migrations/usr/lib/newsgun/
	cp -r deb/migrations/* $(BUILD_DIR)/migrations

package_migrations: build_migrations
	sed -i s/Version:.*/Version:\ $(VERSION)/g $(BUILD_DIR)/migrations/DEBIAN/control
	fakeroot dpkg-deb --build $(BUILD_DIR)/migrations migrations_$(VERSION)_amd64.deb

publish_migrations: package_migrations
	$(eval PACKAGE = $(shell ls *.deb | grep migrations))
	scp $(PACKAGE) aptly@apt.2rll.net:upload/
	ssh aptly@apt.2rll.net aptly repo add 2rll upload/$(PACKAGE)
	ssh apt.2rll.net -l aptly aptly publish update -gpg-key=2rll wheezy 2rll
	rm $(PACKAGE)

build: build_frontend build_migrations

package: package_frontend package_migrations

######## тестирование проекта ########

test: 
	$(eval CONFIG ?= $(shell readlink -e resources/configuration/config.test.gcfg))
	$(eval C := $(shell readlink -e $(CONFIG)))

	$(GO) test $(PROJECT)/common/model/book -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/book/category -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/book/category/utils -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/book/price -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/author -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/publisher -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/user -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/user/subscription -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/user/address -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/subscription -test.config="$(C)"
	$(GO) test $(PROJECT)/common/model/order -test.config="$(C)"
	$(GO) test $(PROJECT)/frontend -test.config="$(C)"
	$(GO) test $(PROJECT)/frontend/form -test.config="$(C)"
	$(GO) test $(PROJECT)/backend -test.config="$(C)"

######## быстрая накатка миграций ########

migrate_up:
	-migrate -url postgres://developer:developer@localhost/openbook -path ./migrations/postgresql up

migrate_down:
	-migrate -url postgres://developer:developer@localhost/openbook -path ./migrations/postgresql down

migrate: migrate_down migrate_up

######## очистка ########

clear:
	rm -rf build
	rm -f *.deb
