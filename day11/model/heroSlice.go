package model

//sort.Sort(data Interface)
type Hero struct {
	Name string
	Age int
}

//声名一个Hero结构体切片类型
type HeroSlice []Hero 

//实现sort.Sort(data Interface)接口方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}