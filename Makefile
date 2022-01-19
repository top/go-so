.PHONY: all so bin clean

all: so bin

so:
	for i in $$(find . -name 'plugin_*.go'); do go build -buildmode=plugin $$i; done;

bin:
	go build main.go

clean:
	rm -rf *.so main
