package model

type Student struct{
	Name string
	Score float64
}

type student struct{
	Name string
	score float64
}

func  Newstudent(n1 string, s1 float64) *student  {
	return &student{
		Name : n1,
		score : s1,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}