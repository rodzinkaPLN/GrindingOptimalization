.PHONY: sb
sb:
	cd web && npm run storybook -- --disable-telemetry
.PHONY: web
web:
	cd web && npm start