FROM golang:1.24.0-alpine3.21

# Go Cover Treemap
# Visualization of Go code coverage via treemaps with heat
# https://github.com/nikolaydubina/go-cover-treemap
RUN go install github.com/nikolaydubina/go-cover-treemap@v1.5.0

COPY containers/coverage.sh /usr/local/bin/coverage.sh
RUN chmod +x /usr/local/bin/coverage.sh

ENTRYPOINT ["/usr/local/bin/coverage.sh"]

CMD ["/usr/local/bin/coverage.sh",  "100"]
