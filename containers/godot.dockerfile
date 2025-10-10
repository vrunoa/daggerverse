FROM barichello/godot-ci:4.4

RUN apt-get update && apt-get install -y --no-install-recommends \
    xvfb x11vnc novnc websockify x11-utils xdotool \
    && rm -rf /var/lib/apt/lists/*

RUN which xvfb-run || echo "xvfb-run not found"

RUN cat > /entrypoint.sh <<'EOF'
#!/bin/bash
set -euo pipefail

# Start virtual X server on :99
Xvfb :99 -screen 0 1280x720x24 &

# Start VNC server pointing to :99
x11vnc -display :99 -nopw -listen 0.0.0.0 -forever -shared &

# Wait a moment for VNC to start
sleep 2

# Set cursor position
xdotool mousemove 0 0

# Start noVNC websockify on port 8080
websockify --web=/usr/share/novnc/ 8080 localhost:5900 &

export LIBGL_ALWAYS_SOFTWARE=1

# Pre-import the project to avoid first-time import delays during tests
godot --headless --import --path /src

# Run Godot in headless mode to execute GUT tests
echo "godot -s --path /src ./addons/gut/gut_cmdln.gd $@"
DISPLAY=:99 godot -s --path /src ./addons/gut/gut_cmdln.gd $@
exit $?
EOF

RUN chmod +x /entrypoint.sh

ENV DISPLAY=:99

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
