package convert

func IsTimeValid(time string) bool {
	if time == "<nil>" {
		return false
	}
	return true
}
