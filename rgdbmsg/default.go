package rgdbmsg

func _default[T int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | float32 | float64 | string](val **T, def T) {
	if *val == nil {
		*val = &def
	}
}
