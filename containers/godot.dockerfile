FROM barichello/godot-ci:4.4

RUN apt-get update && apt-get install -y --no-install-recommends \
    xvfb x11vnc novnc websockify x11-utils xdotool \
    && rm -rf /var/lib/apt/lists/*

RUN which xvfb-run || echo "xvfb-run not found"

RUN cat > /entrypoint.sh <<'EOF'
#!/bin/bash
set -euo pipefail

if [[ -z $1 ]]; then
    echo "Missing tests folder argument"
    exit 1
fi

if [[ -z $2 ]]; then
    echo "Missing gut config path argument"
    exit 1
fi

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

DISPLAY=:99 godot -d -s --path /src ./addons/gut/gut_cmdln.gd -gdir $1 -gconfig=$2 -gexit
EOF
 
RUN chmod +x /entrypoint.sh

ENV DISPLAY=:99

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
