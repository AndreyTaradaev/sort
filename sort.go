package sort 

func bubbleSort(ar []int) {
//	var  IsChange = false
	for i:=len(ar)-1; i>-1; i--{
		//fmt.Println(i)
		IsChange  := false
		for j:=0; j <i ;j++{
			if ar[j] >ar[j+1] {
			ar[j],ar[j+1] = ar[j+1],ar[j]	
			IsChange = true
			}		
		}
		 if !IsChange  {
			 break
		}
	}


	
}


func SelectionSort(ar []int){
	var Minind ,Minval = 0,ar[0]	
	for i:=0; i<len(ar);i++{
		Minind ,Minval = i,ar[i]	
			for j := i; j < len(ar); j++ {				
				if(ar[j]<Minval) {
					Minind = j
					Minval = ar[j]
					
 				}		
			}			
		ar[i], ar[Minind]	 = 	Minval ,ar[i]		
	}


	
}

func InsertSort(ar []int){
	//fmt.Println( ar)
	for i:=1; i<len(ar);i++  {
	j:=i-1	  // номер предыдущего элемента 
	curvalue := ar[i]  //	  значение текущего элемента
	for  ;j>=0 && ar[j]> curvalue ; j-- {   //  проверка пред. элемент больше текущего элемента
		//j--
		ar[j+1] = ar[j]  //если да тогда пред. элемент вставляем в текущий        
	}
	ar[j+1] = curvalue	
	}	
}


func SortMerge(ar [] int) []int {
	if (len(ar) ==1) {		
		return ar
	} 
	var middle = len(ar)/2
	left := ar[:middle] 
	right := ar[middle:]	
	left = SortMerge(left)
	right =SortMerge(right)
	if ( left[len(left)-1]	<=right[0]){ // считаем что  лев и правый массивы отсортированы  поэтому можно безболезненно
	// сложить массивы		
	left = append(left,right...)			
	return left
	}
	s:= merge(left,right)		
return s
}

func merge (left  , right []int) []int{
	var L,R = 0,0
	var s  []int
	for i:=0;L <len(left) && R<len(right)  ; i++ {
		if(left[L] <=right[R] )	{
		s= append(s, left[L])
		L++		
		} else		{
		s=  append(s,  right[R])
		R++		
		}
	}
	if L<len(left){
			s= append(s, left[L:]...)
			}
	if(R<len(right)){
			s = append(s, right[R:]...)
	}
return s
}

func SortQuick(ar[]int) []int {
	if (len(ar)	<=1){
		return  ar
	}	
	var  less,more []int
	pivot := ar[ partion(ar)]
	for i:=1; i<len(ar);i++{
		if ar[i] >pivot{
			more = append(more, ar[i])				
		}else{
			less = append(less, ar[i])
		}		
	}
	less = SortQuick(less)
	more = SortQuick(more)
	less = append(less, more...)
	return less
}

func partion(ar []int) int{
	length := len(ar) 
		mid := length/2
	switch  {
	case ar[mid] < ar[0]:
		ar[0],ar[mid] = ar[mid],ar[0]
	case ar[length-1] < ar[0]:
		ar[length-1],ar[0] = ar[0],ar[length-1]	
	case ar[length-1] < ar[mid]:
		ar[length-1],ar[mid] = ar[mid],ar[length-1]			
	}
	return mid
	}



