TARGETS = bauhaus \
          chessboard \
          color-cells \
          divided-circle \
          fadengraphik \
          farn \
          flower-of-life \
          graph \
          irrgarten \
          moiree \
          raster \
          raster-image \
          relief \
          spirale \
          tictactoe01 \
          tictactoe02 \
          tictactoe03 \
          text-transform \
          vertigo

PNG_FILES = $(TARGETS:=.png)

all: $(TARGETS)

png: $(PNG_FILES)

clean:
	rm -f $(TARGETS)
	rm -f $(PNG_FILES)

.SUFFIXES: .go .png

%: %.go
	go build -o $@ $<

%.png: %
	./$<

