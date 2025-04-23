components {
  id: "laser"
  component: "/main/game/enemycoins/enemies/alien_red/laser.script"
}
components {
  id: "greenExplode"
  component: "/assets/vfx/greenExplode.particlefx"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"laser\"\n"
  "mask: \"player\"\n"
  "mask: \"shield\"\n"
  "mask: \"rocket\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"shoot\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game/sprites.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.01
    y: 0.01
  }
}
