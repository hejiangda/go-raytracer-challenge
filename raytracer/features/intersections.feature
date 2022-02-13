Scenario: An intersection encapsulates t and object
  Given s <- sphere()
  When i <- intersection(3.5,s)
  Then i.t = 3.5
    And i.object = s

Scenario: Aggregating intersections
  Given s <- sphere()
    And i1 <- intersection(1,s)
    And i2 <- intersection(2,s)
  When xs <- intersections(i1,i2)
  Then xs.count = 2
    And xs[0].t=1
    And xs[1].t=2

Scenario: Intersect sets the object on the intersection
  Given r <- ray(point(0,0,-5),vector(0,0,1))
    And s <- sphere()
  When xs <- intersect(s,r)
  Then xs.count =2
    And xs[0].object =s
    And xs[1].object =s

Scenario: The hit, when all intersections have positive t
  Given s <- sphere()
    And i1 <- intersection(1,s)
    And i2 <- intersection(2,s)
    And xs <- intersection(i2,i1)
  When i <- hit(xs)
  Then i = i1

Scenario: The hit, when some intersections have negative t
  Given s <- sphere()
    And i1 <- intersection(-1,s)
    And i2 <- intersection(1,s)
    And xs <- intersections(i2,i1)
  When i <- hit(xs)
  Then i=i2

Scenario: The hit, when all intersections have negative t
  Given s <- sphere()
    And i1 <- intersection(-2,s)
    And i2 <- intersection(-1,s)
    And xs <- intersections(i2,i1)
  When i <- hit(xs)
  Then i is nothing

Scenario: The hit is always the lowest nonnegative intersection
  Given s <- sphere()
    And i1 <- intersection(5,s)
    And i2 <- intersection(7,s)
    And i3 <- intersection(-3,s)
    And i4 <- intersection(2,s)
    And xs <- intersections(i1,i2,i3,i4)
  When i <- hit(xs)
  Then i = i4

Scenario: Intersecting a scaled sphere with a ray
  Given r <- ray(point(0,0,-5),vector(0,0,1))
    And s <- sphere()
  When set_transform(s,scaling(2,2,2))
    And xs <- intersect(s,r)
  Then xs.count = 2
    And xs[0].t = 3
    And xs[1].t = 7

Scenario: Intersecting a translated sphere with a ray
  Given r <- ray(point(0,0,-5),vector(0,0,1))
    And s <- sphere()
  When set_transform(s,translation(5,0,0))
    And xs <- intersect(s,r)
  Then xs.count = 0
