# struct2mapstrstr
convert struct to map[string]string


```
type MssCheck struct {
	Test string `mss:"test"`
	Tas string  `mss:"tassss"`
	Nas int     `mss:"nas"`
	X *string   `mss:"x"`
}
v := MssCheck{
	Test: "iaiu",
	Tas:  "ooo",
	Nas:  3,
}
```  

to

```
map[string]string{
  "test": "iaiu",
  "tassss": "ooo",
  "nas": "3",
}
```
