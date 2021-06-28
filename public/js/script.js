function disableSelectionForm() {
    $('#from').prop('disabled', true)
    $('#to').prop('disabled', true)
}

function allowLineSelection() {
    $('#from').prop('disabled', false)
    $('#to').prop('disabled', false)
}

var loader = function(e){
    let file = e.target.files;
    let fileName = '<span>Fichier sélectionné : </span>'+file[0].name
    let filePicker = document.getElementById("filePicker")
    filePicker.innerHTML = fileName
    filePicker.classList.add("activeFilePicker")
}

$("#calculLaunch").on("click", function() {
    var requestBody = {
        withHeader  : $("#withHeader").prop('checked'),
        seperator   : $("#file-sep").val(),
        somme       : $("#somme").prop('checked'),
        moyenne     : $("#moyenne").prop('checked'),
        median      : $("#mediane").prop('checked'),
        maxValue    : $("#maximum").prop('checked'),
        entireFile  : $("#all").prop('checked'),
        fromLine    : $("#from").val(),
        toLine      : $("#to").val()
    }

    

    var fd = new FormData();
    var files = $('#file')[0].files;
    
    // Check file selected or not
    if(files.length > 0 ){
        if(requestBody.somme || requestBody.maxValue || requestBody.median || requestBody.moyenne) {
            fd.append('file',files[0]);
            fd.append('withHeader', requestBody.withHeader);
            fd.append('params', JSON.stringify(requestBody));
            $.ajax({
                url: '/upload',
                type: 'post',
                data: fd,
                contentType: false,
                processData: false,
                success: function(response){
                    $('#results').html(`<div class="result-text result-success">${response}<div>`)
                },
            });
        }else $('#results').html(`<div class="result-text result-error">Veuillez choisir au moins une opération!<div>`)
    }
    else $('#results').html(`<div class="result-text result-error">Veuillez choisir un fichier SVP.<div>`)
    
})

let fileSource = document.getElementById("file");
fileSource.addEventListener("change", loader)
