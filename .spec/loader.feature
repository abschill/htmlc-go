Feature: Loader Module
Load Chunks into runtime engine based on config feature and local paths derived

Background:
	Given Create Loader part of createLoader(), following config generation

	Scenario:
		When paths are improperly defined, throw error
		Then exit with code 1 or ignore depending on config
	Scenario:
		When paths are properly defined, set up the structure for loader and the metadata for the given chunk syntax.
		If any data was entered into the config, preload it into the non-static templates.
		Then return the loader
