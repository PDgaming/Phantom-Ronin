import OpenGL
from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *
import time

# Import texture loading utility
from utils.load_texture import *

# Window dimensions
window_width = 800
window_height = 600

# World dimensions - adjusting height to match screen aspect ratio better
WORLD_WIDTH = 4000
WORLD_HEIGHT = 800  # Reduced from 800 to match window height better

# Camera position - positioned to look directly at the background plane
camera_x = 0.0
camera_y = 0  # Lowered to 0 so we only see from the bottom of the background
camera_z = 0.0  # Same Z as the background plane (-200) + offset for viewing

# Movement state
keys_pressed = {}

# Game state
last_time = time.time()

# Texture ID for background
background_texture_id = None

def load_background_texture():
    """Load the background texture"""
    global background_texture_id
    background_image_path = "assets/background.png"
    background_texture_id = load_texture(background_image_path)
    if background_texture_id is None:
        print("Failed to load background texture")
        return False
    return True

def draw_background_plane():
    """Draw the background plane perpendicular to the ground"""
    if background_texture_id is None:
        # Fallback: draw a colored plane - using a distinct color for debugging
        glColor3f(0.8, 0.2, 0.8)  # Magenta color for debugging
    else:
        # Enable texture
        glEnable(GL_TEXTURE_2D)
        glBindTexture(GL_TEXTURE_2D, background_texture_id)
        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)
        glColor3f(1.0, 1.0, 1.0)  # White for texture

    # Draw the background plane perpendicular to the ground (Y-Z plane)
    # Position it behind the scene and make it as wide as WORLD_WIDTH
    plane_width = WORLD_WIDTH
    plane_height = WORLD_HEIGHT

    glBegin(GL_QUADS)
    if background_texture_id:
        glTexCoord2f(0.0, 1.0)  # Fixed texture coordinates for proper orientation
    glVertex3f(-plane_width/2, 0, -200)  # Bottom left
    if background_texture_id:
        glTexCoord2f(4.0, 1.0)  # Repeat texture 4 times horizontally
    glVertex3f(plane_width/2, 0, -200)   # Bottom right
    if background_texture_id:
        glTexCoord2f(4.0, 0.0)
    glVertex3f(plane_width/2, plane_height, -200)  # Top right
    if background_texture_id:
        glTexCoord2f(0.0, 0.0)
    glVertex3f(-plane_width/2, plane_height, -200)  # Top left
    glEnd()

    if background_texture_id:
        glDisable(GL_TEXTURE_2D)
        glDisable(GL_BLEND)

    glColor3f(1.0, 1.0, 1.0)  # Reset to white

def draw_floor():
    """Draw a stretched floor cube"""
    glPushMatrix()
    # Position at the bottom and stretch it to cover the full width
    glTranslatef(0, 198, 0)  # Position at the bottom
    
    # Scale the cube to stretch it into a floor
    floor_width = WORLD_WIDTH
    floor_height = 50  # Height of the floor
    floor_depth = 200  # Depth of the floor
    
    glScalef(floor_width / 50.0, floor_height / 50.0, floor_depth / 50.0)  # Scale the cube
    
    glEnable(GL_BLEND)
    glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)
    glColor4f(0.5, 0.3, 0.2, 0.7)  # Brown color for floor, slightly transparent
    glutSolidCube(50)
    glPopMatrix()
    glColor3f(1.0, 1.0, 1.0)  # Reset to white

def setup_3d_scene():
    """Set up the 3D scene with proper lighting"""
    glEnable(GL_DEPTH_TEST)
    glEnable(GL_LIGHTING)
    glEnable(GL_LIGHT0)
    glEnable(GL_COLOR_MATERIAL)
    
    # Set up light
    glLightfv(GL_LIGHT0, GL_POSITION, [1, 1, 1, 0])
    glLightfv(GL_LIGHT0, GL_AMBIENT, [0.2, 0.2, 0.2, 1.0])
    glLightfv(GL_LIGHT0, GL_DIFFUSE, [0.8, 0.8, 0.8, 1.0])

