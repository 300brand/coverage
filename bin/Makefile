BINS := $(dir $(wildcard */main.go))

build:
	go install -v $(addprefix git.300brand.com/coverage/bin/,$(BINS))

start: build
	sky deploy git.300brand.com/coverage/bin/service-StorageReader
	sky deploy git.300brand.com/coverage/bin/service-StorageWriter
	sky deploy git.300brand.com/coverage/bin/service-Feed
	sky deploy git.300brand.com/coverage/bin/service-FeedDownload
	sky deploy git.300brand.com/coverage/bin/service-FeedProcess
	sky deploy git.300brand.com/coverage/bin/service-Queue
	# Two of these
	sky deploy git.300brand.com/coverage/bin/service-ArticleBody
	sky deploy git.300brand.com/coverage/bin/service-ArticleBody
	# Two of these
	sky deploy git.300brand.com/coverage/bin/service-ArticleDownload
	sky deploy git.300brand.com/coverage/bin/service-ArticleDownload
	sky deploy git.300brand.com/coverage/bin/service-Manager
	sky deploy git.300brand.com/coverage/bin/server-jsonrpc

stop:
	sky stop

restart: stop start

.PHONY: build restart start stop
