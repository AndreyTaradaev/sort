package sort

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func CreateSlice(max, size  int) []int {
	
	ar := make([]int, size,size)
	for i := range ar {
		ar[i] = rand.Intn(max*10) - max*5 // ограничиваем случайное значение
	}
return ar
}

func CreateSortedSlice( size  int) []int {
	elem := -1* size/2

	ar := make([]int, size,size)
	for i := range ar {
		ar[i] = elem
		elem ++
	}
return ar
}

func CreateWorstCaseSlice(size  int) []int {
	elem := size/2
	ar := make([]int, size,size)
	for i := range ar {
		ar[i] = elem 
		elem--
	}
return ar
}

func CreateEqualSlice(size  int) []int {
	elem := rand.Intn(size*10) - size*5
	ar := make([]int, size,size)
	for i := range ar {
		ar[i] = elem 		
	}
return ar
}



func CheckSort(ar [] int) bool {

	for i:=1; i< len(ar) ;i++ {
		if  ar[i-1] > ar[i] {
			return false
		}
	}
	return true
	}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	b.Run("very big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100000, 200000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

}

func BenchmarkSelectionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10, 10)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100, 1000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10000, 100000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	b.Run("very big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100000,200000)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})	
}

func BenchmarkInsertSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10, 10)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100, 1000)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10000, 100000)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	
	b.Run("very big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100000, 200000)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
}


func BenchmarkSortMerge(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10, 10)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100, 1000)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10000, 100000)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100000, 200000)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
}


func BenchmarkSortQuick(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10, 10)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100, 1000)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(10000, 100000)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSlice(100000, 200000)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
}


func BenchmarkSortedArray(b *testing.B) {
	size := 100000
	b.Run("bubbleSort", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSortedSlice(size)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SelectionSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSortedSlice(size)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("InsertSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSortedSlice(size)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	b.Run("SortMerge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSortedSlice(size)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SortQuick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateSortedSlice(size)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

}

func BenchmarkWorkCaseArray(b *testing.B) {
	size := 100000
	b.Run("bubbleSort", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateWorstCaseSlice(size)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SelectionSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateWorstCaseSlice(size)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("InsertSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateWorstCaseSlice(size)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	b.Run("SortMerge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateWorstCaseSlice(size)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SortQuick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateWorstCaseSlice(size)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

}

func BenchmarkEqualArray(b *testing.B) {
	size := 100000
	b.Run("bubbleSort", func(b *testing.B) {
		b.StopTimer()
		//b.Logf("Count attempt %d",b.N)
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateEqualSlice(size)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SelectionSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateEqualSlice(size)
			b.StartTimer()
			SelectionSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("InsertSort", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateEqualSlice(size)
			b.StartTimer()
			InsertSort(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})
	b.Run("SortMerge", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateEqualSlice(size)
			b.StartTimer()
			ar = SortMerge(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

	b.Run("SortQuick", func(b *testing.B) {
		b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ar := CreateEqualSlice(size)
			b.StartTimer()
			ar = SortQuick(ar)
			b.StopTimer()
			if(!CheckSort(ar)){
				b.Fatalf("Array not sorted attempt %d",i)
			}
		}
	})

}