package gue

func Gue0[ RT any](f func() (RT, error), e *error) func() RT {
	return func() RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f()
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue00(f func() error, e *error) func() {
	return func() {
		if *e != nil {
			return
		}
		err := f()
		if err != nil {
			*e = err
		}
	}
}
	
func Gue1[R, RT any](f func(R) (RT, error), e *error) func(R) RT {
	return func(r R) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue01[R any](f func(R) error, e *error) func(R) {
	return func(r R) {
		if *e != nil {
			return
		}
		err := f(r)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue2[R, S, RT any](f func(R, S) (RT, error), e *error) func(R, S) RT {
	return func(r R, s S) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue02[R, S any](f func(R, S) error, e *error) func(R, S) {
	return func(r R, s S) {
		if *e != nil {
			return
		}
		err := f(r, s)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue3[R, S, T, RT any](f func(R, S, T) (RT, error), e *error) func(R, S, T) RT {
	return func(r R, s S, t T) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue03[R, S, T any](f func(R, S, T) error, e *error) func(R, S, T) {
	return func(r R, s S, t T) {
		if *e != nil {
			return
		}
		err := f(r, s, t)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue4[R, S, T, U, RT any](f func(R, S, T, U) (RT, error), e *error) func(R, S, T, U) RT {
	return func(r R, s S, t T, u U) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue04[R, S, T, U any](f func(R, S, T, U) error, e *error) func(R, S, T, U) {
	return func(r R, s S, t T, u U) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue5[R, S, T, U, V, RT any](f func(R, S, T, U, V) (RT, error), e *error) func(R, S, T, U, V) RT {
	return func(r R, s S, t T, u U, v V) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u, v)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue05[R, S, T, U, V any](f func(R, S, T, U, V) error, e *error) func(R, S, T, U, V) {
	return func(r R, s S, t T, u U, v V) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u, v)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue6[R, S, T, U, V, W, RT any](f func(R, S, T, U, V, W) (RT, error), e *error) func(R, S, T, U, V, W) RT {
	return func(r R, s S, t T, u U, v V, w W) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u, v, w)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue06[R, S, T, U, V, W any](f func(R, S, T, U, V, W) error, e *error) func(R, S, T, U, V, W) {
	return func(r R, s S, t T, u U, v V, w W) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u, v, w)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue7[R, S, T, U, V, W, X, RT any](f func(R, S, T, U, V, W, X) (RT, error), e *error) func(R, S, T, U, V, W, X) RT {
	return func(r R, s S, t T, u U, v V, w W, x X) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u, v, w, x)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue07[R, S, T, U, V, W, X any](f func(R, S, T, U, V, W, X) error, e *error) func(R, S, T, U, V, W, X) {
	return func(r R, s S, t T, u U, v V, w W, x X) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u, v, w, x)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue8[R, S, T, U, V, W, X, Y, RT any](f func(R, S, T, U, V, W, X, Y) (RT, error), e *error) func(R, S, T, U, V, W, X, Y) RT {
	return func(r R, s S, t T, u U, v V, w W, x X, y Y) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u, v, w, x, y)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue08[R, S, T, U, V, W, X, Y any](f func(R, S, T, U, V, W, X, Y) error, e *error) func(R, S, T, U, V, W, X, Y) {
	return func(r R, s S, t T, u U, v V, w W, x X, y Y) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u, v, w, x, y)
		if err != nil {
			*e = err
		}
	}
}
	
func Gue9[R, S, T, U, V, W, X, Y, Z, RT any](f func(R, S, T, U, V, W, X, Y, Z) (RT, error), e *error) func(R, S, T, U, V, W, X, Y, Z) RT {
	return func(r R, s S, t T, u U, v V, w W, x X, y Y, z Z) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f(r, s, t, u, v, w, x, y, z)
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue09[R, S, T, U, V, W, X, Y, Z any](f func(R, S, T, U, V, W, X, Y, Z) error, e *error) func(R, S, T, U, V, W, X, Y, Z) {
	return func(r R, s S, t T, u U, v V, w W, x X, y Y, z Z) {
		if *e != nil {
			return
		}
		err := f(r, s, t, u, v, w, x, y, z)
		if err != nil {
			*e = err
		}
	}
}
	
