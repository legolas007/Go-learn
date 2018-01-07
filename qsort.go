package main

func qsort(a []int, l int, r int)  {
	if l >= r{
		return
	}
	i ,j:= l,r
	tmp := a[i]
	for i < j{
		for i < j && a[j] >=tmp{
			j--
		}
		a[i] = a[j]
		for i < j && a[i] <=tmp{
			i++
		}
		a[j] = a[i]
	}
	a[i] = tmp
	qsort(a,l,i-1)
	qsort(a,j+1,r)
}

