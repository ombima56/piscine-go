package piscine

func PodiumPosition(podium [][]string) [][]string {
	corrected := make([][]string, len(podium))
	for i := 0; i < len(podium); i++ {
		corrected[i] = podium[len(podium)-1-i]
	}
	return corrected
}
