FROM barichello/godot-ci:4.4
RUN apt-get update && apt-get install -y --no-install-recommends xvfb
RUN which xvfb-run || echo "xvfb-run not found"
ENV DISPLAY=:99
