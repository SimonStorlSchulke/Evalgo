<i class='last-modified'>last modified 12:58 October 20 2018</i>
# Lösung A 2

**Dick**
*Kursiv*

- L1
- l2

```go
func UpdateConfig() {
	var conf Config
	jsondata, err := ioutil.ReadFile("./courseconfig.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(jsondata, &conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	Conf = conf
}
```

