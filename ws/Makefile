# ws package Makefile
# Verify all centrifuge-go public resources are exported

.PHONY: check check-types check-funcs check-consts

# Check all exports match
check: check-types check-funcs check-consts
	@echo ""
	@echo "✓ All centrifuge-go exports are re-exported in ws package"

# Check types match
check-types:
	@echo "Checking types..."
	@go doc -short github.com/centrifugal/centrifuge-go | grep "^type" | awk '{print $$2}' | sort > /tmp/centrifuge_types.txt
	@go doc -short . | grep "^type" | awk '{print $$2}' | sort > /tmp/ws_types.txt
	@if diff -q /tmp/centrifuge_types.txt /tmp/ws_types.txt > /dev/null; then \
		echo "  ✓ Types: $$(wc -l < /tmp/ws_types.txt | tr -d ' ')/$$(wc -l < /tmp/centrifuge_types.txt | tr -d ' ') match"; \
	else \
		echo "  ✗ Types mismatch:"; \
		echo "  Missing:"; \
		diff /tmp/centrifuge_types.txt /tmp/ws_types.txt | grep "^<" | sed 's/< /    /'; \
		echo "  Extra:"; \
		diff /tmp/centrifuge_types.txt /tmp/ws_types.txt | grep "^>" | sed 's/> /    /'; \
		exit 1; \
	fi

# Check functions/vars match (centrifuge-go has funcs, ws re-exports as vars)
check-funcs:
	@echo "Checking functions/vars..."
	@# centrifuge-go: extract standalone funcs (New*, With*) and vars (Err*)
	@go doc -all github.com/centrifugal/centrifuge-go | grep -E "^func (New|With)[A-Za-z]+" | sed 's/func \([A-Za-z]*\).*/\1/' | sort > /tmp/centrifuge_funcs.txt
	@go doc github.com/centrifugal/centrifuge-go | grep -E "^var Err" | sed 's/var \([A-Za-z]*\).*/\1/' | sort >> /tmp/centrifuge_funcs.txt
	@sort -u /tmp/centrifuge_funcs.txt -o /tmp/centrifuge_funcs.txt
	@# ws: extract vars from "// Constructors", "// History options", "// Errors" sections
	@grep -A3 "^// Constructors" centrifugo.go | grep -E "^\s+[A-Z]" | awk '{print $$1}' | sort > /tmp/ws_funcs.txt
	@grep -A4 "^// History options" centrifugo.go | grep -E "^\s+[A-Z]" | awk '{print $$1}' | sort >> /tmp/ws_funcs.txt
	@grep -A2 "^// Errors" centrifugo.go | grep -E "^\s+[A-Z]" | awk '{print $$1}' | sort >> /tmp/ws_funcs.txt
	@sort -u /tmp/ws_funcs.txt -o /tmp/ws_funcs.txt
	@if diff -q /tmp/centrifuge_funcs.txt /tmp/ws_funcs.txt > /dev/null; then \
		echo "  ✓ Funcs/Vars: $$(wc -l < /tmp/ws_funcs.txt | tr -d ' ')/$$(wc -l < /tmp/centrifuge_funcs.txt | tr -d ' ') match"; \
	else \
		echo "  ✗ Funcs/Vars mismatch:"; \
		echo "  Missing:"; \
		diff /tmp/centrifuge_funcs.txt /tmp/ws_funcs.txt | grep "^<" | sed 's/< /    /'; \
		echo "  Extra:"; \
		diff /tmp/centrifuge_funcs.txt /tmp/ws_funcs.txt | grep "^>" | sed 's/> /    /'; \
		exit 1; \
	fi

# Check constants match
check-consts:
	@echo "Checking constants..."
	@# centrifuge-go: extract constants (LogLevel*, State*, SubState*, DeltaType*)
	@go doc -all github.com/centrifugal/centrifuge-go | grep -E "^\s+(LogLevel[A-Z]|State[A-Z]|SubState[A-Z]|DeltaType[A-Z])" | awk '{print $$1}' | sort -u > /tmp/centrifuge_consts.txt
	@# ws: extract constants from source
	@grep -E "^\s+(LogLevel[A-Z]|State[A-Z]|SubState[A-Z]|DeltaType[A-Z])" centrifugo.go | awk '{print $$1}' | sort -u > /tmp/ws_consts.txt
	@if diff -q /tmp/centrifuge_consts.txt /tmp/ws_consts.txt > /dev/null; then \
		echo "  ✓ Constants: $$(wc -l < /tmp/ws_consts.txt | tr -d ' ')/$$(wc -l < /tmp/centrifuge_consts.txt | tr -d ' ') match"; \
	else \
		echo "  ✗ Constants mismatch:"; \
		echo "  Missing:"; \
		diff /tmp/centrifuge_consts.txt /tmp/ws_consts.txt | grep "^<" | sed 's/< /    /'; \
		echo "  Extra:"; \
		diff /tmp/centrifuge_consts.txt /tmp/ws_consts.txt | grep "^>" | sed 's/> /    /'; \
		exit 1; \
	fi

# Show detailed comparison
diff:
	@echo "=== centrifuge-go ==="
	@echo "Types:" && cat /tmp/centrifuge_types.txt
	@echo "Funcs:" && cat /tmp/centrifuge_funcs.txt
	@echo "Consts:" && cat /tmp/centrifuge_consts.txt
	@echo ""
	@echo "=== ws ==="
	@echo "Types:" && cat /tmp/ws_types.txt
	@echo "Funcs:" && cat /tmp/ws_funcs.txt
	@echo "Consts:" && cat /tmp/ws_consts.txt
