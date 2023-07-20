package gameresource

import "path"

type Resource struct {
	Jump    *Jump
	Item    *Item
	Episode *Episode
}

func NewFromPath(resourcePath string) (*Resource, error) {
	resource := &Resource{}
	return resource, resource.LoadWithPath(resourcePath)
}

func (r *Resource) LoadWithPath(resourcePath string) error {

	if r.Jump == nil {
		r.Jump = &Jump{}
	}
	if err :=  r.Jump.LoadFromJson(path.Join(resourcePath, "Json/jump.json")); err != nil {
		return err
	}

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
