# projectile starts one unit above the origin.
# velocity is normalized to 1 unit/tick.
  p <- projectile(point(0,1,0),normalize(vector(1,1,0)))

# gravity -0.1 unit/tick, and wind is -0.01 unit/tick.
  e <- environment(vector(0,-0.1,0),vector(-0.01,0,0))

