# hxx

golang hex dumps made easy.

hxx is library for turning values into memory images that can be printed as hex dumps (or binary, octal or dec dumps).

Objects for dumping are wrapped in a `Dump` object using `NewDump`.  `Dump` objects can then be printed using standard formatting functions such as `fmt.Printf`.

Example:
```

	input0 := `Lorem Ipsum.`
	fmt.Printf("%v\n", hxx.NewDump(input0))

	// Output:
	// 0: 4c 6f 72 65 6d 20 49 70 73 75 6d 2e              Lorem Ipsum.

```
