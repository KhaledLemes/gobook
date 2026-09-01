async function PedirLogin() {
    const but = document.getElementById('entrar')
    const email = document.getElementById('email')
    const senha = document.getElementById('senha')
    const req = await fetch("/api/v1/login", {
        body: {
            "email": email.value,
            "senha": senha.value,
        }
    })
}


const but = document.getElementById('entrar')
but.addEventListener('click', async (e) => {

    e.preventDefault()

    const email = document.getElementById('email')
    const senha = document.getElementById('senha')
    const loginErr = document.getElementById('login-err')
    loginErr.style.color = 'red';

    if (email.value == "" || senha.value == "") {
        loginErr.innerText = ""
        loginErr.innerText = "Campo de e-mail ou senha vazios."
        return null
    } else {
        loginErr.innerText = ""
    }

    const req = await fetch("/api/v1/login", {
        body: JSON.stringify({
            "email": email.value,
            "senha": senha.value
        }),
        method: 'POST',
    })

    switch (req.status) {
        case 200:
            loginErr.style.color = 'green'
            loginErr.innerText = "Login feito! Espere o Khaled parar de preguiça para isso significar alguma coisa <3"
            break
        case 400:
            loginErr.innerText = "Erro no corpo da requisição. Revise os dados."
            break
        case 401:
            loginErr.innerText = "Login ou senha inválidos."
            break
        case 500:
            loginErr.innerText = "Houve um erro interno no nosso servidor. Tente novamente mais tarde!"
            break
        default:
            loginErr.innerText = "Houve um erro desconhecido. Tente novamente mais tarde"
            break
    }
    if (req.status === 200) {

    } else {
        console.log(req.body)
    }
})
    