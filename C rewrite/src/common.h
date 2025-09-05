#ifndef COMMON_H
#define COMMON_H

#include <raylib.h>
#include <raymath.h>

#define WORLD_WIDTH 30.0f
#define WORLD_LENGTH 2.0f

#define GRAVITY -9.8f

typedef struct {
    Vector3 Position;
    float Height;
    float Width;
    float Length;
    Color color;
} Background;

typedef struct {
    Vector3 Position;
    float Width;
    float Height;
    float Length;
    Color color;
} Ground;

typedef struct {
    Vector3 Position;
    float Width;
    float Height;
    float Length;
    Color color;

    Vector3 Velocity;
    Vector3 Acceleration;

    bool IsGrounded;
    int jumpsUsed;
    int State;
} Player;

typedef struct {
    Vector3 Position;
    float Width;
    float Height;
    float Length;
    Color color;
} Wall;

typedef struct {
    Vector3 Position;
    float Width;
    float Height;
    float Length;
    Color color;

    bool final;
} Platform;

typedef struct {
    Platform *Platforms;
    int count;
    int capacity;
} Level;

void drawBackground(Background*);
void drawGround(Ground*);
void drawPlayer(Player*);
void updatePlayer(Player*, bool, Background*, Ground*);
void drawWall(Wall*);
void drawPlatform(Platform*);
Level* createLevel();
void freeLevel(Level*);
void loadLevel(Level*, const char*);
void resetLevel(Level*);
void drawLevel(Level*);
#endif