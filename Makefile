

all:
	go build

test: savvy
	cd .. && savvy/savvy && cd -
