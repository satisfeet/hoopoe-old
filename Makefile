SHELL := /bin/bash

start:
	@gin -p 3000 -a 4000 run \
		--host :4000 \
		--auth bodokaiser:secret \
		--mongo localhost/satisfeet
