BINS := $(dir $(wildcard */main.go))

build:
	go install -v $(addprefix github.com/300brand/coverage/bin/,$(BINS))

start: build
	sky deploy github.com/300brand/coverage/bin/service-StorageReader
	sky deploy github.com/300brand/coverage/bin/service-StorageWriter
	sky deploy github.com/300brand/coverage/bin/service-Feed
	sky deploy github.com/300brand/coverage/bin/service-FeedDownload
	sky deploy github.com/300brand/coverage/bin/service-FeedProcess
	sky deploy github.com/300brand/coverage/bin/service-Queue
	# Two of these
	sky deploy github.com/300brand/coverage/bin/service-ArticleBody
	sky deploy github.com/300brand/coverage/bin/service-ArticleBody
	# Two of these
	sky deploy github.com/300brand/coverage/bin/service-ArticleDownload
	sky deploy github.com/300brand/coverage/bin/service-ArticleDownload
	sky deploy github.com/300brand/coverage/bin/service-Manager
	sky deploy github.com/300brand/coverage/bin/server-jsonrpc

stop:
	sky stop

restart: stop start

.PHONY: build restart start stop
