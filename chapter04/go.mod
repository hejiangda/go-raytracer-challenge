module chapter04

go 1.17

replace canvas => ../canvas

replace tuples => ../tuples

replace matrices => ../matrices

replace transformations => ../transformations

require (
	canvas v0.0.0-00010101000000-000000000000
	transformations v0.0.0-00010101000000-000000000000
	tuples v0.0.0-00010101000000-000000000000
)

require matrices v0.0.0-00010101000000-000000000000
