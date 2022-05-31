Feature: Configuration
Customize your default loader options.

	Background:
	Given Read Config Into Runtime for a call to the createLoader() function

	Scenario: Valid Inline Arguments
		When args are inlined
		Then setup loader with those

	Scenario: Invalid Inline Arguments
		When args are not inlined
		Then check for "cwd/htmlc.json"

	Scenario: Invalid Config JSON
		When "cwd/htmlc.json" is not valid json
		Then throw error

	Scenario: Valid Config JSON
		When "cwd/htmlc.json" is valid json
		Then use the file as config

