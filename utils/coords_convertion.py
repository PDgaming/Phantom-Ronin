def pixels_to_gl_coords(pixel_x, pixel_y, pixel_z=0):
    """
    Convert pixel coordinates to OpenGL world coordinates.
    
    Args:
        pixel_x: X coordinate in pixels (0 to WORLD_WIDTH)
        pixel_y: Y coordinate in pixels (0 to WORLD_HEIGHT)
        pixel_z: Z coordinate in pixels (0 to WORLD_DEPTH)
    
    Returns:
        tuple: (gl_x, gl_y, gl_z) in OpenGL world coordinates
    """
    # Convert pixel coordinates to normalized coordinates (0 to 1)
    norm_x = pixel_x / WORLD_WIDTH
    norm_y = pixel_y / WORLD_HEIGHT
    norm_z = pixel_z / WORLD_DEPTH
    
    # Convert to OpenGL world coordinates
    gl_x = (norm_x - 0.5) * GL_WORLD_WIDTH  # Center at origin
    gl_y = (0.5 - norm_y) * GL_WORLD_HEIGHT  # Flip Y axis (pixels Y=0 is at top)
    gl_z = (norm_z - 0.5) * GL_WORLD_DEPTH   # Center at origin
    
    return gl_x, gl_y, gl_z

def pixels_to_gl_size(pixel_width, pixel_height, pixel_depth):
    """
    Convert pixel dimensions to OpenGL world dimensions.
    
    Args:
        pixel_width: Width in pixels
        pixel_height: Height in pixels
        pixel_depth: Depth in pixels
    
    Returns:
        tuple: (gl_width, gl_height, gl_depth) in OpenGL world dimensions
    """
    gl_width = pixel_width * PIXELS_TO_GL_X
    gl_height = pixel_height * PIXELS_TO_GL_Y
    gl_depth = pixel_depth * PIXELS_TO_GL_Z
    
    return gl_width, gl_height, gl_depth

def gl_coords_to_pixels(gl_x, gl_y, gl_z):
    """
    Convert OpenGL world coordinates to pixel coordinates.
    
    Args:
        gl_x: X coordinate in OpenGL world units
        gl_y: Y coordinate in OpenGL world units
        gl_z: Z coordinate in OpenGL world units
    
    Returns:
        tuple: (pixel_x, pixel_y, pixel_z) in pixel coordinates
    """
    # Convert from OpenGL world coordinates to normalized coordinates
    norm_x = (gl_x / GL_WORLD_WIDTH) + 0.5
    norm_y = 0.5 - (gl_y / GL_WORLD_HEIGHT)  # Flip Y axis back
    norm_z = (gl_z / GL_WORLD_DEPTH) + 0.5
    
    # Convert to pixel coordinates
    pixel_x = norm_x * WORLD_WIDTH
    pixel_y = norm_y * WORLD_HEIGHT
    pixel_z = norm_z * WORLD_DEPTH
    
    return pixel_x, pixel_y, pixel_z
