import csv
import random
import os

def generate_level(level_number, num_platforms):
    """Generates a single level file."""
    filepath = f"level-maps/level{level_number}.csv"
    
    with open(filepath, 'w', newline='') as csvfile:
        fieldnames = ['posX', 'posY', 'posZ', 'width', 'height', 'length', 'final']
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        writer.writeheader()

        pos_x = 3.0
        initial_pos_y = -2.0
        max_pos_y = 10.0
        
        # Make sure platforms are spread out until around 28
        x_increment = (28.0 - pos_x) / num_platforms

        previous_pos_z = 0.0

        for i in range(num_platforms):
            is_final = (i == num_platforms - 1)
            
            width = random.uniform(0.8, 1.0)
            length = random.uniform(0.4, 0.8) # Smaller length range
            height = 0.3

            # Calculate valid posZ range based on length to stay within ground boundaries
            min_z = -0.9 + length / 2
            max_z = 1.1 - length / 2

            if i == 0:
                pos_z = random.uniform(-0.3, 0.3)
            else:
                # Add "trick" platforms with some probability
                if random.random() < 0.2: # 20% chance of a trick
                    delta_z = random.choice([-1, 1]) * random.uniform(0.4, 0.6)
                else:
                    delta_z = random.uniform(-0.1, 0.1)
                
                pos_z = previous_pos_z + delta_z
                
                # Clamp pos_z to be within the valid range
                if pos_z < min_z:
                    pos_z = min_z
                elif pos_z > max_z:
                    pos_z = max_z
            
            previous_pos_z = pos_z

            if i == 0: # First platform
                pos_y = random.uniform(-3.0, -2.0) # Ensure the first platform is low
            else:
                # Calculate posY to have a general upward trend
                base_pos_y = initial_pos_y + ((max_pos_y - initial_pos_y) / (num_platforms -1)) * i if num_platforms > 1 else initial_pos_y
                pos_y = base_pos_y + random.uniform(-1.0, 1.0)

                # Clamp posY to be safe
                if pos_y > 10:
                    pos_y = 10
                elif pos_y < -2:
                    pos_y = -2

            writer.writerow({
                'posX': f"{pos_x:.6f}",
                'posY': f"{pos_y:.6f}",
                'posZ': f"{pos_z:.6f}",
                'width': f"{width:.6f}",
                'height': height,
                'length': f"{length:.6f}",
                'final': str(is_final).lower()
            })

            # Update for next platform
            pos_x += x_increment + random.uniform(-0.2, 0.2)

def main():
    """Generates 10 levels."""
    if not os.path.exists("level-maps"):
        os.makedirs("level-maps")

    for i in range(1, 11):
        num_platforms = random.randint(20, 30)
        generate_level(i, num_platforms)
        print(f"Generated level-maps/level{i}.csv")

if __name__ == "__main__":
    main()
