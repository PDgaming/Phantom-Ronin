import sys
from PIL import Image

from OpenGL.GL import *
from OpenGL.GLUT import *
from OpenGL.GLU import *

from constants import *
from utils.keyboard_handles import *
from utils.coords_convertion import *

def draw_cube(
    x=0, y=0, z=0,
    width=1.0, height=1.0, depth=1.0,
    color=(1.0, 0.5, 0.0, 1),
    rotation=(0, 0, 1, 0),
    use_pixels=False
):
    """
    Draws a rectangular cuboid at (x, y, z) with given dimensions, color (including alpha), and rotation.
    
    Args:
        x, y, z: Position coordinates (pixels if use_pixels=True, OpenGL units if False)
        width, height, depth: Dimensions (pixels if use_pixels=True, OpenGL units if False)
        color: tuple (r, g, b, a) where a is alpha (0.0 transparent, 1.0 opaque)
        rotation: tuple (angle_degrees, x_axis, y_axis, z_axis)
        use_pixels: If True, coordinates and dimensions are in pixels; if False, in OpenGL units
    """
    try:
        # Convert pixel coordinates to OpenGL coordinates if needed
        if use_pixels:
            gl_x, gl_y, gl_z = pixels_to_gl_coords(x, y, z)
            gl_width, gl_height, gl_depth = pixels_to_gl_size(width, height, depth)
        else:
            gl_x, gl_y, gl_z = x, y, z
            gl_width, gl_height, gl_depth = width, height, depth
        
        # Ensure we're in the right matrix mode
        glMatrixMode(GL_MODELVIEW)
        glPushMatrix()
        glTranslatef(gl_x, gl_y, gl_z)
        if rotation is not None:
            angle, rx, ry, rz = rotation
            if angle != 0:
                glRotatef(angle, rx, ry, rz)
        glScalef(gl_width, gl_height, gl_depth)

        # Enable blending and set up for transparency
        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)
        glEnable(GL_ALPHA_TEST)
        glAlphaFunc(GL_GREATER, 0.0)

        # Disable depth writing for proper transparency rendering
        glDepthMask(GL_FALSE)
        glColor4f(*color)
        glutSolidCube(1.0)
        glDepthMask(GL_TRUE)

        glDisable(GL_ALPHA_TEST)
        glDisable(GL_BLEND)
        glPopMatrix()
    except Exception as e:
        print(f"Error in draw_cube: {e}")
        # Ensure matrix stack is popped if there was an error
        try:
            glPopMatrix()
        except:
            pass


def load_background_texture():
    global background_texture_id
    if background_texture_id is not None:
        return background_texture_id

    try:
        image = Image.open("assets/background.png")
        image = image.convert("RGBA")
        image_data = image.tobytes("raw", "RGBA", 0, -1)
        width, height = image.size

        background_texture_id = glGenTextures(1)
        glBindTexture(GL_TEXTURE_2D, background_texture_id)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
        glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)
        glTexImage2D(
            GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0,
            GL_RGBA, GL_UNSIGNED_BYTE, image_data
        )
        glBindTexture(GL_TEXTURE_2D, 0)
        return background_texture_id
    except Exception as e:
        print("Failed to load texture:", e)
        return None

