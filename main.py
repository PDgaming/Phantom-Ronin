import OpenGL
from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *

import time
import csv

from constants import *
from keyboard_handles import *
from debug_functions import *
from utils.draw_text import *
from utils.load_texture import *
from utils.check_aabb_collision import *

class Player:
    def __init__(self, x, y, width, height, speed, texture_opengl_id, bbox_offset_x=0, bbox_offset_y=0, bbox_width=None, bbox_height=None):
        self.x = x
        self.y = y
        self.sprite_width = width # Renamed to avoid confusion with bbox_width
        self.sprite_height = height # Renamed to avoid confusion with bbox_height
        self.speed = speed
        self.texture_id = texture_opengl_id
        self.vx = 0.0
        self.vy = 0.0
        self.on_ground = False
        self.facing_right = True
        self.jump_count = 0

        # Bounding Box
        self.bbox_offset_x = bbox_offset_x
        self.bbox_offset_y = bbox_offset_y
        self.bbox_width = bbox_width if bbox_width is not None else self.sprite_width
        self.bbox_height = bbox_height if bbox_height is not None else self.sprite_height # FIX: Corrected variable name

    def get_bbox(self):
        bbox_x = self.x + self.bbox_offset_x
        bbox_y = self.y + self.bbox_offset_y
        return bbox_x, bbox_y, self.bbox_width, self.bbox_height

    def update(self, dt):
        global camera_x

        # Apply gravity to vertical velocity
        self.vy -= GRAVITY * dt

        # Determine horizontal velocity based on input
        self.vx = 0.0
        if keys_pressed.get("a"):
            self.vx = -self.speed
            self.facing_right = False
        if keys_pressed.get("d"):
            self.vx = self.speed
            self.facing_right = True

        # Handle jump input
        if keys_pressed.get(" "):
            if self.jump_count < MAX_JUMPS:
                self.vy = JUMP_FORCE
                self.jump_count += 1
                keys_pressed[" "] = False # Consume the key press
                self.on_ground = False # Player is now in air

        # --- Collision Resolution (Axis-by-Axis) ---

        # 1. Resolve X-axis movement
        potential_x = self.x + self.vx * dt
        # Temporarily update player's x for horizontal collision check
        self.x = potential_x

        player_bbox_x, player_bbox_y, player_bbox_width, player_bbox_height = self.get_bbox()

        # Iterate over a temporary copy to handle collection modifications safely.
        # This prevents issues if `transition_to_next_level` clears `game_objects`.
        for obj in list(game_objects):
            obj_bbox_x, obj_bbox_y, obj_bbox_width, obj_bbox_height = obj.get_bbox()
            # Check collision using player's new X and current Y
            if check_aabb_collision(player_bbox_x, player_bbox_y, player_bbox_width, player_bbox_height,
                                    obj_bbox_x, obj_bbox_y, obj_bbox_width, obj_bbox_height):
                
                # If we collide with a LevelExit, handle the transition.
                if isinstance(obj, LevelExit):
                    transition_to_next_level()
                    # The `transition_to_next_level` function already clears game_objects
                    # and resets the player position, so we should exit here.
                    return 

                # If we collide with a regular platform, resolve the collision as usual.
                # Collision detected on X-axis, resolve by nudging player out
                if self.vx > 0: # Moving right, hit left side of platform
                    self.x = obj_bbox_x - player_bbox_width - self.bbox_offset_x
                elif self.vx < 0: # Moving left, hit right side of platform
                    self.x = obj_bbox_x + obj_bbox_width - self.bbox_offset_x
                self.vx = 0 # Stop horizontal movement

        # Clamp X position to world boundaries after horizontal collisions
        self.x = max(0, min(self.x, WORLD_WIDTH - self.sprite_width))


        # 2. Resolve Y-axis movement
        potential_y = self.y + self.vy * dt
        # Temporarily update player's y for vertical collision check
        self.y = potential_y

        # Reset on_ground status before re-evaluating for this frame
        self.on_ground = False
        player_bbox_x, player_bbox_y, player_bbox_width, player_bbox_height = self.get_bbox() # Re-get bbox with updated X

        for obj in game_objects:
            obj_bbox_x, obj_bbox_y, obj_bbox_width, obj_bbox_height = obj.get_bbox()
            # Check collision using player's current X and new Y
            if check_aabb_collision(player_bbox_x, player_bbox_y, player_bbox_width, player_bbox_height,
                                    obj_bbox_x, obj_bbox_y, obj_bbox_width, obj_bbox_height):
                # We've already handled `LevelExit` above, so this must be a platform.
                # Collision detected on Y-axis, resolve by nudging player out
                if self.vy > 0: # Moving up, hit bottom side of platform (ceiling)
                    self.y = obj_bbox_y - player_bbox_height - self.bbox_offset_y
                    self.vy = 0 # Stop upward movement
                elif self.vy < 0: # Moving down, hit top side of platform (floor)
                    self.y = obj_bbox_y + obj_bbox_height - self.bbox_offset_y
                    self.vy = 0 # Stop downward movement
                    self.on_ground = True # Player has landed
                    self.jump_count = 0 # Reset jump count on landing

        # Clamp Y position to world boundaries after vertical collisions
        min_world_y_clamped = 190 # Absolute floor height (e.g., matching ground level of background)
        self.y = max(min(self.y, WORLD_HEIGHT - self.sprite_height), min_world_y_clamped)

        # If after all collisions, player is at min_world_y_clamped, ensure on_ground is set
        if self.y == min_world_y_clamped:
            self.on_ground = True
            self.jump_count = 0
    

        # --- End Collision Resolution ---

        # Update camera position based on player's world position
        target_camera_x = self.x - window_width / 2
        camera_x = max(0, min(target_camera_x, WORLD_WIDTH - window_width))

    def draw(self):
        if self.texture_id is None:
            print("Warning: Player texture not loaded")
            return

        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, self.texture_id)

        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)

        glBegin(GL_QUADS)

        tex_s_left = 0.0
        tex_s_right = 1.0

        if not self.facing_right:
            tex_s_left = 1.0
            tex_s_right = 0.0

        glTexCoord2f(tex_s_left, 0.0)
        glVertex2f(self.x - camera_x, self.y + self.sprite_height) # Use sprite_height

        glTexCoord2f(tex_s_right, 0.0)
        glVertex2f(self.x + self.sprite_width - camera_x, self.y + self.sprite_height) # Use sprite_width, sprite_height

        glTexCoord2f(tex_s_right, 1.0)
        glVertex2f(self.x + self.sprite_width - camera_x, self.y) # Use sprite_width

        glTexCoord2f(tex_s_left, 1.0)
        glVertex2f(self.x - camera_x, self.y)

        glEnd()

        glDisable(GL_TEXTURE_2D)
        glDisable(GL_BLEND)

        if DEBUG_DRAW_BBOX:
            bbox_x_screen = (self.x + self.bbox_offset_x) - camera_x
            bbox_y_screen = (self.y + self.bbox_offset_y)
            draw_debug_bbox(bbox_x_screen, bbox_y_screen, self.bbox_width, self.bbox_height, (1.0, 0.0, 0.0, 1.0))


