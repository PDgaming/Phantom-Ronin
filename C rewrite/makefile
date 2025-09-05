CC := gcc
CFLAGS := -std=c11 -Wall -Wextra -O2
LDLIBS := -lraylib -lm -lpthread -ldl -lrt -lX11

SRCDIR := ./
SOURCES := $(SRCDIR)src/main.c $(SRCDIR)src/background.c $(SRCDIR)src/ground.c $(SRCDIR)src/player.c $(SRCDIR)src/wall.c $(SRCDIR)src/platform.c $(SRCDIR)src/level.c
OBJECTS := $(SOURCES:.c=.o)
TARGET := phantom_ronin

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $(OBJECTS) -o $@ $(LDLIBS)

$(SRCDIR)/%.o: $(SRCDIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@

run: $(TARGET)
	./$(TARGET)

clean:
	rm -f $(OBJECTS) $(TARGET)