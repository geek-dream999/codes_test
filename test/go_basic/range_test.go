package go_basic

/*
下面函数通过遍历切片，打印切片的下标和元素值，请问性能上有没有可优化的空间？
答：遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值。
*/
func RangeSlice(slice []int) {
	for index, value := range slice {
		_, _ = index, value
	}
}
