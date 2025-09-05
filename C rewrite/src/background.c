#include "common.h"
#include <raylib.h>
#include <rlgl.h>
void DrawCubeTextureRec(Texture2D texture, Rectangle source, Vector3 position,
                        float width, float height, float length, Color color) {
  float x = position.x;
  float y = position.y;
  float z = position.z;
  float texWidth = (float)texture.width;
  float texHeight = (float)texture.height;

  // Set desired texture to be enabled while drawing following vertex data
  rlSetTexture(texture.id);

  // We calculate the normalized texture coordinates for the desired
  // texture-source-rectangle It means converting from (tex.width, tex.height)
  // coordinates to [0.0f, 1.0f] equivalent
  rlBegin(RL_QUADS);
  rlColor4ub(color.r, color.g, color.b, color.a);

  // Front face
  rlNormal3f(0.0f, 0.0f, 1.0f);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z + length / 2);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z + length / 2);

  // Back face
  rlNormal3f(0.0f, 0.0f, -1.0f);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z - length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z - length / 2);

  // Top face
  rlNormal3f(0.0f, 1.0f, 0.0f);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z - length / 2);

  // Bottom face
  rlNormal3f(0.0f, -1.0f, 0.0f);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z + length / 2);

  // Right face
  rlNormal3f(1.0f, 0.0f, 0.0f);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z - length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z - length / 2);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x + width / 2, y + height / 2, z + length / 2);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x + width / 2, y - height / 2, z + length / 2);

  // Left face
  rlNormal3f(-1.0f, 0.0f, 0.0f);
  rlTexCoord2f(source.x / texWidth, (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z - length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth,
               (source.y + source.height) / texHeight);
  rlVertex3f(x - width / 2, y - height / 2, z + length / 2);
  rlTexCoord2f((source.x + source.width) / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z + length / 2);
  rlTexCoord2f(source.x / texWidth, source.y / texHeight);
  rlVertex3f(x - width / 2, y + height / 2, z - length / 2);

  rlEnd();

  rlSetTexture(0);
}

void drawBackground(Background *background) {
  background->Position.x = 0.0f + (background->Width / 2) - 0.25f;

  if (!background->TextureProvided) {
    DrawCube(background->Position, background->Width, background->Height,
             background->Length, background->color);
  } else {
    DrawCubeTextureRec(background->texture,
                       (Rectangle){0.0f, 0.0f, background->texture.width,
                                   background->texture.height},
                       background->Position, background->Width, 2.0f, 2.0f,
                       WHITE);
  }
}
