package utils

func Apply[Tin any, Tout any](
	arr []Tin, f func(Tin) Tout,
) []Tout {
	out := make([]Tout, len(arr))

	for i, in := range arr {
		out[i] = f(in)
	}

	return out
}
