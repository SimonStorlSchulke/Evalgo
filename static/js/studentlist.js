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

    if(currentMatrikel == 0) {
        loadTask();
    } else {

    if (getPostNumber() > 0) {
        classStr = classStr.concat(getPostNumber())
        $(classStr).addClass("active");
    } else if (currentMatrikel == null){
        $(".a0").addClass("active");
        showInfo();
    }
}
}

//load content to #post-area
function loadContent(src) {
    var xhr = new XMLHttpRequest();

    //Open path
    xhr.open('GET', src, true);

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

function makeRequest(matrikel) {

    //Show Feedback
    $(feedback).removeClass("hidden");

    window.currentMatrikel = matrikel;
    if(currentMatrikel < 1) {
        return
    }
    //open request to /matrikel/post/number
    var path = './';
    var urlArray = window.location.pathname.split('/');
    var postNumber = getPostNumber();
    
    //Open path to postnumber, else take post 1
    if (postNumber > 0) {
        loadContent(path.concat(matrikel, "/post/", postNumber));
    } else {
        loadContent(path.concat("profile/" ,matrikel));
    }
    querystring = "?nr=" + postNumber + "&mat=" + matrikel;
    history.pushState("", document.title, querystring);
}

function loadTask() {

    //Hide Feedback
    $(feedback).addClass("hidden");
    var path = './';
    var urlArray = window.location.pathname.split('/');
    postNumber = getPostNumber();

    //Open path to postnumber, else take post 1
    if (postNumber > 0) {
        loadContent(path.concat("task/", postNumber));
    } else {
        loadContent(path.concat("task/1"));
    }
    querystring = "?nr=" + postNumber + "&mat=0";
    history.pushState("", document.title, querystring);

}

//Ajax Post Loader
function loadAssignment(postNumber) {
    var matrikel = new URL(document.URL).searchParams.get("mat");

    if (matrikel == 0) {
        loadContent("./".concat("/task/", postNumber));
    }

    if (postNumber > 0) {
    loadContent("./".concat(matrikel, "/post/", postNumber));
    } else {
        loadContent("./".concat("profile/" ,matrikel));
    }
    $(".awb").removeClass("active");
    classStr = ".a".concat(postNumber)
    $(classStr).addClass("active");
    querystring = "?nr=" + postNumber + "&mat=" + matrikel;
    history.pushState("", document.title, querystring);
}

//Show Course Info in PostArea
function showInfo() {
    loadContent("./info");
    history.pushState("", document.title, "?nr=0&mat=0");
}

function postLink(number) {
    var path = "?nr=".concat(number, "&mat=", currentMatrikel);
    window.location = path;
}