from OpenGL.GL import *
from OpenGL.GLUT import *

def draw_text(text, x, y, font=GLUT_BITMAP_8_BY_13, color=(1.0, 1.0, 1.0, 1.0)):
    glDisable(GL_TEXTURE_2D)
    glDisable(GL_BLEND)

    glColor4f(color[0], color[1], color[2], color[3])

    glPushMatrix()
    glRasterPos2f(x, y)

    for character in text:
        glutBitmapCharacter(font, ord(character))

    glPopMatrix()

    glEnable(GL_TEXTURE_2D)
    glEnable(GL_BLEND)
