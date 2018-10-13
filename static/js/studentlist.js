window.onload = function () {
    //set params (matrikel and postnumber) to 0 / 0 if not defined in url.
    cMat = MatrikelFromUrl();
    cNr = PostNrFromUrl();
    if (cMat == null || cNr == null) {
        setParams(0, 0)
        cMat = cNr = 0;
    }
    switchStudent(cMat);
    UpdateTaskSwitcher(cNr);

    var matCl = ".".concat(matrikel);
    $(matCl).addClass("selected-profile");
}

//But Not Matrikel
function switchTask(nr) {
    mat = MatrikelFromUrl();
    if (mat == 0 && nr == 0) {
        loadToViewer("./info");
        feedbackHide();
    } else if (mat == 0 && nr != 0) {
        loadToViewer("./".concat("task/", nr));
        feedbackHide();
    }
    else if(mat != 0 && nr == 0) {
        loadToViewer("./".concat("profile/" ,mat));
        feedbackHide();
    } else {
        loadToViewer("./".concat(mat, "/post/", nr));
        feedbackShow();
    }
    setParams(nr, mat);
    UpdateTaskSwitcher(nr);
}

//But not Postnuber
function switchStudent(mat) {
    nr = PostNrFromUrl();
    if (mat == 0 && nr == 0) {
        loadToViewer("./info");
        feedbackHide();
    } else if (mat == 0 && nr != 0) {
        loadToViewer("./".concat("task/", nr));
        feedbackHide();
    }
    else if(mat != 0 && nr == 0) {
        loadToViewer("./".concat("profile/" ,mat));
        feedbackHide();
    } else {
        loadToViewer("./".concat(mat, "/post/", nr));
        feedbackShow();
    }
    setParams(nr, mat);
    UpdateTaskSwitcher(nr);

    //Visual Border
    $(".selected-profile").removeClass("selected-profile");
    matCl = ".".concat(mat);
    $(matCl).addClass("selected-profile");
}

function showInfo() {
    loadToViewer("./info");
    setParams(0, 0);
    UpdateTaskSwitcher(0);
    $(".selected-profile").removeClass("selected-profile");
    $(".0").addClass("selected-profile");
    feedbackHide();
}

function UpdateTaskSwitcher(nr) {
    $(".awb").removeClass("active");
    str = ".a".concat(nr);
    $(str).addClass("active");
}

function deleteSession() {
    document.cookie = "session" + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

//load content to #post-area
function loadToViewer(src) {
    var xhr = new XMLHttpRequest();

    //Open path
    xhr.open('GET', src, true);

    //handle response
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            document.getElementById("post-area").innerHTML = xhr.responseText;
            $("#navigation").addClass("hidden"); //Hide buttons on posts when loaded into viewer
            //highlightJS
            $('pre > code').each(function () {
                hljs.highlightBlock(this);
            });
        }
    };
    xhr.send();
}

function setParams(postNumber, matrikel) {
    querystring = "?nr=" + postNumber + "&mat=" + matrikel;
    history.pushState("", document.title, querystring);
}

function setMat(mat) {
    nr = PostNrFromUrl();
    setParams(nr, mat)
}

function setNr(nr) {
    mat = MatrikelFromUrl();
    setParams(nr, mat)
}

function feedbackShow() {
    if($("#feedback").length) {
        $("#feedback").removeClass("hidden");
    }
}

function feedbackHide() {
    if($("#feedback").length) {
        $("#feedback").addClass("hidden");
    }
}

function MatrikelFromUrl() {
    return new URL(document.URL).searchParams.get("mat");
}

function PostNrFromUrl() {
    return new URL(document.URL).searchParams.get("nr");
}

function PostLink() {
    nr = PostNrFromUrl();
    window.location = "./post?nr="+nr;
}