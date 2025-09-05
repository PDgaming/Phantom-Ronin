#include <raylib.h>

#include "common.h"

void drawPlatform(Platform *platform) {
  DrawCube(platform->Position, platform->Width, platform->Height,
           platform->Length, platform->color);
}