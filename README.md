# Phantom-Ronin

A small 3D/2.5D platformer prototype built in Go using `raylib-go`. Run and jump across platforms, toggle between a side-on view and a pseudo-3D view, and progress through levels.

## Features

- Side-view and pseudo-3D camera toggle (press `R` in-game)
- Double jump mechanics
- CSV-driven level layout with platforms and per-level transitions
- Simple bounding-box collision against ground, walls, and platforms
- Basic UI menus: Start, Level Complete, and Game Over
- Optional debug overlay (player/camera coordinates, level id, bounding boxes)

## Requirements

- Go 1.21+ (module-enabled)
- raylib-go bindings and native raylib dependencies (installed by the module)

Dependencies are declared in `go.mod` and fetched automatically by `go` tooling.

## Setup & Run

```bash
go run .
```

This will open a window titled "Phantom-Ronin" at 800x480.

To build a binary:

```bash
go build -o Phantom.Ronin .
```

## Controls

- `A` / `D`: Move (side-view: left/right; pseudo-3D view: strafe on Z)
- `Space`: Jump (double-jump supported)
- `R`: Toggle camera mode (side-view ↔ pseudo-3D)
- `Esc`: Close the window

## Menus

- Start Menu: Click `Start` to begin, `Exit` to quit.
- Level Transition: After reaching a platform marked as final, click `Next` to advance.
- Game Over: Shown when there are no more levels; click `Exit` to quit.

## Levels

Levels are defined as CSV files under `level-maps/`. Each row defines a platform. Header is optional but supported.

Expected columns (order matters):

```
posX,posY,posZ,width,height,length,final
```

- `posX`, `posY`, `posZ` (float): Platform center position
- `width`, `height`, `length` (float): Platform size
- `final` (bool): If `true`, landing on this platform triggers a level transition

Example row:

```
5.0,-0.2,0.0,1.2,0.3,2.0,false
```

Current level files:

- `level-maps/level1.csv`
- `level-maps/level2.csv`

Level loading occurs in `resetGame` based on `GameState.Level`. When progressing beyond known levels, the game shows `Game Completed!`.

## Project Structure

```
- assets/               # Textures and art (e.g., background.png)
    - background.png
- background.go         # Background geometry/texture handling
- ground.go             # Ground plane drawing
- level.go              # Level struct and CSV parsing
- main.go               # Entry point: window, camera, loop, menus, collisions
- platform.go           # Platform definition and drawing
- player.go             # Player movement, jumping, gravity, and clamping
- utils.go              # Helper: GetBoundingBox
- wall.go               # Left/right bounds walls
- level-maps/           # CSV level definitions
    - level1.csv
    - level2.csv
```

## Technical Notes

- Physics is simple: velocity integration per frame; gravity applied when airborne.
- Collisions uses `raylib` AABB via `GetBoundingBox` and custom resolution logic for platforms (top/bottom/side handling, with minimal-axis push-out on X or Z depending on view).
- Camera follows player differently per mode:
  - Side-view: clamps to background extents and looks down X/Y
  - Pseudo-3D: offsets on X/Y with target at player X/Y
- Debug tools: Toggle `state.isDebug = true` in `main.go` for overlays and bounding boxes.

## Building on Linux

If you encounter native dependency issues for raylib, consult the raylib-go docs. Typical packages for Debian/Ubuntu-based systems include:

```bash
sudo apt-get update
sudo apt-get install -y build-essential pkg-config libx11-dev libxi-dev libgl1-mesa-dev libxcursor-dev libxrandr-dev libxinerama-dev
```