class Platform:
    def __init__(self, x, y, width, height, texture_opengl_id, bbox_offset_x=0, bbox_offset_y=0, bbox_width=None, bbox_height=None):
        self.x = x
        self.y = y
        self.sprite_width = width # Renamed to avoid confusion with bbox_width
        self.sprite_height = height # Renamed to avoid confusion with bbox_height
        self.texture_id = texture_opengl_id

        # Bounding Box
        self.bbox_offset_x = bbox_offset_x
        self.bbox_offset_y = bbox_offset_y
        self.bbox_width = bbox_width if bbox_width is not None else self.sprite_width
        self.bbox_height = bbox_height if bbox_height is not None else self.sprite_height # FIX: Corrected variable name

    def get_bbox(self):
        bbox_x = self.x + self.bbox_offset_x
        bbox_y = self.y + self.bbox_offset_y
        return bbox_x, bbox_y, self.bbox_width, self.bbox_height

    def draw(self):
        if self.texture_id is None:
            print(f"Warning: Platform texture not loaded for platform at ({self.x},{self.y}).")
            return

        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, self.texture_id)

        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)

        glBegin(GL_QUADS)

        glTexCoord2f(0.0, 0.0)
        glVertex2f(self.x - camera_x, self.y + self.sprite_height)

        glTexCoord2f(1.0, 0.0)
        glVertex2f(self.x + self.sprite_width - camera_x, self.y + self.sprite_height)

        glTexCoord2f(1.0, 1.0)
        glVertex2f(self.x + self.sprite_width - camera_x, self.y)

        glTexCoord2f(0.0, 1.0)
        glVertex2f(self.x - camera_x, self.y)

        glEnd()

        glDisable(GL_TEXTURE_2D)
        glDisable(GL_BLEND)

        if DEBUG_DRAW_BBOX:
            bbox_x_screen = (self.x + self.bbox_offset_x) - camera_x
            bbox_y_screen = (self.y + self.bbox_offset_y)
            draw_debug_bbox(bbox_x_screen, bbox_y_screen, self.bbox_width, self.bbox_height, (0.0, 1.0, 0.0, 1.0))


