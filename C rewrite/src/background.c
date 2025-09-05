#include <raylib.h>

#include "common.h"

void drawBackground(Background *background) {
  background->Position.x = 0.0f + (background->Width / 2) - 0.25f;

  DrawCube(background->Position, background->Width, background->Height,
           background->Length, background->color);
}