<i class='last-modified'>last modified 12:13 October 16 2018</i>
## Return User Struct from Matrikel
```go
func FromMatrikel(matrikel int) (User, error) {
	var us User
	jsondata, err := ioutil.ReadFile(fmt.Sprintf("./coursedata/users/%v/profile.json", matrikel))
	if err != nil {
		fmt.Println(err)
		return us, err
	}

	err = json.Unmarshal(jsondata, &us)
	if err != nil {
		fmt.Println(err)
	}
	return us, err
}
```
