venv: requirements.txt
	python3 -m venv venv
	venv/bin/python3 -m pip install --upgrade pip
	venv/bin/pip3 install --upgrade wheel 
	venv/bin/pip3 install --upgrade setuptools
	venv/bin/pip install -r requirements.txt

.PHONY: setup
setup: venv

.PHONY: unittest
unittest:
	venv/bin/python3 -m unittest -v tests-unittest/*.py

.PHONY: pytest
pytest:
	venv/bin/python3 -m pytest -v tests-pytest

.PHONY: test
test: pytest unittest

.PHONY: style
style:
	venv/bin/flake8 --exclude=venv

.PHONY: format
format:
	venv/bin/black --exclude=venv .

.PHONY: clean
clean:
	rm -rf venv
