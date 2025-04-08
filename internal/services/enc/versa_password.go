package enc

var VersaPasswordEncoder = NewPasswordEncoder()

func VersaPassword(p string) (string, error) {
	return VersaPasswordEncoder.Encode(p)
}
