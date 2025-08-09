def check_aabb_collision(
    rect1_x, rect1_y, rect1_width, rect1_height,
    rect2_x, rect2_y, rect2_width, rect2_height
):
    x_overlap = (rect1_x < rect2_x + rect2_width) and (rect1_x + rect1_width > rect2_x)
    y_overlap = (rect1_y < rect2_y + rect2_height) and (rect1_y + rect1_height > rect2_y)
    return x_overlap and y_overlap

def check_aabb_collision_3d(
    rect1_x, rect1_y, rect1_z, rect1_width, rect1_height, rect1_depth,
    rect2_x, rect2_y, rect2_z, rect2_width, rect2_height, rect2_depth
):
    x_overlap = (rect1_x < rect2_x + rect2_width) and (rect1_x + rect1_width > rect2_x)
    y_overlap = (rect1_y < rect2_y + rect2_height) and (rect1_y + rect1_height > rect2_y)
    z_overlap = (rect1_z < rect2_z + rect2_depth) and (rect1_z + rect1_depth > rect2_z)
    return x_overlap and y_overlap and z_overlap
