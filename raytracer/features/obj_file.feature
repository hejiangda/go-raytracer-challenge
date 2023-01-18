Scenario: Ignoring unrecognized lines
  Given gibberish <- a file containing:
    """
    There was a young lady named Bright
    who traveled much faster than light.
    She set out one day
    in a relative way,
    and came back the previous night.
    """
  When parser <- parse_obj_file(gibberish)
  Then parser should have ignored 5 lines

Scenario: Vertex records
  Given file <- a file containing:
    """
    v -1 1 0
    v -1.0000 0.5000 0.0000
    v 1 0 0
    v 1 1 0
    """
  When parser <- parse_obj_file(file)
  Then parser.vertices[1] = point(-1,1,0)
    And parser.vertices[2] = point(-1,0.5,0)
    And parser.vertices[3] = point(1,0,0)
    And parser.vertices[4] = point(1,1,0)

Scenario: Parsing triangle faces
  Given file <- a file containing:
    """
    v -1 1 0
    v -1 0 0
    v 1 0 0
    v 1 1 0

    f 1 2 3
    f 1 3 4
    """
  When parser <- parse_obj_file(file)
    And g <- parser.default_group
    And t1 <- first child of g
    And t2 <- second child of g
  Then t1.p1 = parser.vertices[1]
    And t1.p2 = parser.vertices[2]
    And t1.p3 = parser.vertices[3]
    And t2.p1 = parser.vertices[1]
    And t2.p2 = parser.vertices[3]
    And t2.p3 = parser.vertices[4]

