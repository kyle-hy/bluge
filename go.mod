module github.com/blugelabs/bluge

go 1.16

require (
	github.com/RoaringBitmap/roaring v1.5.0
	github.com/axiomhq/hyperloglog v0.0.0-20230201085229-3ddf4bad03dc
	github.com/bits-and-blooms/bitset v1.8.0
	github.com/blevesearch/go-porterstemmer v1.0.3
	github.com/blevesearch/mmap-go v1.0.4
	github.com/blevesearch/segment v0.9.1
	github.com/blevesearch/snowballstem v0.9.0
	github.com/blevesearch/vellum v1.0.10
	github.com/blugelabs/bluge_segment_api v0.2.0
	github.com/blugelabs/ice v1.0.0
	github.com/caio/go-tdigest v3.1.0+incompatible
	github.com/dgryski/go-metro v0.0.0-20211217172704-adc40b04c140 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/leesper/go_rng v0.0.0-20190531154944-a612b043e353 // indirect
	github.com/spf13/cobra v1.7.0
	golang.org/x/sys v0.12.0
	golang.org/x/text v0.13.0
	gonum.org/v1/gonum v0.7.0 // indirect
)

replace github.com/blugelabs/ice => github.com/zinclabs/ice v1.1.3

replace github.com/blugelabs/bluge_segment_api => github.com/zinclabs/bluge_segment_api v1.0.0
