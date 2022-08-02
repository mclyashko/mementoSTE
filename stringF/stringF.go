package stringF

func DeleteCompositePrefix(in string, prefix string) string {
	if in[0:len(prefix)] == prefix {
		return in[len(prefix):]
	}
	return in
}
