package complexnumbers

import "math"

type Number struct {
    RealPart float64
    ImagPart float64
}

func (n Number) Real() float64 {
	return n.RealPart
}

func (n Number) Imaginary() float64 {
	return n.ImagPart
}

func (n1 Number) Add(n2 Number) Number {
	return Number {RealPart: n1.RealPart + n2.RealPart, ImagPart: n1.ImagPart + n2.ImagPart}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number {RealPart: n1.RealPart - n2.RealPart, ImagPart: n1.ImagPart - n2.ImagPart}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number {RealPart: (n1.RealPart * n2.RealPart - n1.ImagPart * n2.ImagPart), ImagPart: (n1.ImagPart * n2.RealPart) + (n2.ImagPart * n1.RealPart)}
}

func (n Number) Times(factor float64) Number {
	return Number {RealPart: n.RealPart * factor, ImagPart: n.ImagPart * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	return Number {RealPart: (n1.RealPart * n2.RealPart + n1.ImagPart * n2.ImagPart) / (n2.RealPart * n2.RealPart + n2.ImagPart * n2.ImagPart), ImagPart: (n1.ImagPart * n2.RealPart - n2.ImagPart * n1.RealPart) / (n2.RealPart * n2.RealPart + n2.ImagPart * n2.ImagPart)}
}

func (n Number) Conjugate() Number {
	return Number{RealPart: n.RealPart, ImagPart: -1 * n.ImagPart}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.RealPart * n.RealPart + n.ImagPart * n.ImagPart)
}

func (n Number) Exp() Number {
	return Number {RealPart: math.Exp(n.RealPart) * math.Cos(n.ImagPart), ImagPart: math.Exp(n.RealPart) * math.Sin(n.ImagPart)}
}
