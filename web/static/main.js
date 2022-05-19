console.log("JS Loaded")

const url = "127.0.0.1:5555"

var inputForm = document.getElementById("inputForm")

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


inputForm.addEventListener("submit", (e)=>{

    //prevent auto submission
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url,{

        method:"POST",
        body:formdata,
    }).then(
        response => response.text()
    ).then(
        (data) => {
            outputs = data.split("\n")
            for (let out of outputs) {
                document.getElementById("serverMessageBox").appendChild(createPre(out))}
            }

    ).catch(
        error => console.error(error)
    )

})