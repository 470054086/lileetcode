package main

type Abs interface {
	add()
}

type Flag struct {

}

func (f *Flag) add(){
}



func main() {
	 var a Abs
	 f:= Flag{}
	 a = &f
	 a.add()

}
