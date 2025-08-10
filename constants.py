background_texture_id = None


keys_pressed = {}

# Window dimensions
window_width = 800
window_height = 700

# Camera position and speed factor
camera_pos = [-46.5, 0, 2]
camera_target = [-46.5, 0, 0]
camera_up = [0, 1, 0]
CAMERA_MOVE_SPEED = 0.1

# World dimensions
WORLD_WIDTH = 100
WORLD_HEIGHT = 700
WORLD_DEPTH = 1000

# Coordinate conversion constants
# OpenGL world space dimensions for rendering
GL_WORLD_WIDTH = 20.0  # 20 OpenGL units wide
GL_WORLD_HEIGHT = 8.0  # 8 OpenGL units tall
GL_WORLD_DEPTH = 10.0  # 10 OpenGL units deep

# Conversion factors (pixels to OpenGL units)
PIXELS_TO_GL_X = GL_WORLD_WIDTH / WORLD_WIDTH
PIXELS_TO_GL_Y = GL_WORLD_HEIGHT / WORLD_HEIGHT
PIXELS_TO_GL_Z = GL_WORLD_DEPTH / WORLD_DEPTH

# Level tracking
max_levels = 2
current_level = 1