async function fazerLogin(email, senha, errCont) {
    const req = await fetch("/api/v1/login", {
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "email": email,
            "senha": senha
        }),
        method: 'POST',
    })

    if (req.status !== 200) {
        const response = await req.json()
        errCont.innerText = response.error
    } else {
        window.location.replace("/home")
    }
}

document.addEventListener('DOMContentLoaded', (e) => {
    const email = document.getElementById('email')
    const senha = document.getElementById('senha')
    const loginErr = document.getElementById('login-err')
    loginErr.style.color = 'red';

    const but = document.getElementById('entrar')
    but.addEventListener('click', async (e) => {
        e.preventDefault()


        if (email.value == "" || senha.value == "") {
            loginErr.innerText = ""
            loginErr.innerText = "Campo de e-mail ou senha vazios."
            return null
        } else {
            loginErr.innerText = ""
        }

        await fazerLogin(email.value, senha.value, loginErr)

    })

    const jumpStep = document.getElementById('jump-step')
    if (jumpStep) {
        jumpStep.addEventListener('click', async (e) => {
            e.preventDefault()
            loginErr.innerText = ""
            email.value = "amoabra3ao@gmail.com"
            senha.value = "Khaled13@"
            await fazerLogin("amoabra3ao@gmail.com", "Khaled13@", loginErr)
        })
    }
})

