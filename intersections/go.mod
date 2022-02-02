module intersections

go 1.17

replace rays => ../rays

replace tuples => ../tuples

require (
	rays v0.0.0-00010101000000-000000000000
	spheres v0.0.0-00010101000000-000000000000
	transformations v0.0.0-00010101000000-000000000000
)

require (
	matrices v0.0.0-00010101000000-000000000000 // indirect
	tuples v0.0.0-00010101000000-000000000000
)

replace spheres => ../objects/spheres

replace matrices => ../matrices

replace transformations => ../transformations
