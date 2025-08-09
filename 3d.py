# Basic OpenGL 3D scene using PyOpenGL and GLUT
# Now with mouse controls for changing the viewport
# Camera yaw/pitch directions fixed, WASD movement implemented
# Now with V and SPACE to move up and down in the z direction

from OpenGL.GL import *
from OpenGL.GLU import *
from OpenGL.GLUT import *
import sys
import math

# Window dimensions
window_width = 800
window_height = 600

# Camera parameters
angle = 0.0
delta_angle = 0.0

# Camera position in world space (eye position)
cam_pos_x = 0.0
cam_pos_y = 0.0
cam_pos_z = 5.0

# Mouse control state
mouse_left_down = False
mouse_last_x = 0
mouse_last_y = 0
camera_yaw = 0.0    # Horizontal angle (left/right)
camera_pitch = 0.0  # Vertical angle (up/down)
camera_distance = 5.0

# Movement state
move_forward = False
move_backward = False
move_left = False
move_right = False
move_up = False      # SPACE
move_down = False    # V

def init():
    glClearColor(0.1, 0.1, 0.2, 1.0)
    glEnable(GL_DEPTH_TEST)
    glShadeModel(GL_SMOOTH)
    glEnable(GL_COLOR_MATERIAL)
    glEnable(GL_LIGHTING)
    glEnable(GL_LIGHT0)
    glLightfv(GL_LIGHT0, GL_POSITION, [1, 1, 1, 0])
    glLightfv(GL_LIGHT0, GL_AMBIENT, [0.2, 0.2, 0.2, 1.0])
    glLightfv(GL_LIGHT0, GL_DIFFUSE, [0.8, 0.8, 0.8, 1.0])

def update_camera_position():
    """Update camera position based on WASD movement and camera's local axes (yaw and pitch).
    WASD only moves in X and Y (horizontal plane), SPACE/V move in Z (vertical)."""
    global cam_pos_x, cam_pos_y, cam_pos_z
    speed = 0.2  # Increased speed for more noticeable movement

    # Calculate forward vector (local z axis), but ignore vertical (Y) for WASD
    yaw_rad = math.radians(camera_yaw)
    # Only use yaw for horizontal movement
    forward_x = math.sin(yaw_rad)
    forward_y = 0.0
    forward_z = -math.cos(yaw_rad)

    # Calculate right vector (local x axis)
    right_x = math.cos(yaw_rad)
    right_y = 0.0
    right_z = math.sin(yaw_rad)

    # Move along local axes (horizontal plane only for WASD)
    if move_forward:
        cam_pos_x += forward_x * speed
        cam_pos_z += forward_z * speed
    if move_backward:
        cam_pos_x -= forward_x * speed
        cam_pos_z -= forward_z * speed
    if move_left:
        cam_pos_x -= right_x * speed
        cam_pos_z -= right_z * speed
    if move_right:
        cam_pos_x += right_x * speed
        cam_pos_z += right_z * speed
    # Move up and down in world Y axis (vertical) using SPACE and V
    if move_up:
        cam_pos_y += speed
        print(f"Moving up: cam_pos_y = {cam_pos_y}")  # Debug print
    if move_down:
        cam_pos_y -= speed
        print(f"Moving down: cam_pos_y = {cam_pos_y}")  # Debug print

def display():
    global camera_yaw, camera_pitch, camera_distance, cam_pos_x, cam_pos_y, cam_pos_z
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()

    update_camera_position()

    # Calculate camera direction from yaw, pitch, and distance
    yaw_rad = math.radians(camera_yaw)
    pitch_rad = math.radians(camera_pitch)
    # Camera look direction (unit vector)
    dir_x = math.sin(yaw_rad) * math.cos(pitch_rad)
    dir_y = math.sin(pitch_rad)
    dir_z = -math.cos(yaw_rad) * math.cos(pitch_rad)
    # Camera eye position
    eye_x = cam_pos_x
    eye_y = cam_pos_y
    eye_z = cam_pos_z
    # Camera look-at target
    center_x = eye_x + dir_x
    center_y = eye_y + dir_y
    center_z = eye_z + dir_z

    gluLookAt(
        eye_x, eye_y, eye_z,
        center_x, center_y, center_z,
        0.0, 1.0, 0.0
    )

    # Draw a ground plane
    glColor3f(0.3, 0.7, 0.3)
    glBegin(GL_QUADS)
    glVertex3f(-5, -1, -5)
    glVertex3f(-5, -1, 5)
    glVertex3f(5, -1, 5)
    glVertex3f(5, -1, -5)
    glEnd()

    # Draw a colored cube in the center
    glPushMatrix()
    glTranslatef(0.0, 0.0, 0.0)
    glRotatef(angle, 0, 1, 0)
    glColor3f(0.8, 0.2, 0.2)
    glutSolidCube(1.0)
    glPopMatrix()

    # Draw a sphere
    glPushMatrix()
    glTranslatef(2.0, 0.0, 0.0)
    glColor3f(0.2, 0.2, 0.8)
    glutSolidSphere(0.5, 32, 32)
    glPopMatrix()

    glutSwapBuffers()

