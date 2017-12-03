.PHONY: extract
extract: iid
	docker run --rm $(shell cat iid) | tar xf -
	rm iid

iid:
	docker build --no-cache --iidfile iid .
