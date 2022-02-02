module spheres

go 1.17

replace tuples => ../../tuples

replace rays => ../../rays

replace intersections => ../../intersections

replace matrices => ../../matrices

replace transformations => ../../transformations

require (
	rays v0.0.0-00010101000000-000000000000
	transformations v0.0.0-00010101000000-000000000000
	tuples v0.0.0-00010101000000-000000000000
)

require (
	lights v0.0.0-00010101000000-000000000000
	matrices v0.0.0-00010101000000-000000000000
)

replace lights => ../../lights
