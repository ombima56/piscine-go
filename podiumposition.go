package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)/2; i++ {
		q := len(podium) - 1 - i
		podium[i], podium[q] = podium[q], podium[i]
	}
	return podium
}