def reshape(width, height):
    if height == 0:
        height = 1
    glViewport(0, 0, width, height)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    gluPerspective(45, float(width) / float(height), 0.1, 100.0)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()

def idle():
    global angle
    angle += 0.2
    if angle > 360:
        angle -= 360
    glutPostRedisplay()

def keyboard(key, x_pos, y_pos):
    global camera_distance, move_forward, move_backward, move_left, move_right, move_up, move_down
    key = key.lower()
    if key == b'\x1b' or key == b'q':
        sys.exit()
    elif key == b'w':
        move_forward = True
    elif key == b's':
        move_backward = True
    elif key == b'a':
        move_left = True
    elif key == b'd':
        move_right = True
    elif key == b' ':
        move_up = True

def keyboard_up(key, x_pos, y_pos):
    global move_forward, move_backward, move_left, move_right, move_up, move_down
    key = key.lower()
    if key == b'w':
        move_forward = False
    elif key == b's':
        move_backward = False
    elif key == b'a':
        move_left = False
    elif key == b'd':
        move_right = False
    elif key == b' ':
        move_up = False

def special_key(key, x_pos, y_pos):
    """No longer used for vertical movement."""
    pass

def special_key_up(key, x_pos, y_pos):
    """No longer used for vertical movement."""
    pass

def mouse(button, state, x, y):
    global mouse_left_down, mouse_last_x, mouse_last_y
    if button == GLUT_LEFT_BUTTON:
        if state == GLUT_DOWN:
            mouse_left_down = True
            mouse_last_x = x
            mouse_last_y = y
        elif state == GLUT_UP:
            mouse_left_down = False
    elif button == 3:  # Scroll up
        zoom_camera(-0.5)
    elif button == 4:  # Scroll down
        zoom_camera(0.5)

def mouse_motion(x, y):
    global mouse_left_down, mouse_last_x, mouse_last_y
    global camera_yaw, camera_pitch
    if mouse_left_down:
        dx = x - mouse_last_x
        dy = y - mouse_last_y
        # Fix the direction of yaw and pitch and reduce sensitivity
        sensitivity = 0.09  # Lower sensitivity (was 0.5)
        camera_yaw += dx * sensitivity  # X movement increases yaw (right = positive dx)
        camera_pitch -= dy * sensitivity  # Y movement decreases pitch (up = negative dy)
        # Clamp pitch to avoid flipping
        if camera_pitch > 89.0:
            camera_pitch = 89.0
        if camera_pitch < -89.0:
            camera_pitch = -89.0
        mouse_last_x = x
        mouse_last_y = y
        glutPostRedisplay()

def zoom_camera(amount):
    global camera_distance
    camera_distance += amount
    if camera_distance < 1.0:
        camera_distance = 1.0
    if camera_distance > 50.0:
        camera_distance = 50.0
    glutPostRedisplay()

def main():
    glutInit(sys.argv)
    glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB | GLUT_DEPTH)
    glutInitWindowSize(window_width, window_height)
    glutInitWindowPosition(100, 100)
    glutCreateWindow(b"Basic OpenGL 3D Scene")
    init()
    glutDisplayFunc(display)
    glutReshapeFunc(reshape)
    glutIdleFunc(idle)
    glutKeyboardFunc(keyboard)
    glutKeyboardUpFunc(keyboard_up)
    glutSpecialFunc(special_key)
    glutSpecialUpFunc(special_key_up)
    glutMouseFunc(mouse)
    glutMotionFunc(mouse_motion)
    glutMainLoop()

if __name__ == "__main__":
    main()