def setup_camera():
    """Set up the camera to look directly at the background plane (2D-like view)"""
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    
    # Calculate the aspect ratio to maintain proper proportions
    aspect_ratio = float(window_width) / float(window_height)
    
    # Scale the view to make the background cover the entire screen
    # We'll use a zoom factor to scale down the view
    zoom_factor = 4.0  # Increase this to zoom in more
    
    # Calculate view dimensions to fit the screen properly
    view_width = WORLD_WIDTH / zoom_factor
    view_height = WORLD_HEIGHT / zoom_factor
    
    # Adjust view height to match the window aspect ratio exactly
    if aspect_ratio > (view_width / view_height):
        # Window is wider - adjust width to maintain aspect ratio
        view_width = view_height * aspect_ratio
    else:
        # Window is taller - adjust height to maintain aspect ratio
        view_height = view_width / aspect_ratio
    
    glOrtho(-view_width/2, view_width/2, 0, view_height, -1000, 1000)
    
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    
    # Position camera to look directly at the background plane
    # Camera is positioned at the same Z as the background plane (-200) but slightly forward
    camera_z_pos = -150  # Slightly in front of the background plane (-200)
    
    gluLookAt(
        camera_x, camera_y, camera_z_pos,     # Eye position
        camera_x, camera_y, -200,             # Look at point (the background plane)
        0, 1, 0                              # Up vector
    )

def handle_input():
    """Handle keyboard input for movement"""
    global camera_x
    
    speed = 5.0  # Movement speed
    
    if keys_pressed.get(b'a', False) or keys_pressed.get('a', False):
        camera_x -= speed
    if keys_pressed.get(b'd', False) or keys_pressed.get('d', False):
        camera_x += speed
    
    # Limit camera movement to stay within the viewport
    view_width = WORLD_WIDTH / 4.0  # Same as the zoom factor calculation
    max_camera_x = view_width / 2
    min_camera_x = -view_width / 2
    camera_x = max(min_camera_x, min(camera_x, max_camera_x))

def display():
    """Main display function"""
    global last_time

    current_time = time.time()
    dt = current_time - last_time
    last_time = current_time

    handle_input()

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()
    
    setup_camera()
    setup_3d_scene()
    
    # Draw the scene
    draw_background_plane()  # Draw background first (behind everything)
    draw_floor()  # Draw the floor
    
    glutSwapBuffers()

def reshape(width, height):
    """Handle window reshape"""
    global window_width, window_height
    window_width = width
    window_height = height
    glViewport(0, 0, width, height)

def keyboard_down(key, x, y):
    """Handle key press"""
    char_key = key.decode("utf-8").lower() if isinstance(key, bytes) else key.lower()
    keys_pressed[char_key] = True

def keyboard_up(key, x, y):
    """Handle key release"""
    char_key = key.decode("utf-8").lower() if isinstance(key, bytes) else key.lower()
    keys_pressed[char_key] = False

def idle():
    """Idle function for animation"""
    glutPostRedisplay()

def main():
    """Main function"""
    global background_texture_id
    
    glutInit()
    glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB | GLUT_DEPTH)
    glutInitWindowSize(window_width, window_height)
    glutInitWindowPosition(100, 100)
    glutCreateWindow(b"Phantom Ronin")
    
    # Set background color
    glClearColor(0.2, 0.3, 0.4, 1.0)
    
    # Load background texture
    if load_background_texture():
        print("Background texture loaded successfully!")
    else:
        print("Using magenta background plane (no texture) - for debugging")
    
    glutDisplayFunc(display)
    glutReshapeFunc(reshape)
    glutIdleFunc(idle)
    glutKeyboardFunc(keyboard_down)
    glutKeyboardUpFunc(keyboard_up)
    
    glutMainLoop()

if __name__ == "__main__":
    main()
