package days

type Day interface {
	Part1([]string) (string, error)
	Part2([]string) (string, error)
}
