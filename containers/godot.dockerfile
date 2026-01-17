FROM barichello/godot-ci:4.5.1

RUN apt-get update && apt-get install -y --no-install-recommends \
    xvfb x11vnc novnc websockify x11-utils xdotool x11-xserver-utils \
    libxcursor1 libwayland-client0 libxrandr2 libxinerama1 libxi6 libxfixes3 \
    && rm -rf /var/lib/apt/lists/*

RUN which xvfb-run || echo "xvfb-run not found"

RUN cat > /entrypoint.sh <<'EOF'
#!/bin/bash
set -euo pipefail

# Start virtual X server on :99
Xvfb :99 -screen 0 1280x720x24 &
XVFB_PID=$!

# Wait for Xvfb to be ready
for i in {1..30}; do
    if xdpyinfo -display :99 >/dev/null 2>&1; then
        break
    fi
    sleep 0.1
done

# Start VNC server pointing to :99
x11vnc -display :99 -nopw -listen 0.0.0.0 -forever -shared &

# Wait a moment for VNC to start
sleep 2

# Set cursor position
xdotool mousemove 0 0

# Start noVNC websockify on port 8080
websockify --web=/usr/share/novnc/ 8080 localhost:5900 &

export LIBGL_ALWAYS_SOFTWARE=1
export DISPLAY=:99

# Pre-import the project to avoid first-time import delays during tests
godot --headless --import --path /src

# Run Godot in headless mode to execute GUT tests
echo "godot -s --path /src ./addons/gut/gut_cmdln.gd $@"
godot -s --path /src ./addons/gut/gut_cmdln.gd $@
exit $?
EOF

RUN chmod +x /entrypoint.sh

ENV DISPLAY=:99

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
