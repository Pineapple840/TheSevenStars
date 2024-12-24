components {
  id: "enemy"
  component: "/main/scripts/enemy.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Kla\\\'ed - Scout - Base\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/enemy_sprites.atlas\"\n"
  "}\n"
  ""
  position {
    y: -4.0
  }
  rotation {
    x: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemies\"\n"
  "mask: \"bullets\"\n"
  "mask: \"players\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -2.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 11.136363\n"
  "  data: 11.241059\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
