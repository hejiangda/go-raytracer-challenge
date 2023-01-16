package raytracer

type Shape interface {
	GetTransform() *Matrix
	SetTransform(*Matrix)

	GetMaterial() *Material
	SetMaterial(*Material)

	Intersect(r *Ray) (ret []Intersection)
	NormalAt(p *Tuple) *Tuple

	GetParent() Shape
	SetParent(s Shape)
	World2Object(p *Tuple) *Tuple
	Normal2World(p *Tuple) *Tuple
}

func NewShape(shapeType string) Shape {
	switch shapeType {
	case "sphere":
		return NewSphere()
	case "plane":
		return NewPlane()
	}
	return nil
}
