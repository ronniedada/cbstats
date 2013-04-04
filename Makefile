build:
	go clean;
	go install;

clean:
	rm -fr cbstat;
	go clean;

jshint:
	jshint --config jshint_conf.js static/js

test_backend:
	go test -bench='.*' -v;

test: test_backend jshint
