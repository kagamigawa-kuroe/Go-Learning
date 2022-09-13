package main


type Translate interface {
    Translate(x float64, y float64)
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

/// limit de type
type number interface{
	int64 | float32 | uint64
}

func main(){


}