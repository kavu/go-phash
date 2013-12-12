package phash

import "testing"

var hashTestData = []struct {
	file     string
	expected uint64
}{
	{"test_data/jpg/cat.jpg", 11220389026139797626},
	{"test_data/png/gopher.png", 4855808264951085874},
	{"test_data/gif/zebras.gif", 7086864532766789979},
	{"test_data/jpg/zebras.jpg", 7086864532766789979},
	{"test_data/png/zebras.png", 7086864532766789979},
	{"test_data/png/grass.png", 0},
	{"test_data/gif/animated_zebras.gif", 0},
}

func BenchmarkImageHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ImageHash(hashTestData[0].file)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestImageHash(t *testing.T) {
	for _, test := range hashTestData {
		hash, err := ImageHash(test.file)
		if err != nil {
			t.Fatal(err)
		}

		if hash != test.expected {
			t.Errorf("%s: Hash Mismatch: expected %d, got %d", test.file, test.expected, hash)
		}
	}
}
