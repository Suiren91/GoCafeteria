#!/usr/bin/env fish

echo "==> Lint"
golangci-lint run; or exit 1

echo "==> Build"
go build -v ./...; or exit 1

echo "==> Test"
make coverage; or exit 1

echo "==> Check coverage"
set COVERAGE (go tool cover -func=coverage.out | grep '^total:' | awk '{print $3}' | tr -d '%')
echo "Total coverage: $COVERAGE%"
if awk "BEGIN{exit !($COVERAGE<80)}"
    echo "Coverage $COVERAGE% is below the 80% threshold"
    exit 1
end
echo "All checks passed ✅"

