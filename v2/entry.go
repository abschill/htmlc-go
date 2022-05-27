package main

func main() {
	// this gets the full config file with .config as a property
	fsOptions, cwd := getFSOptions()
	// get .config property from full file options
	config := fsOptions.getOptionsFSToConfig()
	// create loader from the config
	loader := CreateLoader(config, cwd)
	loader.Print()
	// define keys from config that you want to bring into the runtime
	/**
	tester input
	{
		"a0": {
				"a1": "foobar"
			},
		"b0": {
			"b1": [
				{
					"name": "b1"
				}
			]
		}
	}**/
	for i, k := range fsOptions.getOptionsFSToInput() {
		// we will assign each input with a type, whether it be bound to a struct, a list, or just be a string/number of some kind
		if k.Type == "value" {
			println("input ", i)
			println("key: ", k.Key)
			println("value: ", k.Value)
		}
		// todo - other types
	}

	/**
	ok so we need to take an array of key value pairs from the config file and then determine the type of the value in each
	once we have the value we will know how to render it against any template expression, so we will need to set up a struct of some sort
	that can identify the bindable methods for the given handle type.
	**/
}
