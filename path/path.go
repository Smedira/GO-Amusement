package Path

type path struct {
	length int
	i1 intersection
	i2 intersection

}

type intersection struct {
	paths []path
}

type space struct {
	area int
	i []intersection
}