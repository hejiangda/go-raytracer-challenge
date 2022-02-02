module chapter06

go 1.17

replace canvas => ../canvas

replace tuples => ../tuples

replace matrices => ../matrices

replace transformations => ../transformations

require (
	canvas v0.0.0-00010101000000-000000000000
	intersections v0.0.0-00010101000000-000000000000
	matrices v0.0.0-00010101000000-000000000000
	spheres v0.0.0-00010101000000-000000000000
	transformations v0.0.0-00010101000000-000000000000
	tuples v0.0.0-00010101000000-000000000000
)

require rays v0.0.0-00010101000000-000000000000

require lights v0.0.0-00010101000000-000000000000

replace spheres => ../objects/spheres

replace rays => ../rays

replace intersections => ../intersections

replace lights => ../lights
