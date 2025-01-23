// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// BuildVolumeApplyConfiguration represents a declarative configuration of the BuildVolume type for use
// with apply.
type BuildVolumeApplyConfiguration struct {
	Name   *string                              `json:"name,omitempty"`
	Source *BuildVolumeSourceApplyConfiguration `json:"source,omitempty"`
	Mounts []BuildVolumeMountApplyConfiguration `json:"mounts,omitempty"`
}

// BuildVolumeApplyConfiguration constructs a declarative configuration of the BuildVolume type for use with
// apply.
func BuildVolume() *BuildVolumeApplyConfiguration {
	return &BuildVolumeApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BuildVolumeApplyConfiguration) WithName(value string) *BuildVolumeApplyConfiguration {
	b.Name = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *BuildVolumeApplyConfiguration) WithSource(value *BuildVolumeSourceApplyConfiguration) *BuildVolumeApplyConfiguration {
	b.Source = value
	return b
}

// WithMounts adds the given value to the Mounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Mounts field.
func (b *BuildVolumeApplyConfiguration) WithMounts(values ...*BuildVolumeMountApplyConfiguration) *BuildVolumeApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMounts")
		}
		b.Mounts = append(b.Mounts, *values[i])
	}
	return b
}
