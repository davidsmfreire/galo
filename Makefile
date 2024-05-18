.PHONY: connect_four galo minesweeper scramble snake

build-dir-%:
	mkdir $*/build 2>/dev/null || true && \
	echo -e "*" > $*/build/.gitignore || true

galo: build-dir-galo
	gcc galo/main.c -o ./galo/build/galo && ./galo/build/galo

minesweeper:
	cd minesweeper && go run main.go

scramble:
	php scramble/main.php

snake: build-dir-snake
	rustc snake/main.rs -o snake/build/main && \
	./snake/build/main
