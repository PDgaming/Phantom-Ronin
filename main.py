import sys
from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *

from constants import *

def draw_axes(length=2.0):
    glBegin(GL_LINES)
    # X axis in red
    glColor3f(1, 0, 0)
    glVertex3f(0, 0, 0)
    glVertex3f(length, 0, 0)
    # Y axis in green
    glColor3f(0, 1, 0)
    glVertex3f(0, 0, 0)
    glVertex3f(0, length, 0)
    # Z axis in blue
    glColor3f(0, 0, 1)
    glVertex3f(0, 0, 0)
    glVertex3f(0, 0, length)
    glEnd()

def draw_cube(
    x=0, y=0, z=0,
    width=1.0, height=1.0, depth=1.0,
    color=(1.0, 0.5, 0.0),
    rotation=(0, 0, 1, 0)
):
    """
    Draws a rectangular cuboid at (x, y, z) with given dimensions, color, and rotation.
    rotation: tuple (angle_degrees, x_axis, y_axis, z_axis)
    """
    glPushMatrix()
    glTranslatef(x, y, z)
    if rotation is not None:
        angle, rx, ry, rz = rotation
        if angle != 0:
            glRotatef(angle, rx, ry, rz)
    glScalef(width, height, depth)
    glColor3f(*color)
    glutSolidCube(1.0)
    glPopMatrix()

def draw_plane(
    x=0, y=0, z=0,
    width=1.0, depth=1.0,
    color=(0.7, 0.7, 0.7),
    rotation=(0, 0, 1, 0)
):
    """
    Draws a rectangular plane at (x, y, z) with given width, depth, color, and rotation.
    rotation: tuple (angle_degrees, x_axis, y_axis, z_axis)
    """
    glPushMatrix()
    glTranslatef(x, y, z)
    if rotation is not None:
        angle, rx, ry, rz = rotation
        if angle != 0:
            glRotatef(angle, rx, ry, rz)
    glScalef(width, 1.0, depth)
    glColor3f(*color)
    glBegin(GL_QUADS)
    glVertex3f(-0.5, 0, -0.5)
    glVertex3f(-0.5, 0, 0.5)
    glVertex3f(0.5, 0, 0.5)
    glVertex3f(0.5, 0, -0.5)
    glEnd()
    glPopMatrix()

def draw_ground():
    glColor3f(0.3, 0.7, 0.3)
    glBegin(GL_QUADS)
    glVertex3f(-5, 0, -5)
    glVertex3f(-5, 0, 5)
    glVertex3f(5, 0, 5)
    glVertex3f(5, 0, -5)
    glEnd()

def display():
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glLoadIdentity()
    # Camera looking at the origin from above and to the side
    gluLookAt(5, 5, 8, 0, 0, 0, 0, 1, 0)

    # Draw ground
    draw_ground()

    # Draw axes
    draw_axes(2.0)

    # Draw a cube at the origin
    draw_cube(x=3, y=0.5, z=0, width=1.0, height=1.0, depth=1.0, color=(1.0, 0.5, 0.0))
    # Draw a vertical plane perpendicular to the ground (e.g., in the XZ plane at y=0.5)
    draw_plane(
        x=0, y=0.5, z=0,
        width=2.0, depth=1.0,
        color=(0.7, 0.7, 0.9),
        rotation=(90, 1, 0, 0)  # Rotate 90 degrees around X to make it vertical
    )

    # Draw a second cube
    draw_cube(x=2, y=0.5, z=2, width=1.0, height=1.0, depth=1.0, color=(0.2, 0.2, 1.0))

    glutSwapBuffers()

def reshape(width, height):
    global window_width, window_height
    window_width = width
    window_height = height
    glViewport(0, 0, width, height)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    gluPerspective(60.0, float(width)/float(height if height != 0 else 1), 0.1, 100.0)
    glMatrixMode(GL_MODELVIEW)

def idle():
    glutPostRedisplay()

def init():
    glEnable(GL_DEPTH_TEST)
    glClearColor(0.2, 0.3, 0.4, 1.0)
    glEnable(GL_COLOR_MATERIAL)
    glEnable(GL_LIGHTING)
    glEnable(GL_LIGHT0)
    glLightfv(GL_LIGHT0, GL_POSITION, [10, 10, 10, 1])
    glLightfv(GL_LIGHT0, GL_DIFFUSE, [1, 1, 1, 1])
    glLightfv(GL_LIGHT0, GL_SPECULAR, [1, 1, 1, 1])

def main():
    glutInit(sys.argv)
    glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB | GLUT_DEPTH)
    glutInitWindowSize(window_width, window_height)
    glutInitWindowPosition(100, 100)
    glutCreateWindow(b"Phantom Ronin 3D Scene")

    init()

    glutDisplayFunc(display)
    glutReshapeFunc(reshape)
    glutIdleFunc(idle)

    glutMainLoop()

if __name__ == "__main__":
    main()
