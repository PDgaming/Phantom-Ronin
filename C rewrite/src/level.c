#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <raylib.h>

#include "common.h"

Level *createLevel() {
  Level *level = malloc(sizeof(Level));
  level->Platforms = NULL;
  level->count = 0;
  level->capacity = 0;

  return level;
}

void freeLevel(Level *level) {
  if (level) {
    free(level->Platforms);
    free(level);
  }
}

void loadLevel(Level *level, const char *filePath) {
  FILE *file = fopen(filePath, "r");
  if (!file) {
    TraceLog(LOG_ERROR, "Failed to open level file: %s", filePath);
    return;
  }

  char line[256];
  int lineNum = 0;

  if (fgets(line, sizeof(line), file)) {
    lineNum++;
  }

  while (fgets(line, sizeof(line), file)) {
    lineNum++;

    line[strcspn(line, "\n")] = 0;

    char *token;
    char *saveptr;
    float values[7];
    int i = 0;

    token = strtok(line, ",");
    while (token && i < 7) {
      values[i] = atof(token);
      token = strtok(NULL, ",");
      i++;
    }

    if (i < 7) {
      continue;
    }

    if (level->count >= level->capacity) {
      level->capacity = level->capacity == 0 ? 10 : level->capacity * 2;
      level->Platforms =
          realloc(level->Platforms, level->capacity * sizeof(Platform));
    }

    Platform *platform = &level->Platforms[level->count];
    platform->Position = (Vector3){values[0], values[1], values[2]};
    platform->Width = values[3];
    platform->Height = values[4];
    platform->Length = values[5];
    platform->final = (values[6] != 0.0f);
    platform->color = BROWN;

    level->count++;
  }

  fclose(file);
}

void resetLevel(Level *level) { level->count = 0; }

void drawLevel(Level *currentLevel) {
  for (int i = 0; i < currentLevel->count; i++) {
    drawPlatform(&currentLevel->Platforms[i]);
  }
}