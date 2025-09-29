package triangle

type Kind int

const (
    NaT Kind = iota
    Equ
    Iso
    Sca
)

func KindFromSides(a, b, c float64) Kind {
    if !valid(a,b,c) {
        return NaT
	}
    if a == b && b == c && a == c {
        return Equ
    }
    if a == b || b == c || a == c {
        return Iso
    }
	return Sca
}
func valid(a, b, c float64) bool {
  ab := a + b > c
  ac := a + c > b
  bc := b + c > a
  passesInequality := ab && ac && bc
  validLength := a > 0 && b > 0 && c > 0

  return passesInequality && validLength
}