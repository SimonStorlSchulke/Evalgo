<i class='last-modified'>last modified 10:32 October 16 2018</i>
## Login Check mit Go
```go
func loggedIn(r *http.Request) (bool, int) {
	var session, cErr = r.Cookie("session")

	if cErr == nil {
		sessionSplitted := strings.Split(session.Value, "<split>")
		storedMat, err := strconv.Atoi(sessionSplitted[0])
		if err != nil {
			return false, 0
		}
		storedPw := sessionSplitted[1]
		currentUser, err := user.FromMatrikel(storedMat)

		if currentUser.GetPassword() == storedPw {
			return true, storedMat
		}
	}
	return false, 0
}
```







