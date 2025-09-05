#include <raylib.h>

#include "common.h"

void drawWall(Wall *wall) {
  DrawCube(wall->Position, wall->Width, wall->Height, wall->Length,
           wall->color);
}