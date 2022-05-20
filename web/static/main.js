console.log("JS Loaded")

const url = "127.0.0.1:5555"

const inputForm = document.getElementById("inputForm");
const fileToLoadButton = document.getElementById("fileToLoad");

function createPre(value) {
    let pre = document.createElement('pre');
    pre.dataset.prefix = "~"
    if (value === "ERRORS:") {
        pre.className = "text-warning"
        pre.dataset.prefix = "!"
    } else if (value.includes("Ru Output:")){
        pre.className = "text-success"
        pre.dataset.prefix = ">"
    }
    pre.textContent = value;
    return pre;
}

function textAreaAdjust(element) {
    element.style.height = "1px";
    element.style.height = (25+element.scrollHeight)+"px";
}

function toggleTheme() {
    const theme = document.documentElement.dataset.theme;
    if (theme === "dark"){
        document.documentElement.dataset.theme = "night"
    } else {
        document.documentElement.dataset.theme = "dark"
    }
}

function runRuScript(){
    const formdata = new FormData(inputForm)
    fetch(url,{

        method:"POST",
        body:formdata,
    }).then(
        response => response.text()
    ).then(
        (data) => {
            const outputs = data.split("\n")
            for (let out of outputs) {
                document.getElementById("serverMessageBox").appendChild(createPre(out))}
        }

    ).catch(
        error => console.error(error)
    )
}

function loadFileAsText(){
    try {
        const fileToLoad = document.getElementById("fileToLoad").files[0];

        const fileReader = new FileReader();
        const scriptBox = document.getElementById("message")
        fileReader.onload = function(fileLoadedEvent){
            scriptBox.value = fileLoadedEvent.target.result;

            textAreaAdjust(scriptBox)

        };

        fileReader.readAsText(fileToLoad, "UTF-8");

        runRuScript();

        document.getElementById("buttonSubmit").classList.add('btn-primary')
        document.getElementById("buttonSubmit").classList.remove('btn-disabled')
        const errorFile = document.getElementById("noFileError");
        errorFile.classList.add("hidden");
    } catch (error) {
        console.error(error);
        const errorFile = document.getElementById("noFileError");
        errorFile.classList.remove("hidden");
    }
}

fileToLoadButton.addEventListener("change", (e)=>{
    // btn-disabled buttonSubmit
    if( e.target.value !== "") {
        console.log("AQUI ESTOY")
        document.getElementById("buttonSubmit").classList.remove('btn-primary')
        document.getElementById("buttonSubmit").classList.add('btn-disabled')
    }
})

inputForm.addEventListener("submit", (e)=>{

    //prevent auto submission
    e.preventDefault()

    runRuScript()

})