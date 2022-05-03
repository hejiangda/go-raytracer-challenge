package raytracer

type Material struct {
	Color           *Tuple
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Pattern         Pattern
	Reflective      float64
	Transparency    float64
	RefractiveIndex float64
}

func NewMaterial() *Material {
	return &Material{Color(1, 1, 1), 0.1, 0.9, 0.9, 200, nil, 0.0, 0, 1}
}
func (m *Material) Equal(n *Material) bool {
	if m.Color.Equal(n.Color) &&
		AlmostEqual(m.Ambient, n.Ambient, Eps) &&
		AlmostEqual(m.Diffuse, n.Diffuse, Eps) &&
		AlmostEqual(m.Specular, n.Specular, Eps) &&
		AlmostEqual(m.Shininess, n.Shininess, Eps) &&
		AlmostEqual(m.Reflective, n.Reflective, Eps) {
		return true
	}
	return false
}
