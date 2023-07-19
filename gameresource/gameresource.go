package gameresource

import "path"

type Resource struct {
	Item    *Item
	Episode *Episode
}

func NewFromPath(resourcePath string) (*Resource, error) {
	resource := &Resource{}
	return resource, resource.LoadWithPath(resourcePath)
}

func (r *Resource) LoadWithPath(resourcePath string) error {
	if r.Item == nil {
		r.Item = &Item{}
	}
	if err := r.Item.LoadFromJson(path.Join(resourcePath, "Json/item.json")); err != nil {
		return err
	}

	if r.Episode == nil {
		r.Episode = &Episode{}
	}
	if err := r.Episode.LoadFromJson(path.Join(resourcePath, "Json/episode.json")); err != nil {
		return err
	}

	return nil
}
