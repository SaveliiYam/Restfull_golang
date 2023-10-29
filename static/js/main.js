document
  .getElementById("button")
  .addEventListener("click", function(){
    const name = document.getElementById("name")
    const age = document.getElementById("age")
    const mail = document.getElementById("mail")

    $.post("https://backend.marathon.ivashev.com/save", {
      name: name.value,
      age:  age.value,
      email: mail.value
    }, function(){
      alert("Данные сохранены")
    })
  })

  