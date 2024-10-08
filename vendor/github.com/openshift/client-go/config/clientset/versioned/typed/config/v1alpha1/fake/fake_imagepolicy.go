// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/api/config/v1alpha1"
	configv1alpha1 "github.com/openshift/client-go/config/applyconfigurations/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImagePolicies implements ImagePolicyInterface
type FakeImagePolicies struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var imagepoliciesResource = v1alpha1.SchemeGroupVersion.WithResource("imagepolicies")

var imagepoliciesKind = v1alpha1.SchemeGroupVersion.WithKind("ImagePolicy")

// Get takes name of the imagePolicy, and returns the corresponding imagePolicy object, and an error if there is any.
func (c *FakeImagePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagepoliciesResource, c.ns, name), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// List takes label and field selectors, and returns the list of ImagePolicies that match those selectors.
func (c *FakeImagePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImagePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagepoliciesResource, imagepoliciesKind, c.ns, opts), &v1alpha1.ImagePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImagePolicyList{ListMeta: obj.(*v1alpha1.ImagePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImagePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imagePolicies.
func (c *FakeImagePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagepoliciesResource, c.ns, opts))

}

// Create takes the representation of a imagePolicy and creates it.  Returns the server's representation of the imagePolicy, and an error, if there is any.
func (c *FakeImagePolicies) Create(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.CreateOptions) (result *v1alpha1.ImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagepoliciesResource, c.ns, imagePolicy), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// Update takes the representation of a imagePolicy and updates it. Returns the server's representation of the imagePolicy, and an error, if there is any.
func (c *FakeImagePolicies) Update(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.UpdateOptions) (result *v1alpha1.ImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagepoliciesResource, c.ns, imagePolicy), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImagePolicies) UpdateStatus(ctx context.Context, imagePolicy *v1alpha1.ImagePolicy, opts v1.UpdateOptions) (*v1alpha1.ImagePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagepoliciesResource, "status", c.ns, imagePolicy), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// Delete takes name of the imagePolicy and deletes it. Returns an error if one occurs.
func (c *FakeImagePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(imagepoliciesResource, c.ns, name, opts), &v1alpha1.ImagePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImagePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImagePolicyList{})
	return err
}

// Patch applies the patch and returns the patched imagePolicy.
func (c *FakeImagePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied imagePolicy.
func (c *FakeImagePolicies) Apply(ctx context.Context, imagePolicy *configv1alpha1.ImagePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ImagePolicy, err error) {
	if imagePolicy == nil {
		return nil, fmt.Errorf("imagePolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(imagePolicy)
	if err != nil {
		return nil, err
	}
	name := imagePolicy.Name
	if name == nil {
		return nil, fmt.Errorf("imagePolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagepoliciesResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeImagePolicies) ApplyStatus(ctx context.Context, imagePolicy *configv1alpha1.ImagePolicyApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ImagePolicy, err error) {
	if imagePolicy == nil {
		return nil, fmt.Errorf("imagePolicy provided to Apply must not be nil")
	}
	data, err := json.Marshal(imagePolicy)
	if err != nil {
		return nil, err
	}
	name := imagePolicy.Name
	if name == nil {
		return nil, fmt.Errorf("imagePolicy.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagepoliciesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.ImagePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImagePolicy), err
}
