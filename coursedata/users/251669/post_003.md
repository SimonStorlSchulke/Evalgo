<i class='last-modified'>last modified 9:37 October 15 2018</i>
```js
xhr.onreadystatechange = function () {
                if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                    var posttext = xhr.responseText;
                    var lines = posttext.split('\n');
                    lines.splice(0, 1);
                    posttext = lines.join('\n');
                    posttext += "\n";
                    document.getElementById("textarea").value = posttext;
                    checkText();
                }
            };
```