// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package external

// UseARNRegionProvider is an interface for retrieving external configuration value for UseARNRegion
type UseARNRegionProvider interface {
	GetS3UseARNRegion() (value bool, ok bool, err error)
}

// ResolveUseARNRegion extracts the first instance of a UseARNRegion from the config slice.
// Additionally returns a boolean to indicate if the value was found in provided configs, and error if one is encountered.
func ResolveUseARNRegion(configs []interface{}) (value bool, ok bool, err error) {
	for _, cfg := range configs {
		if p, pOk := cfg.(UseARNRegionProvider); pOk {
			value, ok, err = p.GetS3UseARNRegion()
			if err != nil {
				return value, false, err
			}
			if ok {
				break
			}
		}
	}

	return value, ok, err
}
