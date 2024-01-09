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
          spirale \
          text-transform \
          vertigo \
          fadengraphik

PNG_FILES = $(TARGETS:=.png)

all: $(TARGETS)

png: $(PNG_FILES)

clean:
	rm $(TARGETS)

.SUFFIXES: .go .png

%: %.go
	go build -o $@ $<

%.png: %
	./$<

