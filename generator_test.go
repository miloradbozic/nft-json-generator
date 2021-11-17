package main

import (
	"testing"
)

func testCreators() []string {
	return []string{"5QYKVtTDPqvGVpgGid5hvTzErvzDuzSJCxban71xSZvA", "A2SnWW5tL9hExpt5FuFXgb45XcBVTYSBhawZFGdaGf4g"}
}

func testList() string {
	var list string
	for i := 0; i < 200; i++ {
		for j := 0; j < 100; j++ {
			list = list + "Some trait"
			if j < 99 {
				list = list + ","
			}
		}

		if i < 1000 {
			list = list + "\r\n"
		}
	}

	return list
}

func benchmarkGenerator(list string, creators []string, paralelize bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateNftJsons(list, creators, paralelize)
	}
}

func BenchmarkGenerator(b *testing.B) {
	benchmarkGenerator(testList(), testCreators(), false, b)
}

func BenchmarkGeneratorParallel(b *testing.B) {
	benchmarkGenerator(testList(), testCreators(), true, b)
}
