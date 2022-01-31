Scenario: Trajectory of a projectile
  Given start <- point(0,1,0)
    And velocity <- normalize(vector(1,1.8,0))*11.25
    And p <- projectile(start, velocity)

    And gravity <- vector(0,-0.1,0)
    And wind <- vector(-0.01,0,0)
    And e <- environment(gravity,wind)

    And c <- canvas(900,550)
