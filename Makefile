ioffice:
	go mod tidy
	go build .

.PHONY: clean
	rm ioffice

.PHONY: install
install: ioffice
	install ioffice /usr/local/bin/ioffice

.PHONY: uninstall
uninstall:
	rm /usr/local/bin/ioffice
