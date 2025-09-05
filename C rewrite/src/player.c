#include <raylib.h>
#include <raymath.h>

#include "common.h"

#define JUMP_STRENGTH 5.0f
#define SPEED 8.0f

void drawPlayer(Player *player) {
  DrawCube(player->Position, player->Width, player->Height, player->Length,
           player->color);
}

void updatePlayer(Player *player, bool isSideView, Background *background,
                  Ground *ground) {
  player->Velocity.x = 0.0f;
  player->Velocity.z = 0.0f;

  if (isSideView) {
    if (IsKeyDown(KEY_A)) {
      player->Velocity.x = -SPEED;
    }
    if (IsKeyDown(KEY_D)) {
      player->Velocity.x = SPEED;
    }
  } else {
    if (IsKeyDown(KEY_A)) {
      player->Velocity.z = SPEED;
    }
    if (IsKeyDown(KEY_D)) {
      player->Velocity.z = -SPEED;
    }
  }

  if (IsKeyPressed(KEY_SPACE)) {
    if (player->IsGrounded) {
      player->Velocity.y = JUMP_STRENGTH;
      player->jumpsUsed = 1;
      player->IsGrounded = false;
    } else if (player->jumpsUsed == 1) {
      player->Velocity.y = JUMP_STRENGTH;
      player->jumpsUsed = 2;
    }
  }

  if (!player->IsGrounded) {
    player->Velocity.y += GRAVITY * GetFrameTime();
  }

  player->Velocity.x *= 0.5f;

  player->Position.x += player->Velocity.x * GetFrameTime();
  player->Position.y += player->Velocity.y * GetFrameTime();
  player->Position.z += player->Velocity.z * GetFrameTime();

  float minZ = ground->Position.z - ground->Length / 2 + player->Length / 2;
  float maxZ = ground->Position.z + ground->Length / 2 - player->Length / 2;

  float clampedZ = Clamp(player->Position.z, minZ, maxZ);

  player->Position.z = clampedZ;
}