# Makefile
#
usage:
	@echo ""
	@echo "Task                 : Description"
	@echo "-----------------    : -------------------"
	@echo "make setup           : Install all necessary dependencies"
	@echo "make build           : Generate production build for current OS"
	@echo "make test            : Execute unit test suite"
	@echo "make re-test         : Execute unit test suite in continuos mode ( retesting on changes )"
	@echo "make restore-vendors : Fetch dependencies to the vendor/ directory"
	@echo ""

build:
	gb build

test:
	gb test

re-test:
	reflex -r src/ $(MAKE) test

dev-deps:
	go get github.com/cespare/reflex
	go get github.com/constabulary/gb/...
	go get github.com/atotto/clipboard

restore-vendors:
	gb vendor restore

ninja:
	./bin/socialcp register eu@gmail.com
	./bin/socialcp add 1 ninja@gmail.com

setup: dev-deps restore-vendors

install:
	@ln -s -f `pwd`/bin/socialcp /usr/local/bin/socialcp
	cp -R ./automator/socialcp-recv1.workflow/ ~/Library/Services/
	cp -R ./automator/socialcp-send1.workflow/ ~/Library/Services/
