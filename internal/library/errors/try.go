package errors

func Try[e any](data e, err error) e {
	if err != nil {
		panic(err)
	}
	return data
}