class LevelExit:
    def __init__(self, x, y, width, height, texture_opengl_id):
        self.x = x
        self.y = y
        self.sprite_width = width
        self.sprite_height = height
        self.texture_id = texture_opengl_id

        # Bounding Box
        self.bbox_offset_x = 0
        self.bbox_offset_y = 0
        self.bbox_width = self.sprite_width
        self.bbox_height = self.sprite_height

    def get_bbox(self):
        bbox_x = self.x + self.bbox_offset_x
        bbox_y = self.y + self.bbox_offset_y
        return bbox_x, bbox_y, self.bbox_width, self.bbox_height
    
    def draw(self):
        if self.texture_id is None:
            print("Warning: LevelExit texture not loaded.")
            return

        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, self.texture_id)
        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)

        glBegin(GL_QUADS)
        glTexCoord2f(0.0, 0.0); glVertex2f(self.x - camera_x, self.y + self.sprite_height)
        glTexCoord2f(1.0, 0.0); glVertex2f(self.x + self.sprite_width - camera_x, self.y + self.sprite_height)
        glTexCoord2f(1.0, 1.0); glVertex2f(self.x + self.sprite_width - camera_x, self.y)
        glTexCoord2f(0.0, 1.0); glVertex2f(self.x - camera_x, self.y)
        glEnd()

        glDisable(GL_TEXTURE_2D)
        glDisable(GL_BLEND)
        
        if DEBUG_DRAW_BBOX:
            bbox_x_screen = (self.x + self.bbox_offset_x) - camera_x
            bbox_y_screen = (self.y + self.bbox_offset_y)
            draw_debug_bbox(bbox_x_screen, bbox_y_screen, self.bbox_width, self.bbox_height, (0.0, 0.0, 1.0, 1.0))


def draw_background():
    if background_texture_id is None:
        print("Warning: Background texture not loaded")
        return

    glEnable(GL_TEXTURE_2D)
    glBindTexture(GL_TEXTURE_2D, background_texture_id)

    glEnable(GL_BLEND)
    glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)

    glBegin(GL_QUADS)

    s_start = camera_x / WORLD_WIDTH
    s_end = (camera_x + window_width) / WORLD_WIDTH

    t_start = 1.0
    t_end = 0.0

    glTexCoord2f(s_start, t_end)
    glVertex2f(0.0, window_height)

    glTexCoord2f(s_end, t_end)
    glVertex2f(window_width, window_height)

    glTexCoord2f(s_end, t_start)
    glVertex2f(window_width, 0.0)

    glTexCoord2f(s_start, t_start)
    glVertex2f(0.0, 0.0)

    glEnd()

    glDisable(GL_TEXTURE_2D)
    glDisable(GL_BLEND)


