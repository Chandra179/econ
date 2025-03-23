# Stock Project Makefile

# Kill a process running on a specific port
# Usage: make kill-port PORT=8080
.PHONY: kill-port
kill-port:
ifndef PORT
	@echo "Error: PORT parameter is required"
	@echo "Usage: make kill-port PORT=<port_number>"
	@exit 1
endif
ifeq ($(OS),Windows_NT)
	@echo "Killing process on port $(PORT) (Windows)"
	@netstat -ano | findstr :$(PORT) | findstr LISTENING > nul && \
	for /f "tokens=5" %a in ('netstat -ano ^| findstr :$(PORT) ^| findstr LISTENING') do taskkill /F /PID %a || echo No process found on port $(PORT)
else
	@echo "Killing process on port $(PORT) (Unix/Linux)"
	@lsof -i :$(PORT) -t | xargs -r kill -9 || echo "No process found on port $(PORT)"
endif 