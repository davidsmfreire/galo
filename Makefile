.PHONY: galo minesweeper

galo:
	gcc galo/main.c -o ./galo/build/galo && ./galo/build/galo

minesweeper:
	cd minesweeper && go run main.go