def iterate():
    glViewport(0, 0, window_width, window_height)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glOrtho(0.0, window_width, 0.0, window_height, 0.0, 1.0)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()


def showScreen():
    global last_time, fps, frame_count, fps_last_update_time

    current_time = time.time()
    dt = current_time - last_time
    last_time = current_time

    frame_count += 1
    if current_time - fps_last_update_time >= 1.0:
        fps = frame_count
        frame_count = 0
        fps_last_update_time = current_time

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()
    iterate()

    draw_background()

    if DEBUG_DRAW_GRID:
        draw_debug_grid(grid_spacing=50, color=(0.5, 0.5, 0.5, 0.5))

    for obj in game_objects:
        obj.draw()

    if player:
        player.update(dt)
        player.draw()

    if SHOW_FPS:
        fps_text = f"FPS: {fps}"
        text_width = 100

        glPushAttrib(GL_CURRENT_BIT | GL_ENABLE_BIT)

        draw_text(fps_text, x=window_width - text_width, y=window_height-30, color=(0.0, 0.0, 0.0, 1.0))

        glPopAttrib()

    glutSwapBuffers()
    glutPostRedisplay()


def load_level_from_csv(file_path):
    global game_objects, platform_texture_id, level_exit_texture_id

    game_objects.clear()

    try:
        with open(file_path, "r") as csv_file:
            csv_reader = csv.DictReader(csv_file)
            for row in csv_reader:
                obj_type = row["type"]
                x = int(row["x"])
                y = int(row["y"])
                width = int(row["width"])
                height = int(row["height"])

                if obj_type == "platform":
                    game_objects.append(Platform(x=x, y=y, width=width, height=height, texture_opengl_id=platform_texture_id))
                elif obj_type == "levelexit":
                    game_objects.append(LevelExit(x=x, y=y, width=width, height=height, texture_opengl_id=level_exit_texture_id))

    except FileNotFoundError:
        print(f"Error: Level file {file_path} not found")
        return
    except Exception as error:
        print(f"Error loading level file {file_path}: {error}")
        return


def transition_to_next_level():
    global current_level, max_levels, player

    if current_level < max_levels:
        current_level += 1
        level_file = f"levels/level{current_level}.csv"
        print(f"Transitioning to level {current_level}: {level_file}")
        load_level_from_csv(level_file)

        if player:
            player.x = 100
            player.y = 190
            player.vx = 0
            player.vy = 0
    else:
        print("Game Completed!")


if __name__ == "__main__":
    glutInit()
    glutInitDisplayMode(GLUT_RGBA | GLUT_DOUBLE | GLUT_DEPTH)
    glutInitWindowSize(window_width, window_height)
    glutInitWindowPosition(0, 0)
    wind = glutCreateWindow(b"OpenGL 2D Platformer")

    last_time = time.time()

    player_image_path = "assets/sprite.png"
    player_texture_id = load_texture(player_image_path)
    if player_texture_id is None:
        print("Failed to load texture. Exiting...")
        exit()

    background_image_path = "assets/background.jpg"
    background_texture_id = load_texture(background_image_path)
    if background_texture_id is None:
        print("Failed to load texture. Exiting...")
        exit()

    platform_image_path = "assets/platform.png"
    platform_texture_id = load_texture(platform_image_path)
    if platform_texture_id is None:
        print("Failed to load texture. Exiting...")
        exit()

    level_exit_image_path = "assets/flag.png"
    level_exit_texture_id = load_texture(level_exit_image_path)
    if level_exit_texture_id is None:
        print("Failed to load texture. Exiting...")
        exit()

    player = Player(
        x=100, y=190, width=128, height=128, speed=400,
        texture_opengl_id=player_texture_id,
        bbox_offset_x=30, bbox_offset_y=0, bbox_width=64, bbox_height=128
    )

    load_level_from_csv("levels/level1.csv")

    glutDisplayFunc(showScreen)
    glutIdleFunc(showScreen)
    glutKeyboardFunc(keyboard_down)
    glutKeyboardUpFunc(keyboard_up)

    glutMainLoop()
