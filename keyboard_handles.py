from constants import *

def keyboard_down(key, x, y):
    char_key = key.decode("utf-8").lower()
    keys_pressed[char_key] = True
    # print(f"Key pressed: {char_key}")


def keyboard_up(key, x, y):
    char_key = key.decode("utf-8").lower()
    keys_pressed[char_key] = False
    # print(f"Key released: {char_key}")
