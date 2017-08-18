// Package erratum contains errata related to error handling
package erratum

const testVersion = 2

// Use Frobs a resource given a ResourceOpener and an input to Frob with
func Use(opener ResourceOpener, input string) (err error) {
	var res Resource
	// open the resource, retrying in case of TransientError
	for {
		res, err = opener()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			} else {
				return err
			}
		}
		break
	}

	// ensure we close the resource
	defer res.Close()

	// handle panic with FrobError
	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				res.Defrob(frobErr.defrobTag)
			}
			err = r.(error)
		}
	}()

	// Frob the resource
	res.Frob(input)

	// err is either nil or the panic error
	return err
}
