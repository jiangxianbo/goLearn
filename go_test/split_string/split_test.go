package split_string

import (
	"reflect"
	"testing"
)

//func TestSplit1(t *testing.T) {
//	got := Split("babcbef", "b")
//	want := []string{"", "a", "c", "ef"}
//	if !reflect.DeepEqual(got, want) {
//		// 测试用例失败
//		t.Errorf("want:%v, but got:%v", want, got)
//	}
//}
//
//func TestSplit2(t *testing.T) {
//	got := Split("a:b:c", ":")
//	want := []string{"a", "b", "c"}
//	if !reflect.DeepEqual(got, want) {
//		// 测试用例失败
//		t.Errorf("want:%v, but got:%v", want, got)
//	}
//}
//
//func TestSplit3(t *testing.T) {
//	got := Split("abcef", "bc")
//	want := []string{"a", "ef"}
//	if !reflect.DeepEqual(got, want) {
//		t.Fatalf("want:%v, but got:%v", want, got)
//	}
//}

// 测试组
//func TestSplit(t *testing.T) {
//	type testCase struct {
//		str  string
//		sep  string
//		want []string
//	}
//	testGroup := []testCase{
//		{"babcbef","b", []string{"", "a", "c", "ef"}},
//		{"a:b:c",":", []string{ "a", "b", "c"}},
//		{"abcef","bc", []string {"a", "ef"}},
//		{"沙河有沙又有河","有", []string {"沙河", "沙又", "河"}},
//	}
//	for _, tc := range testGroup {
//		got := Split(tc.str, tc.sep)
//		if !reflect.DeepEqual(got, tc.want) {
//			t.Fatalf("want:%v, but got:%v", tc.want, got)
//		}
//	}
//}

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": {"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case2": {"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": {"abcef", "bc", []string{"a", "ef"}},
		"case4": {"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v, but got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
