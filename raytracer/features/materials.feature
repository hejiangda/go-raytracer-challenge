Background:
  Given m ← material()
    And position ← point(0, 0, 0)

Scenario: Lighting with the eye between the light and the surface
  Given eyev ← vector(0, 0, -1)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv)
  Then result = color(1.9, 1.9, 1.9)

Scenario: Lighting with the eye between light and surface, eye offset 45°
  Given eyev ← vector(0, √2/2, -√2/2)
    And normalv ← vector(0, 0, -1)
    And light ← point_light(point(0, 0, -10), color(1, 1, 1))
  When result ← lighting(m, light, position, eyev, normalv)
  Then result = color(1.0, 1.0, 1.0)

Scenario: Lighting with eye opposite surface, light offset 45°
 	  Given eyev ← vector(0, 0, -1)
 	    And normalv ← vector(0, 0, -1)
 	    And light ← point_light(point(0, 10, -10), color(1, 1, 1))
 	  When result ← lighting(m, light, position, eyev, normalv)
 	  Then result = color(0.7364, 0.7364, 0.7364)

Scenario: Lighting with eye in the path of the reflection vector
 	  Given eyev ← vector(0, -√2/2, -√2/2)
 	    And normalv ← vector(0, 0, -1)
 	    And light ← point_light(point(0, 10, -10), color(1, 1, 1))
 	  When result ← lighting(m, light, position, eyev, normalv)
 	  Then result = color(1.6364, 1.6364, 1.6364)

Scenario: Lighting with the light behind the surface
 	  Given eyev ← vector(0, 0, -1)
 	    And normalv ← vector(0, 0, -1)
 	    And light ← point_light(point(0, 0, 10), color(1, 1, 1))
 	  When result ← lighting(m, light, position, eyev, normalv)
 	  Then result = color(0.1, 0.1, 0.1)

Scenario: Lighting with the surface in shadow
	Given eyev <- vector(0,0,-1)
	  And normalv <- vector(0,0,-1)
	  And light <- point_light(point(0,0,-10),color(1,1,1))
	  And in_shadow <- true
	When result <- lighting(m,light,position,eyev,normalv,in_shadow)
	Then result = color(0.1,0.1,0.1)

Scenario: Reflectivity for the default material
 	  Given m ← material()
 	  Then m.reflective = 0.0

Scenario: Transparency and Refractive Index for the default material
  Given m ← material()
  Then m.transparency = 0.0
	And m.refractive_index = 1.0
