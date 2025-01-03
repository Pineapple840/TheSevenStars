components {
  id: "enemy_general"
  component: "/main/scripts/enemy_general.script"
}
components {
  id: "enemy2"
  component: "/main/scripts/enemy2.script"
}
components {
  id: "bullet_sound"
  component: "/game resources/audio/sfx/laserSmall_001.ogg"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game_images/enemy2_weapons.tilesource\"\n"
  "}\n"
  ""
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
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 12.1303\n"
  "  data: 12.218407\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "bullet_producer"
  type: "factory"
  data: "prototype: \"/main/game_objects/enemy2_bullet.go\"\n"
  ""
}
