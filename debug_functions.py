from OpenGL.GL import *

from constants import *

def draw_debug_bbox(x, y, width, height, color):
    glDisable(GL_TEXTURE_2D)
    glDisable(GL_BLEND)

    glColor4f(color[0], color[1], color[2], color[3])
    glLineWidth(2.0)

    glBegin(GL_LINE_LOOP)
    glVertex2f(x, y)
    glVertex2f(x + width, y)
    glVertex2f(x + width, y + height)
    glVertex2f(x, y + height)
    glEnd()

    glLineWidth(1.0)


def draw_debug_grid(grid_spacing, color):
    glDisable(GL_TEXTURE_2D)
    glDisable(GL_BLEND)

    glColor4f(color[0], color[1], color[2], color[3])
    glLineWidth(1.0)

    glBegin(GL_LINES)

    start_x_on_screen = -int(camera_x % grid_spacing)
    for x in range(start_x_on_screen, window_width + grid_spacing, grid_spacing):
        glVertex2f(x, 0.0)
        glVertex2f(x, window_height)

    for y in range(0, window_height + grid_spacing, grid_spacing):
        glVertex2f(0.0, y)
        glVertex2f(window_width, y)

    glEnd()
    glLineWidth(1.0)
