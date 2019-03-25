package utils

func GetMessage(crocsUse CrocsUse) string {
	switch crocsUse {
	case DontUse:
		return "Não."
	case UseWithSocks:
		return "Sim, com meias."
	default:
		return "Sim!"
	}
}
