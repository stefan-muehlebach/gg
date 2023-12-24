TARGETS = bauhaus \
          chessboard \
          color-cells \
          divided-circle \
          farn \
          flower-of-life \
          graph \
          irrgarten \
          moiree \
          raster \
          raster-image \
          relief \
          scatter \
          spirale \
          text-transform \
          vertigo

PNG_FILES = $(TARGETS:=.png)

all: $(TARGETS)

png: $(PNG_FILES)

clean:
	rm $(TARGETS)
	rm $(PNG_FILES)

.SUFFIXES: .go .png

%: %.go
	go build -o $@ $<

%.png: %
	./$<

