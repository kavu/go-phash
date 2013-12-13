package phash

import (
	"fmt"
	"testing"
)

var hashTestData = []struct {
	file     string
	expected uint64
}{
	{"test_data/jpg/cat.jpg", 11220389026139797626},
	{"test_data/png/gopher.png", 4855808264951085874},
	{"test_data/gif/zebras.gif", 7086864532766789979},
	{"test_data/jpg/zebras.jpg", 7086864532766789979},
	{"test_data/png/zebras.png", 7086864532766789979},
	// {"test_data/png/grass.png", 54086765383280},
	// {"test_data/gif/animated_zebras.gif", 7086864532766789979},
}

// ImageHash

func BenchmarkImageHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ImageHashDCT(hashTestData[0].file)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestImageHash(t *testing.T) {
	for _, test := range hashTestData {
		hash, err := ImageHashDCT(test.file)
		if err != nil {
			t.Fatal(err)
		}

		if hash != test.expected {
			t.Errorf("%s: Hash Mismatch: expected %d, got %d", test.file, test.expected, hash)
		}
	}
}

func ExampleImageHashDCT() {
	hash, err := ImageHashDCT("test_data/jpg/cat.jpg")
	if err != nil {
		panic(err)
	}
	fmt.Println(hash)
	// Output: 11220389026139797626
}

// HammingDistanceForHashes

func BenchmarkHammingDistanceForHashesForSameImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := HammingDistanceForHashes(7086864532766789979, 7086864532766789979)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHammingDistanceForHashesForDifferentImages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := HammingDistanceForHashes(7086864532766789979, 11220389026139797626)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestHammingDistanceForHashesForSameImage(t *testing.T) {
	d, err := HammingDistanceForHashes(7086864532766789979, 7086864532766789979)
	if err != nil {
		t.Fatal(err)
	}

	if d != 0 {
		t.Errorf("Distance Mismatch: expected 0, got %d", d)
	}
}

func TestHammingDistanceForHashesForDifferentImages(t *testing.T) {
	d, err := HammingDistanceForHashes(11220389026139797626, 4855808264951085874)
	if err != nil {
		t.Fatal(err)
	}

	if d == 0 {
		t.Error("Distance Mismatch: got unexpected 0")
	}
}

func ExampleHammingDistanceForHashes() {
	d, err := HammingDistanceForHashes(11220389026139797626, 4855808264951085874)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	// Output: 30
}

// HammingDistanceForFiles

func BenchmarkHammingDistanceForFilesForSameImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := HammingDistanceForFiles("test_data/jpg/cat.jpg", "test_data/jpg/cat.jpg")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkHammingDistanceForFilesForDifferentImages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := HammingDistanceForFiles("test_data/jpg/cat.jpg", "test_data/png/gopher.png")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestHammingDistanceForFilesForSameImage(t *testing.T) {
	d, err := HammingDistanceForFiles("test_data/jpg/cat.jpg", "test_data/jpg/cat.jpg")
	if err != nil {
		t.Fatal(err)
	}

	if d != 0 {
		t.Errorf("Distance Mismatch: expected 0, got %d", d)
	}
}

func TestHammingDistanceForFilesForDifferentImages(t *testing.T) {
	d, err := HammingDistanceForFiles("test_data/jpg/cat.jpg", "test_data/png/gopher.png")
	if err != nil {
		t.Fatal(err)
	}

	if d == 0 {
		t.Error("Distance Mismatch: got unexpected 0")
	}
}

func ExampleHammingDistanceForFiles() {
	d, err := HammingDistanceForFiles("test_data/jpg/cat.jpg", "test_data/png/gopher.png")
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	// Output: 30
}
