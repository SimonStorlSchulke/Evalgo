function deleteSession() {
    document.cookie = "session" + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

var currentMatrikel;

function selectStudent(matrikel) {
    window.currentMatrikel = matrikel;
    var currentCard = $(matrikel)
    $(".selected-profile").removeClass("selected-profile");
    var m = ".".concat(matrikel);
    $(m).addClass("selected-profile");
}


function getPostNumber() {
    return new URL(document.URL).searchParams.get("nr");
}

//select active assignment
window.onload = function () {
    currentMatrikel = new URL(document.URL).searchParams.get("mat");
    selectStudent(currentMatrikel);
    makeRequest(currentMatrikel);
    var classStr = ".a";
    if (getPostNumber() > 0) {
        classStr = classStr.concat(getPostNumber())
        $(classStr).addClass("active");
    } else {
        $(".a1").addClass("active");
    }
}

//Ajax Post Loader
function makeRequest(matrikel) {
    window.currentMatrikel = matrikel;
    if(currentMatrikel < 1) {
        return
    }
    var xhr = new XMLHttpRequest();

    //open request to /matrikel/post/number
    var path = '/';
    var urlArray = window.location.pathname.split('/');

    var postNumber = getPostNumber();

    //Open path to postnumber, else take post 1
    if (postNumber > 0) {
        xhr.open('GET', path.concat(matrikel, "/post/", postNumber), true);
    } else {
        xhr.open('GET', path.concat(matrikel, "/post/1"), true);
    }

    //handle response
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            document.getElementById("post-area").innerHTML = xhr.responseText;
            $("#navigation").addClass("hidden");
            //highlightJS
            $('pre > code').each(function () {
                hljs.highlightBlock(this);
            });
        }
    };
    xhr.send();
}

//Show Course Info in PostArea
function showInfo() {

    selectStudent(0);
    $(".nav-item").removeClass("active");

    var xhr = new XMLHttpRequest();
    //Open path
    xhr.open('GET', "/info", true);

    //handle response
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            document.getElementById("post-area").innerHTML = xhr.responseText;
            //highlightJS
            $('pre > code').each(function () {
                hljs.highlightBlock(this);
            });
        }
    };
    xhr.send();
}

function postLink(number) {
    var path = "?nr=".concat(number, "&mat=", currentMatrikel);
    window.location = path;
}