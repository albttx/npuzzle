
BINARY	=	npuzzle 
SOURCES	=	$(shell find . -type f -name "*.go")
PKG		=	./npuzzle

all: $(BINARY)
	
$(BINARY)	: $(SOURCES)
	@go get ./...
	@go build ./cmd/$@
clean:
	@rm -f $(BINARY)
re: clean all
run: all
	@./$(BINARY) -f map/3_solvable.txt
random: all
	@./$(BINARY) -size 3
