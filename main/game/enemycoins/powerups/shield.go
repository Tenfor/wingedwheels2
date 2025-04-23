components {
  id: "shield"
  component: "/main/game/enemycoins/powerups/shield.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"shield\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.9
  }
  scale {
    x: 0.1
    y: 0.1
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"shield\"\n"
  "mask: \"enemy\"\n"
  "mask: \"coin\"\n"
  "mask: \"laser\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "    id: \"hitbox\"\n"
  "  }\n"
  "  data: 64.0\n"
  "}\n"
  ""
}
