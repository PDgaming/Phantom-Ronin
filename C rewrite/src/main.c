#include <stdio.h>
#include <string.h>

#define RAYGUI_IMPLEMENTATION
#include "../include/raygui.h"
#include <raylib.h>
#include <raymath.h>

#include "common.h"

#define screenWidth 800
#define screenHeight 480

typedef struct {
  int Level;
  bool isSideView;
  bool isDebug;
  char menuState[20];
} GameState;

Color backgroundColor = {255, 182, 193, 255};
Level *currentLevel;
bool exitButton;
bool startButton;
bool transitionButton;

void resetGame(GameState *state, Player *player, Level *currentLevel) {
  player->Position = (Vector3){0.0f, -1.0f, 0.0f};
  player->Velocity = (Vector3){0.0f, 0.0f, 0.0f};
  player->IsGrounded = true;
  player->jumpsUsed = 0;
  resetLevel(currentLevel);

  switch (state->Level) {
  case 1:
    loadLevel(currentLevel, "./level-maps/level1.csv");
    break;
  case 2:
    loadLevel(currentLevel, "./level-maps/level2.csv");
    break;
  default:
    printf("Game Completed!");
    state->Level = 0;
    strcpy(state->menuState, "gameOver");
  }
}

int main() {
  InitWindow(screenWidth, screenHeight, "Phantom-Ronin");

  GameState state = {1, true, false, "startMenu"};

  Camera3D camera = {0};
  camera.position = (Vector3){0.0f, 0.0f, 10.0f};
  camera.target = (Vector3){0.0f, 0.0f, 0.0f};
  camera.up = (Vector3){0.0f, 1.0f, 0.0f};
  camera.fovy = 45.0f;
  camera.projection = CAMERA_PERSPECTIVE;

  Texture2D backgroundTexture = LoadTexture("../assets/background.png");

  Background background = {
      {0.0f, 0.0f, -1.0f}, screenHeight, WORLD_WIDTH, 0.1f, BLUE, true,
      backgroundTexture};

  Ground ground = {{0.0f, -2.0f, 0.1f}, WORLD_WIDTH, 0.2f, 2.0f, RED};

  Player player = {
      {0.0f, -1.0f, 0.0f}, 0.5,   1.0f, 0.5f, GREEN, {0.0f, 0.0f, 0.0f},
      {0.0f, 0.0f, 0.0f},  false, 0,    0};

  Wall leftWall = {{-0.5f, -0.6f, 0.0f}, 0.5f, 3.0f, 2.2f, DARKBROWN};

  Wall rightWall = {{WORLD_WIDTH, -0.6f, 0.0f}, 0.5f, 3.0f, 2.2f, DARKBROWN};

  currentLevel = createLevel();
  resetGame(&state, &player, currentLevel);

  SetTargetFPS(120);
  while (!WindowShouldClose()) {
    if (strcmp(state.menuState, "inGame") == 0 ||
        strcmp(state.menuState, "gameOver") == 0) {
      if (IsKeyPressed(KEY_R)) {
        state.isSideView = !state.isSideView;
      }

      updatePlayer(&player, state.isSideView, &background, &ground);
    }

    BoundingBox playerBox =
        (BoundingBox){{player.Position.x - player.Width / 2,
                       player.Position.y - player.Height / 2,
                       player.Position.z - player.Length / 2},
                      {player.Position.x + player.Width / 2,
                       player.Position.y + player.Height / 2,
                       player.Position.z + player.Length / 2}};
    BoundingBox groundBox =
        (BoundingBox){{ground.Position.x - ground.Width / 2,
                       ground.Position.y - ground.Height / 2,
                       ground.Position.z - ground.Length / 2},
                      {ground.Position.x + ground.Width / 2,
                       ground.Position.y + ground.Height / 2,
                       ground.Position.z + ground.Length / 2}};

    BoundingBox leftWallBox =
        (BoundingBox){{leftWall.Position.x - leftWall.Width / 2,
                       leftWall.Position.y - leftWall.Height / 2,
                       leftWall.Position.z - leftWall.Length / 2},
                      {leftWall.Position.x + leftWall.Width / 2,
                       leftWall.Position.y + leftWall.Height / 2,
                       leftWall.Position.z + leftWall.Length / 2}};

    BoundingBox rightWallBox =
        (BoundingBox){{rightWall.Position.x - rightWall.Width / 2,
                       rightWall.Position.y - rightWall.Height / 2,
                       rightWall.Position.z - rightWall.Length / 2},
                      {rightWall.Position.x + rightWall.Width / 2,
                       rightWall.Position.y + rightWall.Height / 2,
                       rightWall.Position.z + rightWall.Length / 2}};

    if (CheckCollisionBoxes(playerBox, groundBox)) {
      if (player.Velocity.y <= 0) {
        player.IsGrounded = true;
        player.jumpsUsed = 0;
        player.Velocity.y = 0.0f;

        player.Position.y =
            ground.Position.y + (ground.Height / 2) + (player.Height / 2);
      } else {
        player.Velocity.y = 0.0f;
        player.Position.y =
            ground.Position.y - (ground.Height / 2) - (player.Height / 2);
      }
    } else {
      player.IsGrounded = false;
    }

    if (CheckCollisionBoxes(playerBox, leftWallBox)) {
      player.Position.x =
          leftWall.Position.x + leftWall.Width / 2 + player.Width / 2;
    }
    if (CheckCollisionBoxes(playerBox, rightWallBox)) {
      player.Position.x =
          rightWall.Position.x - rightWall.Width / 2 - player.Width / 2;
    }

    for (int i = 0; i < currentLevel->count; i++) {
      Platform *platform = &currentLevel->Platforms[i];

      BoundingBox platformBox =
          (BoundingBox){{platform->Position.x - platform->Width / 2,
                         platform->Position.y - platform->Height / 2,
                         platform->Position.z - platform->Length / 2},
                        {platform->Position.x + platform->Width / 2,
                         platform->Position.y + platform->Height / 2,
                         platform->Position.z + platform->Length / 2}};

      if (CheckCollisionBoxes(playerBox, platformBox)) {
        float playerBottom = player.Position.y - player.Height / 2;
        float platformTop = platform->Position.y + platform->Height / 2;
        float platformBottom = platform->Position.y - platform->Height / 2;

        if (playerBottom >= platformTop - 0.1f) {
          player.Position.y = platformTop + player.Height / 2;
          player.IsGrounded = true;
          player.jumpsUsed = 0;
          player.Velocity.y = 0.0f;

          if (platform->final) {
            printf("Level Completed!\n");
            printf("Transitioning to Level %d\n", state.Level);
            strcpy(state.menuState, "levelTransition");
          }
        } else if (player.Position.y + player.Height / 2 <=
                       platformBottom + 0.05 &&
                   player.Velocity.y > 0) {
          player.Position.y = platformBottom - player.Height / 2;
          player.Velocity.y = 0.0f;
        } else {
          float playerTop = player.Position.y + player.Height / 2;

          if (playerTop > platformBottom && playerBottom < platformTop) {
            if (state.isSideView) {
              float playerLeft = player.Position.x - player.Width / 2;
              float playerRight = player.Position.x + player.Width / 2;
              float platformLeft = platform->Position.x - platform->Width / 2;
              float platformRight = platform->Position.x + platform->Width / 2;

              float overlapLeft = playerRight - platformLeft;
              float overlapRight = platformRight - playerLeft;

              if (overlapLeft < overlapRight) {
                player.Position.x -= overlapLeft;
              } else {
                player.Position.x += overlapRight;
              }
            } else {
              float playerFront = player.Position.z + player.Length / 2;
              float playerBack = player.Position.z - player.Length / 2;
              float platformFront = platform->Position.z + platform->Length / 2;
              float platformBack = platform->Position.z - platform->Length / 2;

              float overlapBack = playerFront - platformBack;
              float overlapFront = platformFront - playerBack;

              if (overlapBack < overlapFront) {
                player.Position.z -= overlapBack;
              } else {
                player.Position.z += overlapFront;
              }
            }
          }
        }
      }
    }

    if (state.isSideView) {
      float clampX = Clamp(player.Position.x, 3.15f, background.Width - 3.7f);
      float clampY =
          Clamp(player.Position.y, 0.1f, background.Height - player.Height);

      camera.position = (Vector3){clampX, clampY, 6.0f};
      camera.target = (Vector3){clampX, clampY, player.Position.z};
    } else {
      camera.position =
          (Vector3){player.Position.x + 5, player.Position.y + 2, 4.0f};
      camera.target = (Vector3){player.Position.x, player.Position.y, 0.0f};
    }

    BeginDrawing();
    ClearBackground(backgroundColor);

    BeginMode3D(camera);

    if (state.isDebug) {
      DrawBoundingBox(playerBox, RED);
      DrawBoundingBox(groundBox, GREEN);
      DrawBoundingBox(leftWallBox, BLUE);
      DrawBoundingBox(rightWallBox, BLUE);
    }

    drawBackground(&background);
    drawGround(&ground);
    drawWall(&leftWall);
    drawWall(&rightWall);

    drawLevel(currentLevel);

    drawPlayer(&player);

    EndMode3D();

    if (strcmp(state.menuState, "startMenu") == 0) {
      DrawText("Phanton Ronin", 80, 150, 80, RED);
      startButton = GuiButton(
          (Rectangle){screenWidth / 2.0f - 50, 250, 100, 40}, "Start");

      if (startButton) {
        strcpy(state.menuState, "inGame");
        resetGame(&state, &player, currentLevel);
      }

      exitButton =
          GuiButton((Rectangle){screenWidth / 2.0f - 50, 300, 100, 40}, "Exit");

      if (exitButton) {
        break;
      }
    }

    if (strcmp(state.menuState, "levelTransition") == 0) {
      DrawText("Level Completed!", 80, 150, 80, RED);
      transitionButton =
          GuiButton((Rectangle){screenWidth / 2.0f - 50, 250, 100, 40}, "Next");

      if (transitionButton) {
        strcpy(state.menuState, "inGame");
        state.Level++;
        resetGame(&state, &player, currentLevel);
      }
    }

    if (strcmp(state.menuState, "gameOver") == 0) {
      DrawText("Game Completed!", 70, 190, 80, RED);
      exitButton =
          GuiButton((Rectangle){screenWidth / 2.0f - 50, 280, 100, 40}, "Exit");

      if (exitButton) {
        break;
      }
    }

    if (state.isDebug) {
      DrawText(TextFormat("Player: %.2f, %.2f, %.2f", player.Position.x,
                          player.Position.y, player.Position.z),
               10, 40, 18, RED);
      DrawText(TextFormat("Camera: %.2f, %.2f, %.2f", camera.position.x,
                          camera.position.y, camera.position.z),
               10, 60, 18, RED);
      DrawText(TextFormat("Level: %d", state.Level), 10, 80, 18, RED);
    }

    DrawText(TextFormat("Level: %d", state.Level), 10, 30, 18, RED);

    DrawFPS(10, 10);
    EndDrawing();
  }

  CloseWindow();
  freeLevel(currentLevel);

  return 0;
}