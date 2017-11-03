all:
	$(error please pick a target)

install-reveal:
	npm install reveal-md

run-slides:
	./node_modules/.bin/reveal-md slides.md --theme serif
