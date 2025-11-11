install_pw:
	go run github.com/playwright-community/playwright-go/cmd/playwright@v0.5200.1 install --with-deps

run_local: install_pw
	LOCAL_PATH=$(shell pwd) go run .