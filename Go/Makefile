PROJECTNAME := $(shell basename "$(PWD)")
STIME 		:= $(shell date +%s)

build: 
	@echo ">  Building Program..."
	GOPRIVATE=github go build -ldflags="-s -w" -o bin/${PROJECTNAME} main.go; 
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

## rest: start without docker
.PHONY: rest
rest: build
	@echo "  >  Starting Program..."
	./bin/${PROJECTNAME} rest
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"