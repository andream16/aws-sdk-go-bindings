package format

// StrToPtr is an helper function used to return a pointer to a passed string.
func StrToPtr(s string) *string {
	return &s
}

// Int64ToPtr is an helper function used to return a pointer to a passed int64.
func Int64ToPtr(i int64) *int64 {
	return &i
}
