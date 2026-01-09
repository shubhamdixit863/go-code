package funPackage

func Cal(a int,b int,c int ,d int ,e int)(int,int,int){
	Sum:=a+b+c+d+e
	Sub:=a-b-c-d-e
	Mul:=a*b*c*d*e
	return Sum,Sub,Mul
}

func DivRem(g int,h int)(int,int){
	div:=g/h
	rem:=g%h
	return div,rem
}