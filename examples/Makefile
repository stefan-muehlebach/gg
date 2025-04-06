TARGETS = beziers \
          circles \
          clip \
          colormap \
          colors \
          colorWheel \
          concat \
          crisp \
          cubic \
          ellipses \
          fontmap \
          gofont \
		  gradient-conic \
		  gradient-linear \
		  gradient-mandelbrot \
		  gradient-radial \
		  gradient-text \
		  gradient-transformed \
          lines \
          linewidth \
          lorem \
          mask \
          meme \
          pattern-fill \
          quadratic \
		  rotated-image \
		  rotated-text \
		  sine \
		  spiral \
		  star \
		  stars \
		  text \
		  tiling \
		  transform \
		  unicode \
		  wrap

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

