# Player object
player = None

# List of game objects
game_objects = []

# Texture IDs for background and level exit
background_texture_id = None
platform_texture_id = None
level_exit_texture_id = None

# Dictionary to store pressed keys
keys_pressed = {}

# Time tracking
last_time = 0.0

# Debug-Info
DEBUG_DRAW_BBOX = False
DEBUG_DRAW_GRID = False

# FPS
SHOW_FPS = True
fps = 0
frame_count = 0
fps_last_update_time = 0.0

# Window dimensions
window_width = 800
window_height = 700

# Physics constants
GRAVITY = 980.0
JUMP_FORCE = 350.0
MAX_JUMPS = 2

# Camera position and speed factor
camera_x = 0.0
camera_y = 0.0
CAMERA_SPEED_FACTOR = 0.5

# Viewport dimensions
WORLD_WIDTH = 4000
WORLD_HEIGHT = 800

# Level tracking
max_levels = 2
current_level = 1