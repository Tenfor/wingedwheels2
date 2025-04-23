components {
  id: "powerup"
  component: "/main/game/enemycoins/powerups/powerup.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"powerup_magnet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game/sprites.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"coin\"\n"
  "mask: \"player\"\n"
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
  "  data: 22.5\n"
  "}\n"
  ""
}