def draw_background(width=GL_WORLD_WIDTH):
    """
    Draws a background quad with the background texture,
    repeating the texture horizontally to fill the specified width.
    The quad covers the entire world in world coordinates, so that
    the camera/view/projection still works.

    Args:
        width (float): The width of the background in OpenGL world units.
                       The texture will repeat to fill this width, but each
                       repeat will be at the original scale of the texture.
    """
    try:
        tex_id = load_background_texture()

        glPushMatrix()
        glDisable(GL_DEPTH_TEST)
        glEnable(GL_BLEND)
        glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)
        glEnable(GL_TEXTURE_2D)

        # Disable lighting so the background is uniformly bright
        lighting_was_enabled = glIsEnabled(GL_LIGHTING)
        if lighting_was_enabled:
            glDisable(GL_LIGHTING)

        if tex_id is not None:
            glBindTexture(GL_TEXTURE_2D, tex_id)
            glColor4f(1.0, 1.0, 1.0, 1.0)
            # Set texture wrapping to repeat
            glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_REPEAT)
            glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_REPEAT)
        else:
            glColor4f(0.7, 0.7, 0.7, 1.0)

        # Draw a single quad in world coordinates at the far back (z = -GL_WORLD_DEPTH/2)
        z = -GL_WORLD_DEPTH / 2 + 0.01  # Slightly in front of far plane to avoid z-fighting

        # Calculate how many times to repeat the texture horizontally
        # The texture coordinates should be set to this value to use GL_REPEAT
        repeat_x = width / GL_WORLD_WIDTH
        
        half_width = width / 2.0
        half_height = GL_WORLD_HEIGHT / 2.0

        glBegin(GL_QUADS)
        # Lower-left
        glTexCoord2f(0.0, 0.0)
        glVertex3f(-half_width, -half_height, z)
        # Lower-right
        glTexCoord2f(repeat_x, 0.0)
        glVertex3f(half_width, -half_height, z)
        # Upper-right
        glTexCoord2f(repeat_x, 1.0)
        glVertex3f(half_width, half_height, z)
        # Upper-left
        glTexCoord2f(0.0, 1.0)
        glVertex3f(-half_width, half_height, z)
        glEnd()

        if tex_id is not None:
            glBindTexture(GL_TEXTURE_2D, 0)
        glDisable(GL_TEXTURE_2D)
        glDisable(GL_BLEND)
        glEnable(GL_DEPTH_TEST)

        # Restore lighting state if it was enabled before
        if lighting_was_enabled:
            glEnable(GL_LIGHTING)

        glPopMatrix()
    except Exception as e:
        print(f"Error in draw_background: {e}")
        try:
            glPopMatrix()
        except:
            pass

def update_camera():
    """
    Moves the camera in the x and y axes based on 'a' and 'd' keys.
    """
    move_x = 0.0
    move_y = 0.0
    # 'a' moves left (negative x), 'd' moves right (positive x)
    if keys_pressed.get('a', False):
        move_x -= CAMERA_MOVE_SPEED
    if keys_pressed.get('d', False):
        move_x += CAMERA_MOVE_SPEED
    # Optionally, you can add 'w'/'s' for y axis movement
    # if keys_pressed.get('w', False):
    #     move_y += CAMERA_MOVE_SPEED
    # if keys_pressed.get('s', False):
    #     move_y -= CAMERA_MOVE_SPEED

    # Move both camera_pos and camera_target to keep the view direction
    if move_x != 0.0 or move_y != 0.0:
        camera_pos[0] += move_x
        camera_pos[1] += move_y
        camera_target[0] += move_x
        camera_target[1] += move_y

    print(camera_pos)
    print(camera_target)

def display():
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glEnable(GL_BLEND)
    glBlendFunc(GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA)
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()

    # Update camera position based on keyboard input
    update_camera()

    # Camera looking at the origin from above and to the side
    gluLookAt(camera_pos[0], 0, 1.9,
            camera_target[0], camera_target[1], camera_target[2],
            camera_up[0], camera_up[1], camera_up[2])

    # Ensure matrix stack is in a clean state before drawing
    try:
        # Draw background using pixel coordinates (covers the entire world)
        draw_background(width=WORLD_WIDTH)
        
        # draw_cube(
        #     x=1000, y=400, z=0,  # Pixel coordinates
        #     width=100, height=150, depth=0,  # Pixel dimensions
        #     color=(1.0, 0.5, 0.0, 0.5),
        #     rotation=(0, 0, 0, 0),
        #     use_pixels=True
        # )
    except Exception as e:
        print(f"Error in display: {e}")
        # Reset matrix stack if there was an error
        glMatrixMode(GL_MODELVIEW)
        glLoadIdentity()
    
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
    
    # Initialize matrix stack properly
    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()

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
    glutKeyboardFunc(keyboard_down)
    glutKeyboardUpFunc(keyboard_up)

    glutMainLoop()

if __name__ == "__main__":
    main()
