NAME = pinger
VERSION = 0.1

build_image:
	docker run -t $(NAME):$(VERSION) .

tag_latest:
	docker tag -f $(NAME):$(VERSION) $(NAME):latest

