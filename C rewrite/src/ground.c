#include <raylib.h>

#include "common.h"

void drawGround(Ground *ground) {
  ground->Position.x = 0.0f + (ground->Width / 2) - 0.25f;

  DrawCube(ground->Position, ground->Width, ground->Height, ground->Length,
           ground->color);
}