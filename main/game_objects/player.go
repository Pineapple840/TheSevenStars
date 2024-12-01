components {
  id: "player"
  component: "/main/scripts/player.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"players\"\n"
  "mask: \"enemies\"\n"
  "mask: \"walls\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 6.0\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"Box\"\n"
  "  }\n"
  "  data: 2.0\n"
  "  data: 3.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Main Ship - Base - Full health\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 48.0\n"
  "  y: 48.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game resources/images/player_ship.atlas\"\n"
  "}\n"
  ""
  position {
    x: 6.0
  }
}
embedded_components {
  id: "bullet_producer"
  type: "factory"
  data: "prototype: \"/main/game_objects/bullet.go\"\n"
  ""
}