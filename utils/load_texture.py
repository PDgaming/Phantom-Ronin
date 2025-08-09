from PIL import Image
from OpenGL.GL import *

def load_texture(image_path):
    new_texture_id = None
    try:
        image = Image.open(image_path)
        image = image.transpose(Image.FLIP_TOP_BOTTOM)

        if image.mode != "RGBA":
            image = image.convert("RGBA")

        img_data = image.tobytes("raw", "RGBA", 0, -1)
        width, height = image.size

        new_texture_id = glGenTextures(1)
        glBindTexture(GL_TEXTURE_2D, new_texture_id)

        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST)
        glTexParameterf(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST)

        glTexImage2D(
            GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0,
            GL_RGBA, GL_UNSIGNED_BYTE, img_data
        )

        return new_texture_id

    except FileNotFoundError:
        print(f"Error: Image file {image_path} not found")
        return None
    except Exception as error:
        print(f"Error loading image {image_path}: {error}")
        return None
