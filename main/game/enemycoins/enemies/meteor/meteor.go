components {
  id: "rock"
  component: "/main/game/enemycoins/enemies/meteor/meteor.script"
}
components {
  id: "greyExplode"
  component: "/assets/vfx/greyExplode.particlefx"
}
components {
  id: "greenishExplode"
  component: "/assets/vfx/greenishExplode.particlefx"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"meteor_green\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game/sprites.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.18
    y: 0.18
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"player\"\n"
  "mask: \"shield\"\n"
  "mask: \"rocket\"\n"
  "mask: \"radar\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      x: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 21.0\n"
  "}\n"
  ""
}
