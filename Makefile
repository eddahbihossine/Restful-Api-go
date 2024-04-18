SRC = main.go
BIN = executable

# Path: Makefile

all : $(SRC)
	go build -o $(BIN) $(SRC)

clean:
	rm -f $(BIN)
fclean: clean

