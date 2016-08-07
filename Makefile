ELM_FILES = $(shell find elm/ -type f -name '*.elm' -or -name '*.js')
LESS_FILES = $(shell find assets/css/ -type f -name '*.less')
PUBLIC_FILES = $(shell find public/ -type f)
TEMPLATE_FILES = $(shell find templates/ -type f)

all: public/elm.min.js public/main.css bindata.go

.PHONY: clean

clean:
	rm -f public/elm.js public/elm.min.js public/main.css bindata.go

public/elm.js: $(ELM_FILES)
	cd elm && elm make --warn --output ../public/elm.js --yes src/BuildPage.elm src/JobPage.elm src/PipelinePage.elm

public/main.css: $(LESS_FILES)
	lessc --clean-css="--advanced" assets/css/main.less $@

public/elm.min.js: public/elm.js
	uglifyjs < $< > $@

bindata.go: $(PUBLIC_FILES) $(TEMPLATE_FILES)
	go-bindata ${DEV} -pkg web templates/... public/...
	go fmt bindata.go
